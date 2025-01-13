package app

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"file-transfer/db"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router          *mux.Router
	Server          *http.Server
	Logger          *log.Logger
	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection
	BlobStorage     *db.BlobStorage
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Logger = log.New(os.Stdout, "server: ", log.Flags())

	logMiddleware := NewLogMiddleware(a.Logger)
	a.Router.Use(logMiddleware.Func())

	a.initRoutes()
}

func (a *App) Run(ctx *context.Context, addr string) {
	a.Server = &http.Server{Addr: addr, Handler: a.Router}

	a.MongoCollection, a.MongoClient = db.InitMongo(ctx)
	a.BlobStorage, _ = db.InitBlobStorage("files")

	a.Logger.Println("Server is ready to handle requests at :8080")
	if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.Logger.Fatalf("Could not listen on :8080: %v\n", err)
	}
}

func (a *App) Close(ctx context.Context) error {
	a.Logger.Println("Server is shutting down...")

	a.Server.SetKeepAlivesEnabled(false)
	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		return err
	}

	if err := a.MongoClient.Disconnect(ctx); err != nil {
		a.Logger.Fatalf("Could not gracefully shutdown the MongoDB client: %v\n", err)
		return err
	}

	a.Logger.Println("Server stopped")
	return nil
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
