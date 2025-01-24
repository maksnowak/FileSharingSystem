package app

import (
	"bytes"
	"context"
	"encoding/json"
	"file-transfer/models"
	"fmt"
	"io"
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
		defer CleanDatabase(t, a.MongoCollection)
		server := httptest.NewServer(a.Server.Handler)
		expected := models.File{
			FileName: "test.txt",
			UserID:   "123",
			Path:     "path/test.txt",
		}

		body, err := json.Marshal(expected)
		assert.NoError(t, err)

		reader := bytes.NewReader(body)
		resp, err := http.Post(server.URL+"/file", "application/json", reader)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var actual models.File
		err = json.NewDecoder(resp.Body).Decode(&actual)
		assert.NoError(t, err)

		assert.Equal(t, expected.FileName, actual.FileName)
		assert.Equal(t, expected.UserID, actual.UserID)
		assert.Equal(t, expected.Path, actual.Path)
	})

	t.Run("it should return all files", func(t *testing.T) {
		defer CleanDatabase(t, a.MongoCollection)
		server := httptest.NewServer(a.Server.Handler)
		file := models.File{
			FileName: "test.txt",
			UserID:   "123",
			Path:     "path/test.txt",
		}

		// Create files
		for i := 0; i < 3; i++ {
			file.FileName = fmt.Sprintf("test%d.txt", i)
			file.Path = fmt.Sprintf("path/test%d.txt", i)
			body, err := json.Marshal(file)
			assert.NoError(t, err)

			reader := bytes.NewReader(body)
			resp, err := http.Post(server.URL+"/file", "application/json", reader)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusCreated, resp.StatusCode)
		}

		// Get all files
		resp, err := http.Get(server.URL + "/files")
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var actual []models.File
		data, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		err = json.Unmarshal(data, &actual)
		assert.NoError(t, err)

		assert.Len(t, actual, 3)
	})

	t.Run("it should modify file", func(t *testing.T) {
		defer CleanDatabase(t, a.MongoCollection)
		server := httptest.NewServer(a.Server.Handler)
		initial := models.File{
			FileName: "test.txt",
			UserID:   "123",
			Path:     "path/test.txt",
		}
		expected := models.File{
			FileName: "test.txt",
			UserID:   "123",
			Path:     "path/test.txt",
			BlobURL:  "http://example.com",
		}

		// Create file
		body, err := json.Marshal(initial)
		assert.NoError(t, err)

		reader := bytes.NewReader(body)
		resp, err := http.Post(server.URL+"/file", "application/json", reader)
		assert.NoError(t, err)

		var actual models.File
		err = json.NewDecoder(resp.Body).Decode(&actual)
		assert.NoError(t, err)

		// Update file
		body, err = json.Marshal(expected)
		assert.NoError(t, err)
		req, err := http.NewRequest("PUT", server.URL+"/file/"+actual.FileID.Hex(), bytes.NewReader(body))
		assert.NoError(t, err)

    req.Header.Set("Content-Type", "application/json")
    resp, err = http.DefaultClient.Do(req)
    assert.NoError(t, err)

    assert.Equal(t, http.StatusOK, resp.StatusCode)

    body, err = io.ReadAll(resp.Body)
    assert.NoError(t, err)

    err = json.Unmarshal(body, &actual)
    assert.NoError(t, err)

		assert.Equal(t, expected.FileName, actual.FileName)
		assert.Equal(t, expected.UserID, actual.UserID)
		assert.Equal(t, expected.Path, actual.Path)
		assert.Equal(t, expected.BlobURL, actual.BlobURL)
	})

	t.Run("it should delete file", func(t *testing.T) {
		defer CleanDatabase(t, a.MongoCollection)
		server := httptest.NewServer(a.Server.Handler)
		file := models.File{
			FileName: "test.txt",
			UserID:   "123",
			Path:     "path/test.txt",
		}

		// Create file
		body, err := json.Marshal(file)
		assert.NoError(t, err)

		reader := bytes.NewReader(body)
		resp, err := http.Post(server.URL+"/file", "application/json", reader)
		assert.NoError(t, err)

		var actual models.File
		err = json.NewDecoder(resp.Body).Decode(&actual)
		assert.NoError(t, err)

		// Delete file
		req, err := http.NewRequest(http.MethodDelete, server.URL+"/file/"+actual.FileID.Hex(), nil)
		assert.NoError(t, err)

		resp, err = http.DefaultClient.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Get file
		resp, err = http.Get(server.URL + "/file/" + actual.FileID.Hex())
		assert.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("it should return all user files", func(t *testing.T) {
		defer CleanDatabase(t, a.MongoCollection)
		server := httptest.NewServer(a.Server.Handler)
		file := models.File{
			FileName: "test.txt",
			UserID:   "123",
			Path:     "path/test.txt",
		}
		file2 := models.File{
			FileName: "test2.txt",
			UserID:   "456",
			Path:     "path/test2.txt",
		}

		// Create files
		reqBody, err := json.Marshal(file)
		assert.NoError(t, err)

		reader := bytes.NewReader(reqBody)
		resp, err := http.Post(server.URL+"/file", "application/json", reader)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		reqBody, err = json.Marshal(file2)
		assert.NoError(t, err)

		reader = bytes.NewReader(reqBody)
		resp, err = http.Post(server.URL+"/file", "application/json", reader)

		// Get all files
		resp, err = http.Get(server.URL + "/file/user/123")
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		actual := []models.File{}
		data, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &actual)
		assert.NoError(t, err)

		assert.Len(t, actual, 1)
	})
}
