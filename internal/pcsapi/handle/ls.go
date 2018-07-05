package handle

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/iikira/BaiduPCS-Go/baidupcs"
	"github.com/iikira/BaiduPCS-Go/internal/pcsapi/api"
	"github.com/iikira/BaiduPCS-Go/internal/pcsconfig"
)

func List(c *gin.Context) {
	order := c.Query("order")
	orderType := c.Query("order_type")
	path := c.Query("path")
	orderOptions := &baidupcs.OrderOptions{}
	switch order {
	case "asc":
		orderOptions.Order = baidupcs.OrderAsc
	case "desc":
		orderOptions.Order = baidupcs.OrderDesc
	default:
		orderOptions.Order = baidupcs.OrderAsc
	}
	switch orderType {
	case "time":
		orderOptions.By = baidupcs.OrderByTime
	case "name":
		orderOptions.By = baidupcs.OrderByName
	case "size":
		orderOptions.By = baidupcs.OrderBySize
	default:
		orderOptions.By = baidupcs.OrderByName
	}
	body, err := pcsconfig.Config.ActiveUserBaiduPCS().PrepareFilesDirectoriesList(path, orderOptions)
	if err != nil {
		api.ServerError(c, "获取失败", err)
		return
	}
	b, e := ioutil.ReadAll(body)
	if e != nil {
		api.ServerError(c, "获取失败", err)
		return
	}
	var files map[string]interface{}
	json.Unmarshal(b, &files)
	if msg, ok := files["error_msg"]; ok {
		api.ServerError(c, msg.(string), files)
		return
	}
	if data, ok := files["list"]; ok {
		api.Success(c, "获取成功", data)
		return
	}
	api.Success(c, "没有获取到数据", files)
}
