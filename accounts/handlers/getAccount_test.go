package handlers

import (
	"accounts/db"
	"accounts/models"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var all Handler = GetAllAccounts
var one Handler = GetAccountByID

func TestGetAllAccountsEmpty(t *testing.T) {
	rec := Call(nil, http.MethodGet, "/accounts", all, nil)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "null", rec.Body.String())
}

func TestGetAllAccountsStandard(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	_, _ = coll.InsertOne(context.Background(), u2)
	_, _ = coll.InsertOne(context.Background(), u3)
	rec := Call(nil, http.MethodGet, "/accounts", all, nil)
	assert.Equal(t, http.StatusOK, rec.Code)
	res, _ := json.Marshal([]models.User{u1, u2, u3})
	assert.JSONEq(t, string(res), rec.Body.String())
}

func TestGetAccountByIDBadID(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	router := chi.NewRouter()
	router.Get("/accounts/{user_id}", GetAccountByID)
	rec := Call(reg, http.MethodGet, "/accounts/wrong_id", nil, router)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid account ID format\n", rec.Body.String())
}

func TestGetAccountByIDNonexistent(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	router := chi.NewRouter()
	router.Get("/accounts/{user_id}", GetAccountByID)
	rec := Call(nil, http.MethodGet, "/accounts/67599a97ca18148343366a54", nil, router)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "Account not found\n", rec.Body.String())
}

func TestGetAccountByIDStandard(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	_, _ = coll.InsertOne(context.Background(), u2)
	_, _ = coll.InsertOne(context.Background(), u3)
	router := chi.NewRouter()
	router.Get("/accounts/{user_id}", GetAccountByID)
	rec := Call(nil, http.MethodGet, "/accounts/102938475610293847561029", nil, router)
	assert.Equal(t, http.StatusOK, rec.Code)
	res, _ := json.Marshal(u2)
	assert.JSONEq(t, string(res), rec.Body.String())
}
