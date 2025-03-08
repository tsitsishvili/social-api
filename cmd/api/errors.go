package main

import (
	"log"
	"net/http"
)

func (app *application) statusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusInternalServerError, "there was an unexpected error")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusNotFound, "not found")
}
