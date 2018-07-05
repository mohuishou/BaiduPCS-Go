package user

import (
	"github.com/gin-gonic/gin"
	"github.com/iikira/Baidu-Login"
	"github.com/iikira/BaiduPCS-Go/internal/pcsapi/api"
	"github.com/iikira/BaiduPCS-Go/internal/pcsconfig"
)

var bc = baidulogin.NewBaiduClinet()
var loginJson *baidulogin.LoginJSON

const (
	VerifyPhone = "mobile"
	VerifyEmail = "email"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	code := c.PostForm("code")
	codeStr := c.PostForm("code_str")
	// 参数校验
	if username == "" || password == "" {
		api.BadRequest(c, "用户名密码不能为空", nil)
		return
	}
	loginJson = bc.BaiduLogin(username, password, code, codeStr)
	switch loginJson.ErrInfo.No {
	case "0":
		setUser(c, loginJson)
	case "400023", "400101": // 需要验证手机或邮箱
		api.Error(c, 1, "需要验证手机号或邮箱账号", nil)
	case "500001", "500002": // 验证码图片
		codeStr = loginJson.Data.CodeString
		if codeStr == "" {
			api.ServerError(c, "未找到codeString", nil)
			return
		}
		api.Error(
			c,
			2,
			"请输入验证码",
			"https://wappass.baidu.com/cgi-bin/genimage?"+codeStr,
		)
	}
}

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	verifyType := c.PostForm("verify_type")
	if verifyType != VerifyEmail && verifyType != VerifyPhone {
		api.BadRequest(c, "验证方式不合法", nil)
		return
	}
	api.Success(c, bc.SendCodeToUser(verifyType, loginJson.Data.Token), nil)
}

// VerifyCode 验证
func VerifyCode(c *gin.Context) {
	code := c.PostForm("code")
	verifyType := c.PostForm("verify_type")
	lj := bc.VerifyCode(verifyType, loginJson.Data.Token, code, loginJson.Data.U)
	if lj.ErrInfo.No != "0" {
		api.ServerError(c, lj.ErrInfo.Msg, nil)
		return
	}
	setUser(c, lj)
}

// setUser 设置用户账号
func setUser(c *gin.Context, lj *baidulogin.LoginJSON) {
	baidu, err := pcsconfig.Config.SetupUserByBDUSS(
		lj.Data.BDUSS,
		lj.Data.PToken,
		lj.Data.SToken,
	)
	if err != nil {
		api.ServerError(c, "登录失败", err)
		return
	}
	api.Success(c, baidu.Name+"登录成功", nil)
}
