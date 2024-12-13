package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router          *mux.Router
	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection
}

func (a *App) Initialize() {
	ctx := context.TODO()

	a.Router = mux.NewRouter()
	a.MongoCollection, a.MongoClient = initMongo(&ctx)
  a.initRoutes()
}

func (a *App) Run(addr string) {
  http.Handle("/", a.Router)
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		panic(fmt.Sprintf("Could not start server: %s", err))
	}
}

func (a *App) initRoutes() {
  a.Router.HandleFunc("/health", a.healthCheck).Methods(http.MethodGet)

	a.Router.HandleFunc("/file", a.createFile).Methods(http.MethodPost)
	a.Router.HandleFunc("/files", a.getAllFiles).Methods(http.MethodGet)
	a.Router.HandleFunc("/file/{file_id}", a.getFile).Methods(http.MethodGet)
	a.Router.HandleFunc("/file/{file_id}", a.updateFile).Methods(http.MethodPut)
	a.Router.HandleFunc("/file/{file_id}", a.deleteFile).Methods(http.MethodDelete)
}

func (a *App) Close() {
  ctx := context.TODO()
  if err := a.MongoClient.Disconnect(ctx); err != nil {
    panic(fmt.Sprintf("Could not close connection: %s", err))
  }
}

func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
  respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (a *App) createFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	f := File{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload:" + err.Error())
		return
	}
	defer r.Body.Close()

	if err := f.createFile(&ctx, a.MongoCollection); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, f)
}

func (a *App) getAllFiles(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	files, err := getAllFiles(&ctx, a.MongoCollection)
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

	f := File{FileID: id}
	if err := f.getFile(&ctx, a.MongoCollection); err != nil {
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

	f := File{FileID: id}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload:" + err.Error())
		return
	}
	defer r.Body.Close()

	if err := f.updateFile(&ctx, a.MongoCollection); err != nil {
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

	f := File{FileID: id}
	if err := f.deleteFile(&ctx, a.MongoCollection); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
