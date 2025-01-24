package app

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"file-transfer/models"

	"github.com/stretchr/testify/assert"
)

func TestLocalStorage(t *testing.T) {
	a := RunTestApp(t)
	defer a.Close(context.Background())
	defer KillDatabase()

	t.Run("it should upload and download file", func(t *testing.T) {
		server := httptest.NewServer(a.Server.Handler)
		expected := models.FileResponse{
			UserID: "123",
			Path:   "test.txt",
			URL:    "http://localhost:8080/files/123/test.txt",
			Size:   13,
		}

		fileData := "Hello, World!"
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", "test.txt")
		assert.NoError(t, err)

		_, err = io.Copy(part, strings.NewReader(fileData))
		assert.NoError(t, err)

		err = writer.WriteField("metadata", `{"user_id":"123","path":"test.txt"}`)
		assert.NoError(t, err)

		writer.Close()

		resp, err := http.Post(server.URL+"/upload", writer.FormDataContentType(), body)
		assert.NoError(t, err)

		var actual models.FileResponse
		err = json.NewDecoder(resp.Body).Decode(&actual)
		assert.NoError(t, err)

		assert.Equal(t, expected.UserID, actual.UserID)
		assert.Equal(t, expected.Path, actual.Path)
		assert.Equal(t, expected.URL, actual.URL)

		resp, err = http.Post(server.URL+"/download", "application/json", strings.NewReader(`{"user_id":"123","path":"test.txt"}`))
		assert.NoError(t, err)

    respFile, err := io.ReadAll(resp.Body)
    assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "attachment; filename=test.txt", resp.Header.Get("Content-Disposition"))
		assert.Equal(t, "application/octet-stream", resp.Header.Get("Content-Type"))
		assert.Equal(t, fileData, string(respFile))
	})
}
