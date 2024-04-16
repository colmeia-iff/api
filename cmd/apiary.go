package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.mod/core"
	"go.mod/entity"
	"go.mod/rest"
)

func ApiaryRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createApiaryHandler)

	return r
}

func createApiaryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body entity.Apiary
	if err := rest.ParseBody(w, r, &body); err != nil {
		return
	}
	manager := core.ApiaryManagerNew()
	err := manager.Create(ctx, body)
	if err != nil {
		rest.SendError(w, err)
		return
	}

}
