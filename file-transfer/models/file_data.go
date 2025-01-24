package models

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
)

type FileData struct {
	UserID string `json:"user_id"`
	Path   string `json:"path"`
	Data   []byte `json:"data"`
}

type FileUploadRequest struct {
	UserID string `json:"user_id"`
	Path   string `json:"path"`
	File   multipart.File
}

type FileDownloadRequest struct {
	UserID string `json:"user_id"`
	Path   string `json:"path"`
}

type FileResponse struct {
	UserID string `json:"user_id"`
	Path   string `json:"path"`
	URL    string `json:"url"`
	Size   int64  `json:"size"`
}

func ParseUploadRequest(r *http.Request) (*FileUploadRequest, error) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	metadataStr := r.FormValue("metadata")
	if metadataStr == "" {
		return nil, fmt.Errorf("missing metadata")
	}

	var req FileUploadRequest
	if err := json.Unmarshal([]byte(metadataStr), &req); err != nil {
		return nil, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}

	req.File = file
	return &req, nil
}
