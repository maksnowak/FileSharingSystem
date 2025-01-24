package app

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"file-transfer/db"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Router          *mux.Router
	Server          *http.Server
	Logger          *log.Logger
	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection
	BlobStorage     db.BlobStorage
}

func (a *App) Initialize(ctx *context.Context) {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Logger = log.New(os.Stdout, "server: ", log.Flags())

	a.MongoCollection, a.MongoClient = db.InitMongo(ctx)

	_ = godotenv.Load("./.env")

	storageType := os.Getenv("STORAGE_TYPE")

	if storageType == "local" {
		path := os.Getenv("LOCAL_STORAGE_PATH")
		a.BlobStorage, _ = db.InitLocalBlobStorage(path)
	} else {
		a.BlobStorage, _ = db.InitAzureBlobStorage("files")
	}

	logMiddleware := NewLogMiddleware(a.Logger)
  corsMiddleware := NewCorsMiddleware(
    []string{"http://localhost:8080"}, // Allowed origins
    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},     // Allowed methods
    []string{"Content-Type", "Authorization"},              // Allowed headers
    true,
  )
	a.Router.Use(logMiddleware.Func())
  a.Router.Use(corsMiddleware.Func())

	a.initRoutes()
}

func (a *App) Run(ctx *context.Context, addr string) {
	a.Server = &http.Server{Addr: addr, Handler: a.Router}

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
