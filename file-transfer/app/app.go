package app

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

func (a *App) Close() {
	ctx := context.TODO()
	if err := a.MongoClient.Disconnect(ctx); err != nil {
		panic(fmt.Sprintf("Could not close connection: %s", err))
	}
}

func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
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
