package handlers

import (
	"accounts/db"
	"accounts/models"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
	"net/http"
)

func GetPasswordSalt(w http.ResponseWriter, r *http.Request) {
	uname := chi.URLParam(r, "username")

	collection := db.GetCollection("users")
	var user models.User

	err := collection.FindOne(context.Background(), bson.M{"username": uname}).Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not find user with this name", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"username": user.Username, "passwordSalt": user.PasswordSalt})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&credentials)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	collection := db.GetCollection("users")
	var user models.User

	err = collection.FindOne(context.Background(), bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not find user with this name", http.StatusBadRequest)
		return
	}

	if credentials.PasswordHash != user.PasswordHash {
		log.Println("Password incorrect")
		http.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}
