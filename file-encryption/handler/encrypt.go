package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"net/http"
	_ "os"
)

// Encrypt godoc
//
//	@Description	Encrypt a file
//	@Tags			files
//	@Accept			mpfd
//	@Produce		octet-stream
//	@Param			file			formData	file	true	"File which contents to encrypt"
//	@Param			password-hash	formData	string	true	"Hash of the password that protect the file"
//	@Success		200				{object}	int		"Encrypted file"
//	@Failure		400				{object}	int		"Invalid request body"
//	@Failure		500				{object}	int		"Server could not read the file"
//	@Failure		500				{object}	int		"Server could not encrypt the file"
//	@Router			/encrypt [post]
func Encrypt(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password-hash")
	if len(password) == 0 {
		http.Error(w, "Cannot protect a file with an empty password", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	plainText, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cipherText, err := encrypt(plainText, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(cipherText)
}

func encrypt(plainText []byte, password string) (cipherText []byte, err error) {
	key := sha256.Sum256([]byte(password))
	block, _ := aes.NewCipher(key[:])
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	cipherText = gcm.Seal(nonce, nonce, plainText, nil)
	return cipherText, nil
}
