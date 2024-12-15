package app

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "file-transfer/db"
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
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Logger = log.New(os.Stdout, "server: ", log.Flags())

	logMiddleware := NewLogMiddleware(a.Logger)
	a.Router.Use(logMiddleware.Func())

	a.initRoutes()
	a.initDocs()
}

//	@title			File transfer API
//	@version		0.2
//	@description	Webserver providing saving and retrieval of files from MongoDB

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// @BasePath	/
func (a *App) Run(ctx *context.Context, addr string) {
	a.Server = &http.Server{Addr: addr, Handler: a.Router}
	// a.MongoCollection, a.MongoClient = db.InitMongo(ctx)

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		a.Logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		a.Server.SetKeepAlivesEnabled(false)
		if err := a.Server.Shutdown(ctx); err != nil {
			a.Logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	a.Logger.Println("Server is ready to handle requests at :8080")
	if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.Logger.Fatalf("Could not listen on :8080: %v\n", err)
	}

	<-done
	a.Logger.Println("Server stopped")
}

func (a *App) Close(ctx context.Context) error {
	// if err := a.MongoClient.Disconnect(ctx); err != nil {
	// 	return err
	// }

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
