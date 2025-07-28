package dto

import "time"

type FileUploadedResult struct {
	Filename   string    `json:"filename"`
	Path       string    `json:"path"`
	UploadedAt time.Time `json:"uploaded_at"`
}
