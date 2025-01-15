package app

import (
	"file-transfer/db"
	"file-transfer/models"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFileIntegrationTests(t *testing.T) {
	ts := RunTestApp(t)
	defer ts.Close()

	t.Run("it should return 200 when health is ok", func(t *testing.T) {
		resp, err := http.Get("localhost:8080/health")

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})
}
