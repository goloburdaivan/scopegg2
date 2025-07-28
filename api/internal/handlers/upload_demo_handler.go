package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"scopegg2/internal/dto"
)

type DemoService interface {
	UploadAndNotify(ctx context.Context, file io.Reader, filename string) (*dto.FileUploadedResult, error)
}

type UploadDemoHandler struct {
	demoService DemoService
}

func NewUploadDemoHandler(demoService DemoService) *UploadDemoHandler {
	return &UploadDemoHandler{demoService: demoService}
}

func (h *UploadDemoHandler) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := fileHeader.Open()
	defer file.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	_, err = h.demoService.UploadAndNotify(ctx, file, fileHeader.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
