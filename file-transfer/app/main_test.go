package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
	"time"
)

func runTestApp(t *testing.T) *App {
	a := &App{}

	a.Router = mux.NewRouter().StrictSlash(true)
	a.Logger = log.New(os.Stdout, "server: ", log.Flags())

	a.MongoCollection, a.MongoClient = db.InitMongo(ctx)

	a.BlobStorage, _ = db.InitLocalBlobStorage("files")

	a.initRoutes()
}

func SetupDatabase() {
	log.Println("Setting up database...")
	cmd := exec.Command("docker", "run", "--rm", "-d", "-p", "27017:27017", "--name", "testDB", "mongo")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	db.Client, err = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	} else {
		log.Println("Database set up successfully")
	}
}

func KillDatabase() {
	log.Println("Killing database...")
	cmd := exec.Command("docker", "kill", "testDB")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	} else {
		log.Println("Database killed")
	}
}

func CleanDatabase() {
	coll := db.GetCollection("files")
	_, _ = coll.DeleteMany(context.Background(), bson.D{})
}

func TestMain(m *testing.M) {
	SetupDatabase()
	code := m.Run()
	KillDatabase()
	os.Exit(code)
}
