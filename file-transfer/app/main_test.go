package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"

	"file-transfer/db"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RunTestApp(t *testing.T) *App {
	ctx := context.Background()
	a := &App{}

	a.Router = mux.NewRouter().StrictSlash(true)
	a.Logger = log.New(os.Stdout, "server: ", log.Flags())

	a.MongoCollection, a.MongoClient = db.InitMongo(&ctx)

	a.BlobStorage, _ = db.InitLocalBlobStorage("files")

	a.initRoutes()

	return a
}

func SetupDatabase() {
	log.Println("Setting up database...")
	cmd := exec.Command("docker", "run", "--rm", "-d", "-p", "27017:27017", "--name", "testDB", "mongo")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
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

func CleanDatabase(a *App) {
	_, _ = a.MongoCollection.DeleteMany(context.Background(), bson.D{})
}

func TestHealthCheck(t *testing.T) {
	a := RunTestApp(t)
	defer a.Close(context.Background())

	t.Run("it should return 200 OK", func(t *testing.T) {
		resp, err := http.Get("http://localhost:8080/health")

		if err != nil {
			t.Fatalf("Could not send GET request: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})
}

func TestMain(m *testing.M) {
	SetupDatabase()
	code := m.Run()
	KillDatabase()
	os.Exit(code)
}
