package handler

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plainText := []byte("this is testing, this is testing... 123!")
	password := "123"
	cipherText, err := encrypt(plainText, password)
	if err != nil {
		t.Errorf("error during encryption: expected nil; got %v\n", err)
	}

	if slices.Equal(plainText, cipherText) {
		t.Errorf("ciphertext (%v) is the same as plaintext (%v)\n", cipherText, plainText)
	}
}

func TestCallEncrypt(t *testing.T) {
	var reqBody bytes.Buffer
	w := multipart.NewWriter(&reqBody)
	filePart, err := w.CreateFormFile("file", "file")
	if err != nil {
		t.Fatalf("internal error in a test; cannot create form file")
	}
	_, err = io.WriteString(filePart, "I am a byte string")
	if err != nil {
		t.Fatalf("internal error in a test; cannot write to filePart")
	}

	err = w.WriteField("password-hash", "123")
	if err != nil {
		t.Fatalf("internal error in a test; cannot create password-hash field")
	}
	_ = w.Close()

	var resp *http.Response
	{
		req, _ := http.NewRequest(http.MethodPost, "encrypt", &reqBody)
		req.Header.Set("Content-Type", w.FormDataContentType())
		rec := httptest.NewRecorder()
		Encrypt(rec, req)
		resp = rec.Result()
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, resp.StatusCode)
	}
}

func TestCallEncryptNoFile(t *testing.T) {
	var reqBody bytes.Buffer
	w := multipart.NewWriter(&reqBody)

	err := w.WriteField("password-hash", "123")
	if err != nil {
		t.Fatalf("internal error in a test; cannot create password-hash field")
	}
	_ = w.Close()

	var resp *http.Response
	{
		req, _ := http.NewRequest(http.MethodPost, "encrypt", &reqBody)
		req.Header.Set("Content-Type", w.FormDataContentType())
		rec := httptest.NewRecorder()
		Encrypt(rec, req)
		resp = rec.Result()
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, resp.StatusCode)
	}
}

func TestCallEncryptNoPasswordHash(t *testing.T) {
	var reqBody bytes.Buffer
	w := multipart.NewWriter(&reqBody)
	filePart, err := w.CreateFormFile("file", "file")
	if err != nil {
		t.Fatalf("internal error in a test; cannot create form file")
	}
	_, err = io.WriteString(filePart, "I am a byte string")
	if err != nil {
		t.Fatalf("internal error in a test; cannot write to filePart")
	}
	_ = w.Close()

	var resp *http.Response
	{
		req, _ := http.NewRequest(http.MethodPost, "encrypt", &reqBody)
		req.Header.Set("Content-Type", w.FormDataContentType())
		rec := httptest.NewRecorder()
		Encrypt(rec, req)
		resp = rec.Result()
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, resp.StatusCode)
	}
}
