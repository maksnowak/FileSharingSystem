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

func TestUploadLocalStorage(t *testing.T) {
	a := RunTestApp(t)
	defer a.Close(context.Background())
	defer KillDatabase()

	t.Run("it should upload and return file", func(t *testing.T) {
		server := httptest.NewServer(a.Server.Handler)
		expected := models.FileDataRequest{
			Path: "test.txt",
		}

		fileData := "Hello, World!"
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", "test.txt")
		assert.NoError(t, err)

		_, err = io.Copy(part, strings.NewReader(fileData))
		assert.NoError(t, err)

		err = writer.WriteField("path", expected.Path)
		assert.NoError(t, err)

		writer.Close()

		resp, err := http.Post(server.URL+"/files/upload", writer.FormDataContentType(), body)
		assert.NoError(t, err)

		var actual models.FileDataResponse
		err = json.NewDecoder(resp.Body).Decode(&actual)
		assert.NoError(t, err)

		assert.Equal(t, expected.Path, actual.Path)
	})
}

func TestDownloadLocalStorage(t *testing.T) {
	a := RunTestApp(t)
	defer a.Close(context.Background())
	defer KillDatabase()

	t.Run("it should upload and download file", func(t *testing.T) {
		server := httptest.NewServer(a.Server.Handler)
		expected := models.FileDataRequest{
			Path: "test.txt",
		}

		fileData := "Hello, World!"
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", "test.txt")
		assert.NoError(t, err)

		_, err = io.Copy(part, strings.NewReader(fileData))
		assert.NoError(t, err)

		err = writer.WriteField("path", expected.Path)
		assert.NoError(t, err)

		writer.Close()

		resp, err := http.Post(server.URL+"/files/upload", writer.FormDataContentType(), body)
		assert.NoError(t, err)

		var actual models.FileDataResponse
		err = json.NewDecoder(resp.Body).Decode(&actual)

		resp, err = http.Post(server.URL+"/files/download", "application/json", strings.NewReader(`{"path":"test.txt"}`))
		assert.NoError(t, err)

		var actualDownload models.FileDataResponse
		err = json.NewDecoder(resp.Body).Decode(&actualDownload)
		assert.NoError(t, err)

		assert.Equal(t, expected.Path, actualDownload.Path)
		assert.Equal(t, actual.URL, actualDownload.URL)
	})
}
