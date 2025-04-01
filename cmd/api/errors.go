package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw(
		"internal server error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusInternalServerError, "there was an unexpected error")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw(
		"bad request",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw(
		"conflict request",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw(
		"not found error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusNotFound, "not found")
}
