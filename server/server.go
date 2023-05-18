package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mx52jing/sync-transfer/server/controller"
	"github.com/mx52jing/sync-transfer/server/middleware"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run() {
		router := gin.Default()
		staticFiles, _ := fs.Sub(FS, "frontend/dist")
		router.StaticFS("/static", http.FS(staticFiles))
		v1Group := router.Group("/api/v1")
		{
			v1Group.GET("/addresses", controller.AddressController)
			v1Group.GET("/qrcodes", controller.QrcodesController)
			v1Group.GET("/uploads/:path", controller.UploadsController)
			v1Group.POST("/texts", middleware.GenUploadsDir, controller.TextsController)
		}
		// 处理 404 访问不存在的路由时的情况
		router.NoRoute(func(ctx *gin.Context) {
			path := ctx.Request.URL.Path
			if strings.HasPrefix(path, "/static/") {
				reader, err := staticFiles.Open("index.html")
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				stat, err := reader.Stat()
				if err != nil {
					log.Fatal(err)
				}
				ctx.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil)
			} else {
				ctx.Status(http.StatusNotFound)
			}
			// ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9527/static/index.html")
		})
		router.Run(":9527")
}