package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mx52jing/sync-transfer/server/utils"
)

func AddressController(ctx *gin.Context) {
	ipList, err := utils.GetIp(); 
	if err != nil {
		ctx.JSON(http.StatusOK, utils.CommonResponse{ Code: -1, Message: "获取IP失败", Data: err })
	}
	data := utils.CommonResponse{ Code: 0, Message: "获取IP成功", Data: gin.H{ "addresses": ipList }}
	ctx.JSON(http.StatusOK, data)
}