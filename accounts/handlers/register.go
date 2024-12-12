package handlers

import (
	"accounts/db"
	"accounts/models"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"time"
)

// Register 	godoc
//
//	@Summary		Create an account
//	@Description	Create a User record in the database
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Register	true	"Necessary account details"
//	@Success		200		{object}	models.HTTP200	"Account created successfully"
//	@Failure		400		{object}	models.HTTP400	"Invalid request body"
//	@Failure		500		{object}	models.HTTP500	"Server could not save the account"
//	@Router			/accounts/ [post]
func Register(w http.ResponseWriter, r *http.Request) {
	// reading request
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	// checking if any required fields are missing
	if user.Username == "" || user.Email == "" || user.PasswordHash == "" || user.PasswordSalt == "" || user.Role == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// check if role is valid
	if user.Role != "user" && user.Role != "admin" {
		http.Error(w, "Wrong role", http.StatusBadRequest)
		return
	}

	collection := db.GetCollection("users")

	// check if username or email is already taken
	var existingUser models.User
	err = collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		http.Error(w, "Username taken", http.StatusBadRequest)
		return
	}

	err = collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		http.Error(w, "Email taken", http.StatusBadRequest)
		return
	}

	// everything ok, inserting user
	user.CreatedAt = time.Now()
	user.OwnedFiles = []string{}
	user.SharedFiles = []string{}
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, "Error saving the user account in database", http.StatusInternalServerError)
		return
	}
	// success
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "Account created successfully"})
}
