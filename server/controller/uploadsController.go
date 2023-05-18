package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func getUploadsDir() (uploadsDir string) {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		return
	}
	execDir := filepath.Dir(execPath)
	uploadsDir = filepath.Join(execDir, "uploads")
	return 
}

func UploadsController(ctx *gin.Context) {
	if path := ctx.Param("path"); path != "" {
		uploadsDir := getUploadsDir()
		filePath := filepath.Join(uploadsDir, path)
		// ctx.Header("Content-Type", "application/octet-stream")
		// ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.FileAttachment(filePath, path)
		return
	}
	ctx.Status(http.StatusBadRequest)
}