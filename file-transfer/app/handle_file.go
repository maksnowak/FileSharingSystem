package app

import (
	"context"
	"encoding/json"
	"net/http"

	"file-transfer/db"
	"file-transfer/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// createFile godoc
//
//	@Summary		Create a new file
//	@Description	Create a new file record in the database
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			file	body		models.File	true	"File object to create"
//	@Success		201		{object}	models.File	"Created file object"
//	@Failure		400		{string}	string		"Invalid request payload"
//	@Failure		500		{string}	string		"Internal server error"
//	@Router			/files [post]
func (a *App) createFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	f := models.File{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload:"+err.Error())
		return
	}
	defer r.Body.Close()

	file, err := db.CreateFile(&ctx, a.MongoCollection, f)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, file)
}

// getFile godoc
//
//	@Summary		Retrieve a specific file
//	@Description	Get information about a file by its ID
//	@Tags			files
//	@Produce		json
//	@Param			file_id	path		string		true	"File ID"
//	@Success		200		{object}	models.File	"Retrieved file object"
//	@Failure		400		{string}	string		"Invalid file ID"
//	@Failure		500		{string}	string		"Internal server error"
//	@Router			/files/{file_id} [get]
func (a *App) getFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	vars := mux.Vars(r)
	id, err := bson.ObjectIDFromHex(vars["file_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid file ID")
		return
	}

	f := models.File{FileID: id}
	f, err = db.GetFile(&ctx, a.MongoCollection, f)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			respondWithError(w, http.StatusNotFound, "File not found")
			return
		}
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
//	@Success		200	{array}		models.File	"Every existing file"
//	@Failure		500	{string}	string		"Internal server error"
//	@Router			/files [get]
func (a *App) getAllFiles(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	files, err := db.GetAllFiles(&ctx, a.MongoCollection)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(files) == 0 {
		respondWithError(w, http.StatusNoContent, string(json.RawMessage("[]")))
		return
	}
	respondWithJSON(w, http.StatusOK, files)
}

// getFilesByUser godoc
//
//	@Summary		Retrieve files by user
//	@Description	Retrieve information about all files uploaded by a specific user
//	@Tags			files
//	@Produce		json
//	@Param			user_id	path		string		true	"User ID"
//	@Success		200		{array}		models.File	"Files uploaded by the user"
//	@Failure		500		{string}	string		"Internal server error"
//	@Router			/files/user/{user_id} [get]
func (a *App) getFilesByUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	vars := mux.Vars(r)
	userID := vars["user_id"]

	files, err := db.GetFilesByUserID(&ctx, a.MongoCollection, userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(files) == 0 {
		respondWithError(w, http.StatusNoContent, string(json.RawMessage("[]")))
		return
	}
	respondWithJSON(w, http.StatusOK, files)
}

// updateFile godoc
//
//	@Summary		Update an existing file
//	@Description	Update the details of an existing file by its ID
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			file_id	path		string		true	"File ID"
//	@Param			file	body		models.File	true	"File object with updated information"
//	@Success		200		{object}	models.File	"Updated file object"
//	@Failure		400		{string}	string		"Invalid request payload or file ID"
//	@Failure		500		{string}	string		"Internal server error"
//	@Router			/files/{file_id} [put]
func (a *App) updateFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	vars := mux.Vars(r)
	id, err := bson.ObjectIDFromHex(vars["file_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid file ID")
		return
	}

	f := models.File{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload:"+err.Error())
		return
	}
	defer r.Body.Close()

  f.FileID = id
	f, err = db.UpdateFile(&ctx, a.MongoCollection, f)
  if err != nil {
		if err.Error() == "mongo: no documents in result" {
			respondWithError(w, http.StatusNotFound, "File not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, f)
}

// deleteFile godoc
//
//	@Summary		Delete a specific file
//	@Description	Remove a file from the database by its ID
//	@Tags			files
//	@Produce		json
//	@Param			file_id	path		string				true	"File ID"
//	@Success		200		{object}	map[string]string	"Result: success"
//	@Failure		400		{string}	string				"Invalid file ID"
//	@Failure		500		{string}	string				"Internal server error"
//	@Router			/files/{file_id} [delete]
func (a *App) deleteFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	vars := mux.Vars(r)
	id, err := bson.ObjectIDFromHex(vars["file_id"])
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			respondWithError(w, http.StatusNotFound, "File not found")
			return
		}
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
