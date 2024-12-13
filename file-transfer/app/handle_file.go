package app

import (
	"context"
	"encoding/json"
	"net/http"

	"file-transfer/db"
	"file-transfer/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *App) createFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	f := models.File{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload:"+err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.CreateFile(&ctx, a.MongoCollection, f); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, f)
}

// getAllFiles	godoc
//
//	@Summary		Retrieve all files
//	@Description	Retrieve information about all existing files
//	@Tags			files
//	@Produce		json
//	@Success		200	{array}		models.File		"Every existing file"
//	@Router			/files [get]
func (a *App) getAllFiles(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	files, err := db.GetAllFiles(&ctx, a.MongoCollection)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, files)
}

func (a *App) getFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["file_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid file ID")
		return
	}

	f := models.File{FileID: id}
	f, err = db.GetFile(&ctx, a.MongoCollection, f)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, f)
}

func (a *App) updateFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["file_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid file ID")
		return
	}

	f := models.File{FileID: id}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload:"+err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.UpdateFile(&ctx, a.MongoCollection, f); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, f)
}

func (a *App) deleteFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["file_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid file ID")
		return
	}

	f := models.File{FileID: id}
	if err := db.DeleteFile(&ctx, a.MongoCollection, f); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}