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

var login Handler = Login

func TestGetPasswordSaltNonexistent(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	router := chi.NewRouter()
	router.Get("/login/{username}", GetPasswordSalt)
	rec := Call(nil, http.MethodGet, "/login/nonexistent_user", nil, router)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "Could not find user with this name\n", rec.Body.String())
}

func TestGetPasswordSaltStandard(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	_, _ = coll.InsertOne(context.Background(), u2)
	_, _ = coll.InsertOne(context.Background(), u3)
	router := chi.NewRouter()
	router.Get("/login/{username}", GetPasswordSalt)
	rec := Call(nil, http.MethodGet, "/login/Test2", nil, router)
	assert.Equal(t, http.StatusOK, rec.Code)
	res := models.Salt{
		Username:     u2.Username,
		PasswordSalt: u2.PasswordSalt,
	}
	exp, _ := json.Marshal(res)
	assert.JSONEq(t, string(exp), rec.Body.String())
}

func TestLoginWrongBody(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	rec := Call(nil, http.MethodGet, "/login/", login, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid body\n", rec.Body.String())
}

func TestLoginNonexistent(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	cred := models.Credentials{
		Username:     "Nonexistent",
		PasswordHash: "Wrong_Password_Hash",
	}
	rec := Call(cred, http.MethodGet, "/login/", login, nil)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "Could not find user with this name\n", rec.Body.String())
}

func TestLoginWrongPassword(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	cred := models.Credentials{
		Username:     u1.Username,
		PasswordHash: "Wrong_Password_Hash",
	}
	rec := Call(cred, http.MethodGet, "/login/", login, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Incorrect password\n", rec.Body.String())
}

func TestLoginStandard(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	_, _ = coll.InsertOne(context.Background(), u2)
	_, _ = coll.InsertOne(context.Background(), u3)
	cred := models.Credentials{
		Username:     u1.Username,
		PasswordHash: u1.PasswordHash,
	}
	rec := Call(cred, http.MethodGet, "/login/", login, nil)
	assert.Equal(t, http.StatusOK, rec.Code)
	exp, _ := json.Marshal(u1)
	assert.JSONEq(t, string(exp), rec.Body.String())
}
