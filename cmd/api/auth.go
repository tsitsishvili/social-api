package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"github.com/tsitsishvili/social/internal/store"
	"net/http"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

//registerUserHandler godoc

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// has password
	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
	}

	ctx := r.Context()

	plainToken := uuid.New().String()

	hash := sha256.Sum256([]byte(plainToken))
	hashedToken := hex.EncodeToString(hash[:])

	//store user
	if err := app.store.Users.CreateAndInvite(ctx, user, hashedToken, app.config.mail.exp); err != nil {
		switch {
		case errors.Is(err, store.ErrDuplicateEmail):
			app.badRequestResponse(w, r, err)
		case errors.Is(err, store.ErrDuplicateUsername):
			app.badRequestResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}

		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}
