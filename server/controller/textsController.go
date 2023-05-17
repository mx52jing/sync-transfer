package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mx52jing/sync-transfer/server/utils"
)

func TextsController(ctx *gin.Context) {
	randomStr, isStrExist := ctx.Get("randomStr"); 
	dirPath, _ := ctx.Get("uploadsDirPath")
	if !isStrExist {
		// 如果uuid生成失败就重新生成
		randomStr = uuid.New().String()
	}
	dir, _ := dirPath.(string)
	// 类型断言，将randomStr转换为string类型
	str, _ := randomStr.(string)
	// 拼接文件路径
	fullPath := filepath.Join(dir, str + ".txt")
	var texts utils.CommonUploadRequest
	if err := ctx.ShouldBindJSON(&texts); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.CommonResponse{Code: utils.FAILED_CODE, Message: "失败", Data: err.Error()})
		return
	}
	// 写入文件
	err := os.WriteFile(fullPath, []byte(texts.Raw), 0666)
	if err != nil {
		fmt.Println("os.WriteFile", err)
		return
	}	
	ctx.JSON(http.StatusOK, gin.H{"code": utils.SUCCESS_CODE, "data": gin.H{"url": fullPath }})
}