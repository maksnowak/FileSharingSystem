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

// UpdateAccount 	godoc
//
//	@Summary		Update an account
//	@Description	Update an account with given ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string			true	"ID of the user to update"
//	@Param			request	body		models.Update	true "Data to be updated (no need for all the fields)"
//	@Success		200		{object}	models.HTTP200	"Account updated successfully"
//	@Failure		400		{object}	models.HTTP400	"Invalid ID format or request body"
//	@Failure		404		{object}	models.HTTP404	"No user with given ID was found"
//	@Failure		500		{object}	models.HTTP500	"Server could not update the account"
//	@Router			/accounts/{user_id} [put]
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

// DeleteAccount 	godoc
//
//	@Summary		Delete an account
//	@Description	Delete an account with given ID
//	@Tags			accounts
//	@Produce		json
//	@Param			user_id	path		string			true	"ID of the user to delete"
//	@Success		200		{object}	models.HTTP200	"Account deleted successfully"
//	@Failure		400		{object}	models.HTTP400	"Invalid ID format"
//	@Failure		404		{object}	models.HTTP404	"No user with given ID was found"
//	@Failure		500		{object}	models.HTTP500	"Server could not delete the account"
//	@Router			/accounts/{user_id} [delete]
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
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "Account deleted successfully"})
}
