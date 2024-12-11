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
	json.NewEncoder(w).Encode(map[string]string{"username": user.Username, "passwordSalt": user.PasswordSalt})
}
