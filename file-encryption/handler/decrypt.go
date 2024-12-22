package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"io"
	"net/http"
)

// Decrypt godoc
//
//	@Description	Decrypt a file
//	@Tags			files
//	@Accept			mpfd
//	@Produce		octet-stream
//	@Param			file			formData	file	true	"Encrypted file which contents to decrypt"
//	@Param			password-hash	formData	string	true	"Hash of the password that protect the file"
//	@Success		200				{object}	int		"Decrypted plain file"
//	@Failure		400				{object}	int		"Invalid request body"
//	@Failure		403				{object}	int		"Server could not decrypt the file on account on bad password"
//	@Failure		500				{object}	int		"Server could not read the file"
//	@Router			/decrypt [post]
func Decrypt(w http.ResponseWriter, r *http.Request) {
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

	cipherText, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := sha256.Sum256([]byte(password))
	block, _ := aes.NewCipher(key[:])
	gcm, _ := cipher.NewGCM(block)

	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		http.Error(w, "Bad password", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(plainText)
}

func decrypt(cipherText []byte, password string) (plainText []byte, err error) {
	key := sha256.Sum256([]byte(password))
	block, _ := aes.NewCipher(key[:])
	gcm, _ := cipher.NewGCM(block)

	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err = gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
