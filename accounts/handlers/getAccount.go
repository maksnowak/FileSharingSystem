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

// GetAllAccounts 	godoc
//
//	@Summary		Retrieve all account data
//	@Description	Retrieve information about all existing accounts
//	@Tags			accounts
//	@Produce		json
//	@Success		200	{array}		models.User		"Every existing account"
//	@Failure		500	{object}	models.HTTP500	"Server could not retrieve or process the data"
//	@Router			/accounts/ [get]
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

// GetAccountByID 	godoc
//
//	@Summary		Retrieve an account
//	@Description	Retrieve information about an account with given ID
//	@Tags			accounts
//	@Produce		json
//	@Param			user_id	path		string			true	"ID of the user to retrieve"
//	@Success		200		{object}	models.User		"Account retrieved successfully"
//	@Failure		400		{object}	models.HTTP400	"Invalid ID format"
//	@Failure		404		{object}	models.HTTP404	"No user with given ID was found"
//	@Router			/accounts/{user_id} [get]
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
