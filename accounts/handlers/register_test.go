package handlers

import (
	"accounts/db"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var fn Handler = Register

func TestRegisterStandard(t *testing.T) {
	defer CleanDatabase()
	rec := Call(reg, http.MethodPost, "/accounts", fn, nil)
	// Validate the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message":"Account created successfully"}`, rec.Body.String())
}

func TestRegisterWrongBody(t *testing.T) {
	defer CleanDatabase()
	rec := Call(nil, http.MethodPost, "/accounts", fn, nil)
	// Validate the response
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid body\n", rec.Body.String())
}

func TestRegisterMissingFields(t *testing.T) {
	defer CleanDatabase()
	// no username
	msg := reg
	msg.Username = ""
	// Validate the response
	rec := Call(msg, http.MethodPost, "/accounts", fn, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Missing required fields\n", rec.Body.String())
	msg = reg
	msg.Email = ""
	rec = Call(msg, http.MethodPost, "/accounts", fn, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Missing required fields\n", rec.Body.String())
	msg = reg
	msg.PasswordHash = ""
	rec = Call(msg, http.MethodPost, "/accounts", fn, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Missing required fields\n", rec.Body.String())
	msg = reg
	msg.PasswordSalt = ""
	rec = Call(msg, http.MethodPost, "/accounts", fn, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Missing required fields\n", rec.Body.String())
	msg = reg
	msg.Role = ""
	rec = Call(msg, http.MethodPost, "/accounts", fn, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Missing required fields\n", rec.Body.String())
}

func TestRegisterInvalidRole(t *testing.T) {
	defer CleanDatabase()
	msg := reg
	msg.Role = "piesek"
	rec := Call(msg, http.MethodPost, "/accounts", fn, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Wrong role\n", rec.Body.String())
}

func TestRegisterUsernameTaken(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	msg := reg
	msg.Username = u1.Username
	rec := Call(msg, http.MethodPost, "/accounts", fn, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Username taken\n", rec.Body.String())
}

func TestRegisterEmailTaken(t *testing.T) {
	defer CleanDatabase()
	coll := db.GetCollection("users")
	_, _ = coll.InsertOne(context.Background(), u1)
	msg := reg
	msg.Email = u1.Email
	rec := Call(msg, http.MethodPost, "/accounts", fn, nil)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Username taken\n", rec.Body.String())
}
