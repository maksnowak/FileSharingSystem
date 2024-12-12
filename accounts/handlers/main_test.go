package handlers

import (
	"accounts/db"
	"accounts/models"
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

type Handler func(w http.ResponseWriter, r *http.Request)

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
	coll := db.GetCollection("users")
	_, _ = coll.DeleteMany(context.Background(), bson.D{})
}

func Call(msg interface{}, method string, target string, fn Handler, router chi.Router) *httptest.ResponseRecorder {
	var req *http.Request
	if msg != nil {
		body, _ := json.Marshal(msg)
		req = httptest.NewRequest(method, target, bytes.NewReader(body))
	} else {
		req = httptest.NewRequest(method, target, nil)
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	// Call the handler
	if router != nil {
		router.ServeHTTP(rec, req)
	} else {
		fn(rec, req)
	}

	return rec
}

func TestMain(m *testing.M) {
	SetupDatabase()
	code := m.Run()
	KillDatabase()
	os.Exit(code)
}

var fixedTime = time.Date(2024, 12, 12, 21, 55, 52, 460000000, time.UTC)

var id1, _ = bson.ObjectIDFromHex("123456789012345678901234")
var u1 = models.User{
	ID:           id1,
	Username:     "Test1",
	Email:        "test1@test.com",
	PasswordHash: "Hash1",
	PasswordSalt: "Salt1",
	CreatedAt:    fixedTime,
	Role:         "admin",
	OwnedFiles:   []string{},
	SharedFiles:  []string{},
}

var id2, _ = bson.ObjectIDFromHex("102938475610293847561029")
var u2 = models.User{
	ID:           id2,
	Username:     "Test2",
	Email:        "test2@test.com",
	PasswordHash: "Hash2",
	PasswordSalt: "Salt2",
	CreatedAt:    fixedTime,
	Role:         "user",
	OwnedFiles:   []string{},
	SharedFiles:  []string{},
}

var id3, _ = bson.ObjectIDFromHex("67599a97ca18148343366a54")
var u3 = models.User{
	ID:           id3,
	Username:     "Test3",
	Email:        "test3@test.com",
	PasswordHash: "Hash3",
	PasswordSalt: "Salt3",
	CreatedAt:    fixedTime,
	Role:         "user",
	OwnedFiles:   []string{},
	SharedFiles:  []string{},
}

var reg = models.Register{
	Username:     "Test1",
	Email:        "test1@test.com",
	PasswordHash: "Hash1",
	PasswordSalt: "Salt1",
	Role:         "admin",
}

var up = models.Update{
	Email:        "Updated@test.com",
	PasswordHash: "UpdatedHash",
	PasswordSalt: "UpdatedSalt",
	OwnedFiles:   []string{"bike"},
	SharedFiles:  []string{"dog"},
}
