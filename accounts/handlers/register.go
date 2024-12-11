package handlers

import (
	"accounts/db"
	"accounts/models"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"net/http"
	"time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// reading request
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	// checking if any required fields are missing
	if user.Username == "" || user.Email == "" || user.PasswordHash == "" || user.PasswordSalt == "" || user.Role == "" {
		log.Println("Incomplete data")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// check if role is valid
	if user.Role != "user" && user.Role != "admin" {
		log.Println(err)
		http.Error(w, "Wrong role", http.StatusBadRequest)
		return
	}

	collection := db.GetCollection("users")

	// check if username or email is already taken
	var existingUser models.User
	err = collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		log.Println(err)
		http.Error(w, "Username taken", http.StatusBadRequest)
		return
	}

	err = collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		log.Println(err)
		http.Error(w, "Email taken", http.StatusBadRequest)
		return
	}

	// everything ok, inserting user
	user.CreatedAt = time.Now()
	user.OwnedFiles = []string{}
	user.SharedFiles = []string{}
	log.Println(user)
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error saving the user account in database", http.StatusInternalServerError)
		return
	}
	// success
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "Account created successfully"})
}
