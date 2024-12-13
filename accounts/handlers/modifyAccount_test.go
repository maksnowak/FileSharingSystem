package handlers

import (
	"accounts/db"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUpdateAccountBadID(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	router := chi.NewRouter()
	router.Put("/accounts/{user_id}", UpdateAccount)
	rec := Call(nil, http.MethodPut, "/accounts/wrong_id", nil, router)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid account ID format\n", rec.Body.String())
}

func TestUpdateAccountBadBody(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	router := chi.NewRouter()
	router.Put("/accounts/{user_id}", UpdateAccount)
	rec := Call(nil, http.MethodPut, "/accounts/123456789012345678901234", nil, router)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid JSON body\n", rec.Body.String())
}

func TestUpdateAccountNonexistent(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	router := chi.NewRouter()
	router.Put("/accounts/{user_id}", UpdateAccount)
	rec := Call(up, http.MethodPut, "/accounts/67599a97ca18148343366a54", nil, router)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "Account not found\n", rec.Body.String())
}

func TestUpdateAccountStandard(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	_, _ = coll.InsertOne(context.Background(), u2)
	_, _ = coll.InsertOne(context.Background(), u3)
	router := chi.NewRouter()
	router.Put("/accounts/{user_id}", UpdateAccount)
	rec := Call(up, http.MethodPut, "/accounts/67599a97ca18148343366a54", nil, router)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "Account updated successfully"}`, rec.Body.String())
	res := u3
	res.Email = up.Email
	res.PasswordHash = up.PasswordHash
	res.PasswordSalt = up.PasswordSalt
	res.OwnedFiles = up.OwnedFiles
	res.SharedFiles = up.SharedFiles
	router.Get("/accounts/{user_id}", GetAccountByID)
	rec = Call(nil, http.MethodGet, "/accounts/67599a97ca18148343366a54", nil, router)
	assert.Equal(t, http.StatusOK, rec.Code)
	js, _ := json.Marshal(res)
	assert.JSONEq(t, string(js), rec.Body.String())
}

func TestDeleteAccountBadID(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	router := chi.NewRouter()
	router.Delete("/accounts/{user_id}", DeleteAccount)
	rec := Call(nil, http.MethodDelete, "/accounts/wrong_id", nil, router)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid account ID format\n", rec.Body.String())
}

func TestDeleteAccountNonexistent(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	router := chi.NewRouter()
	router.Delete("/accounts/{user_id}", DeleteAccount)
	rec := Call(nil, http.MethodDelete, "/accounts/67599a97ca18148343366a54", nil, router)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "Account not found\n", rec.Body.String())
}

func TestDeleteAccountStandard(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	_, _ = coll.InsertOne(context.Background(), u2)
	_, _ = coll.InsertOne(context.Background(), u3)
	router := chi.NewRouter()
	router.Delete("/accounts/{user_id}", DeleteAccount)
	rec := Call(up, http.MethodDelete, "/accounts/67599a97ca18148343366a54", nil, router)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "Account deleted successfully"}`, rec.Body.String())
	router.Get("/accounts/{user_id}", GetAccountByID)
	rec = Call(nil, http.MethodGet, "/accounts/67599a97ca18148343366a54", nil, router)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "Account not found\n", rec.Body.String())
}
