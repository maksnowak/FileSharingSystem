package app

import (
	"context"
	"encoding/json"
	// "file-transfer/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router          *mux.Router
	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
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
	serv := &http.Server{Addr: addr, Handler: a.Router}

	fmt.Printf("http: Listening on %v\n", addr)
	if err := serv.ListenAndServe(); err != nil {
		fmt.Print(err)
	}

	// a.MongoCollection, a.MongoClient = db.InitMongo(ctx)
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
