package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.mod/core"
	"go.mod/entity"
	"go.mod/rest"
)

func DataRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", dataCreateHandler)
	r.Get("/", dataGetHandler)
	return r
}

func dataCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body entity.Data
	if err := rest.ParseBody(w, r, &body); err != nil {
		rest.SendError(w, err)
		return
	}
	err := core.ApiaryManagerNew().DataCreateInfo(ctx, body)
	if err != nil {
		rest.SendError(w, err)
		return
	}

}

func dataGetHandler(w http.ResponseWriter, r *http.Request) {

}
