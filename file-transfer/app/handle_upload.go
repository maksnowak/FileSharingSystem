package app

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"file-transfer/models"
)

// uploadFile godoc
//
//  @Summary    Upload a file
//  @Description    Upload a file to the server
//  @Tags    files
//  @Accept    json
//  @Produce    json
//  @Param    file    formData    file    true    "File to upload"
//  @Param    path    body    string    true    "Path to store the file"
//  @Success    200    {object}    models.FileDataResponse    "Uploaded file object"
//  @Failure    400    {string}    string    "Invalid request payload"
//  @Failure    500    {string}    string    "Internal server error"
//  @Router    /files/upload [post]
func (a *App) uploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	f := models.FileDataRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload:"+err.Error())
		return
	}

	// Parse multipart form with 32MB max memory
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		respondWithError(w, http.StatusBadRequest, "File too large or invalid form")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid file")
		return
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error reading file")
		return
	}

	fileData := models.FileData{
		Path: f.Path,
		Data: buffer,
	}

	URL, err := a.BlobStorage.UploadFile(ctx, fileData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, models.FileDataResponse{
		Path: f.Path,
		URL:  URL,
	})
}

// downloadFile godoc
//
//  @Summary    Download a file
//  @Description    Download a file from the server
//  @Tags    files
//  @Accept    json
//  @Produce    json
//  @Param    path    body    string    true    "Path to the file"
//  @Success    200    {object}    models.FileDataResponse    "Downloaded file object"
//  @Failure    400    {string}    string    "Invalid request payload"
//  @Failure    500    {string}    string    "Internal server error"
//  @Router    /files/download [post]
func (a *App) downloadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	f := models.FileDataRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload:"+err.Error())
		return
	}

	file, err := a.BlobStorage.DownloadFile(ctx, f.Path)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+f.Path)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(file.Data)
}
