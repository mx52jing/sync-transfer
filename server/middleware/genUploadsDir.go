package middleware

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mx52jing/sync-transfer/server/utils"
)

func GenUploadsDir(ctx *gin.Context) {
	// 获取当前应用的可执行文件路径
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("os.Executable() err=>", err)
		ctx.JSON(http.StatusOK, gin.H{ "code": utils.FAILED_CODE, "message": "获取可执行文件路径失败"})
		ctx.Abort()
	}
	// 获取可执行文件所在的目录
	execPathDir := filepath.Dir(execPath)
	// 拼接 uploads 文件夹路径，将uploads文件夹，放到可执行文件的目录中
	uploadsDirPath := filepath.Join(execPathDir, "uploads")
	// 查看目录是否存在，如果不存在就创建
	_, statErr := os.Stat(uploadsDirPath) 
	if os.IsNotExist(statErr) {
		// 目录不存在，创建它
		err := os.MkdirAll(uploadsDirPath, 0755)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{ "code": utils.FAILED_CODE, "message": "创建存储目录失败"})
			ctx.Abort()
			return
		}
	}else if statErr != nil {
		// 发生其他错误
		ctx.JSON(http.StatusOK, gin.H{ "code": utils.FAILED_CODE, "message": err.Error()})
		ctx.Abort()
		return
	}
	// 生成一个随机字符串
	randomStr := uuid.New().String()
	ctx.Set("randomStr", randomStr)
	ctx.Set("uploadsDirPath", uploadsDirPath)
	// 执行下一个中间件
	ctx.Next()
}