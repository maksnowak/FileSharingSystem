package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"file-transfer/models"
)

// uploadFile godoc
//
// @Summary Upload a file
// @Description Upload file with path and content
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param metadata formData string true "JSON metadata with path"
// @Param file formData file true "File content"
// @Success 200 {object} models.FileDataResponse
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /files/upload [post]
func (a *App) uploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	req, err := models.ParseUploadRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to parse request: "+err.Error())
		return
	}
	defer req.File.Close()

	// Read file content
	file, err := io.ReadAll(req.File)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to read file: "+err.Error())
		return
	}

	fileData := models.FileData{
		UserID: req.UserID,
		Path:   req.Path,
		Data:   file,
	}

	// Upload to blob storage
	blobURL, err := a.BlobStorage.UploadFile(ctx, fileData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to upload: "+err.Error())
		return
	}

	response := models.FileResponse{
		UserID: req.UserID,
		Path:   req.Path,
		URL:    blobURL,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// downloadFile godoc
//
// @Summary Download a file
// @Description Download a file with path
// @Tags files
// @Accept json
// @Produce octet-stream
// @Param file body models.FileDownloadRequest true "File metadata"
// @Success 200 {file} file "File content"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /files/download [post]
func (a *App) downloadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	f := models.FileDownloadRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	// Download from blob storage
	file, err := a.BlobStorage.DownloadFile(ctx, f.UserID, f.Path)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to download: "+err.Error())
		return
	}

	// Set headers for file download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(file.Path)))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(file.Data)
}
