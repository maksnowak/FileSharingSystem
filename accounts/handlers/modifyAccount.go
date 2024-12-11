package handlers

import (
	"accounts/db"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
	"net/http"
)

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "user_id")

	objectID, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid account ID format", http.StatusBadRequest)
		return
	}

	var updateFields map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateFields); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	// delete fields that can't be updated
	delete(updateFields, "_id")
	delete(updateFields, "username")
	delete(updateFields, "createdAt")
	delete(updateFields, "role")

	update := bson.M{"$set": updateFields}
	collection := db.GetCollection("users")
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)

	if err != nil {
		log.Println(err)
		http.Error(w, "Could not update account", http.StatusInternalServerError)
		return
	}

	if res.MatchedCount == 0 {
		log.Println("Could not find account")
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "Account updated successfully"})
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "user_id")

	objectID, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid account ID format", http.StatusBadRequest)
		return
	}

	collection := db.GetCollection("users")
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectID})

	if err != nil {
		log.Println(err)
		http.Error(w, "Could not delete account", http.StatusInternalServerError)
		return
	}

	if res.DeletedCount == 0 {
		log.Println("Could not find account")
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account deleted successfully"})
}
