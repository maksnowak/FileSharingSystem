package app

import (
	"bytes"
	"context"
	"encoding/json"
	"file-transfer/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileIntegrationTests(t *testing.T) {
	a := RunTestApp(t)
	defer a.Close(context.Background())
	defer KillDatabase()

	t.Run("it should create and return file", func(t *testing.T) {
		server := httptest.NewServer(a.Server.Handler)
		expected := models.File{
			FileName: "test.txt",
			UserID:   "123",
			Path:     "path/test.txt",
		}

		// Create file
		body, err := json.Marshal(expected)
		assert.NoError(t, err)

		reader := bytes.NewReader(body)
		resp, err := http.Post(server.URL+"/file", "application/json", reader)
		assert.NoError(t, err)

		var actual models.File
		err = json.NewDecoder(resp.Body).Decode(&actual)
		assert.NoError(t, err)

		assert.Equal(t, expected.FileName, actual.FileName)
		assert.Equal(t, expected.UserID, actual.UserID)
		assert.Equal(t, expected.Path, actual.Path)
	})

	t.Run("it should return file", func(t *testing.T) {
		defer CleanDatabase(t, a.MongoCollection)
		server := httptest.NewServer(a.Server.Handler)
		expected := models.File{
			FileName: "test.txt",
			UserID:   "123",
			Path:     "path/test.txt",
		}

		// Create file
		a.MongoCollection.InsertOne(context.TODO(), expected)

		// Get file
		resp, err := http.Get(server.URL + "/file/" + expected.FileID.Hex())
		assert.NoError(t, err)

		var actual models.File
		err = json.NewDecoder(resp.Body).Decode(&actual)
		assert.NoError(t, err)
	})
}
