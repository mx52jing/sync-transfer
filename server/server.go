package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mx52jing/sync-transfer/server/controller"
)

func Run() {
		router := gin.Default()
		router.Static("/static", "frontend/dist")
		v1Group := router.Group("/api/v1")
		{
			v1Group.GET("/addresses", controller.AddressController)
		}
		// 处理 404 访问不存在的路由时的情况
		router.NoRoute(func(ctx *gin.Context) {
			ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9527/static/index.html")
		})
		router.Run(":9527")
}