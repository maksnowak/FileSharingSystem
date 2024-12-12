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

// UpdateAccount 	godoc
//
//	@Summary		Get the user's password salt
//	@Description	Get the password salt of the user with given username
//	@Tags			login
//	@Produce		json
//	@Param			username	path		string			true	"Username of the account to retrieve the password salt for"
//	@Success		200			{object}	models.Salt		"Password salt retrieved successfully"
//	@Failure		404			{object}	models.HTTP404	"No account with given username was found"
//	@Router			/login/{username} [get]
func GetPasswordSalt(w http.ResponseWriter, r *http.Request) {
	uname := chi.URLParam(r, "username")

	collection := db.GetCollection("users")
	var user models.User

	err := collection.FindOne(context.Background(), bson.M{"username": uname}).Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not find user with this name", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"username": user.Username, "passwordSalt": user.PasswordSalt})
}

// UpdateAccount 	godoc
//
//	@Summary		Verify users password
//	@Description	Verify users password and return the User if it is correct (login successful)
//	@Tags			login
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Credentials	true	"Users login credentials"
//	@Success		200		{object}	models.User			"User credentials valid (login successful)"
//	@Failure		400		{object}	models.HTTP400		"Invalid request body or password"
//	@Failure		404		{object}	models.HTTP404		"No account with given username was found"
//	@Router			/login/ [get]
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
		http.Error(w, "Could not find user with this name", http.StatusNotFound)
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
