package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mx52jing/sync-transfer/server/utils"
	"github.com/skip2/go-qrcode"
)

func QrcodesController(ctx *gin.Context) {
	content := ctx.Query("content")
	if content == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": utils.FAILED_CODE, "message": "生成图片失败" })
		return
	}
	ctx.Data(http.StatusOK, "image/png", png)
}