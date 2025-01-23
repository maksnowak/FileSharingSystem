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

func TestDecrypt(t *testing.T) {
	plainText := []byte("this is testing, this is testing... 123!")
	password := "123"
	cipherText, err := encrypt(plainText, password)
	if err != nil {
		t.Errorf("error during encryption: expected nil; got %v\n", err)
	}

	plainText2, err := decrypt(cipherText, password)
	if err != nil {
		t.Errorf("error during decryption: expected nil; got %v\n", err)
	}

	if !slices.Equal(plainText, plainText2) {
		t.Errorf("plaintext after decryption (%v) does not match the original (%v)\n", plainText2, plainText)
	}
}

func TestDecryptBadPassword(t *testing.T) {
	plainText := []byte("this is testing, this is testing... 123!")
	cipherText, err := encrypt(plainText, "123")
	if err != nil {
		t.Errorf("error during encryption: expected nil; got %v\n", err)
	}

	_, err = decrypt(cipherText, "456")
	if err == nil {
		t.Errorf("decryption with bad password should fail")
	}
}

func TestCallDecrypt(t *testing.T) {
	var reqBody bytes.Buffer
	w := multipart.NewWriter(&reqBody)
	filePart, err := w.CreateFormFile("file", "file")
	if err != nil {
		t.Fatalf("internal error in a test; cannot create form file")
	}
	plainText := []byte("I am a byte string")
	password := "123"
	cipherText, _ := encrypt(plainText, password)
	_, err = filePart.Write(cipherText)
	if err != nil {
		t.Fatalf("internal error in a test; cannot write to filePart")
	}

	err = w.WriteField("password-hash", password)
	if err != nil {
		t.Fatalf("internal error in a test; cannot create password-hash field")
	}
	_ = w.Close()

	var resp *http.Response
	{
		req, _ := http.NewRequest(http.MethodPost, "decrypt", &reqBody)
		req.Header.Set("Content-Type", w.FormDataContentType())
		rec := httptest.NewRecorder()
		Decrypt(rec, req)
		resp = rec.Result()
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %v; got %v\n", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading from response body %v\n", err)
	}

	if !slices.Equal(plainText, body) {
		t.Fatalf("expected %v; got %v\n", string(plainText), string(body))
	}
}

func TestCallDecryptBadPassword(t *testing.T) {
	var reqBody bytes.Buffer
	w := multipart.NewWriter(&reqBody)
	filePart, err := w.CreateFormFile("file", "file")
	if err != nil {
		t.Fatalf("internal error in a test; cannot create form file")
	}
	plainText := []byte("I am a byte string")
	cipherText, _ := encrypt(plainText, "123")
	_, err = filePart.Write(cipherText)
	if err != nil {
		t.Fatalf("internal error in a test; cannot write to filePart")
	}

	err = w.WriteField("password-hash", "456")
	if err != nil {
		t.Fatalf("internal error in a test; cannot create password-hash field")
	}
	_ = w.Close()

	var resp *http.Response
	{
		req, _ := http.NewRequest(http.MethodPost, "decrypt", &reqBody)
		req.Header.Set("Content-Type", w.FormDataContentType())
		rec := httptest.NewRecorder()
		Decrypt(rec, req)
		resp = rec.Result()
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("decryption with bad password should fail\n")
	}
}

func TestCallDecryptNoFile(t *testing.T) {
	var reqBody bytes.Buffer
	w := multipart.NewWriter(&reqBody)
	err := w.WriteField("password-hash", "123")
	if err != nil {
		t.Fatalf("internal error in a test; cannot create password-hash field")
	}
	_ = w.Close()

	var resp *http.Response
	{
		req, _ := http.NewRequest(http.MethodPost, "decrypt", &reqBody)
		req.Header.Set("Content-Type", w.FormDataContentType())
		rec := httptest.NewRecorder()
		Decrypt(rec, req)
		resp = rec.Result()
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected %v; got %v\n", http.StatusOK, resp.StatusCode)
	}
}

func TestCallDecryptNoPasswordHash(t *testing.T) {
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
		req, _ := http.NewRequest(http.MethodPost, "decrypt", &reqBody)
		req.Header.Set("Content-Type", w.FormDataContentType())
		rec := httptest.NewRecorder()
		Decrypt(rec, req)
		resp = rec.Result()
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected %v; got %v\n", http.StatusOK, resp.StatusCode)
	}
}
