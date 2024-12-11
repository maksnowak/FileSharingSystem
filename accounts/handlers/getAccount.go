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

func GetAllAccounts(w http.ResponseWriter, _ *http.Request) {
	var users []models.User
	ctx := context.TODO()
	collection := db.GetCollection("users")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		http.Error(w, "Error while reading database", http.StatusInternalServerError)
		return
	}

	err = cursor.All(ctx, &users)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error while mapping data", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "user_id")

	objectID, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid account ID format", http.StatusBadRequest)
		return
	}

	var user models.User
	collection := db.GetCollection("users")
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}
