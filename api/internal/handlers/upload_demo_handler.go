package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
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

	if err := validateDemoFile(fileHeader); err != nil {
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

func validateDemoFile(fileHeader *multipart.FileHeader) error {
	if filepath.Ext(fileHeader.Filename) != ".dem" {
		return errors.New("only .dem files are allowed")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	pr, pw := io.Pipe()
	tee := io.TeeReader(file, pw)

	go func() {
		_, _ = io.Copy(io.Discard, tee)
		pw.Close()
	}()

	parser := demoinfocs.NewParser(pr)

	_, err = parser.ParseNextFrame()
	if err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("file is not a valid CS2 demo: %w", err)
	}

	return nil
}
