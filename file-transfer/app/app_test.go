package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"file-transfer/db"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func RunTestApp(t *testing.T) *App {
	log.Println("Setting up database...")
	cmd := exec.Command(
		"docker", "run",
		"--rm", "-d", "-p", "27017:27017",
		"--name", "testDB",
		"-e", "MONGO_INITDB_ROOT_USERNAME=root", "-e", "MONGO_INITDB_ROOT_PASSWORD=example",
		"mongo",
	)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}

	ctx := context.Background()
	a := &App{}

	a.Router = mux.NewRouter().StrictSlash(true)
	a.Logger = log.New(os.Stdout, "server: ", log.Flags())

	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	a.MongoClient, err = mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = a.MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	a.MongoCollection = a.MongoClient.Database("files").Collection("Files")

	a.BlobStorage, _ = db.InitLocalBlobStorage("files")

	a.initRoutes()

	a.Server = &http.Server{Addr: ":8080", Handler: a.Router}

	return a
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

func CleanDatabase(t *testing.T, collection *mongo.Collection) {
	ctx := context.TODO()
	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		t.Error(err)
	}
}

func TestHealthCheck(t *testing.T) {
	a := RunTestApp(t)
	defer a.Close(context.Background())
	defer KillDatabase()

	t.Run("it should return 200", func(t *testing.T) {
		server := httptest.NewServer(a.Router)
		resp, err := http.Get(server.URL + "/health")
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})
}
