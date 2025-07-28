package server

import (
	"github.com/gin-gonic/gin"
	"scopegg2/internal/handlers"
)

func NewRouter(uploadHandler *handlers.UploadDemoHandler) *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 1 << 30
	router.POST("/upload-demo", uploadHandler.Upload)
	return router
}
