package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func DataRoutr() http.Handler {
	r := chi.NewRouter()
	r.Post("/", dataCreateHandler)
	r.Get("/", dataGetHandler)
	return r
}

func dataCreateHandler(w http.ResponseWriter, r *http.Request) {

}

func dataGetHandler(w http.ResponseWriter, r *http.Request) {

}
