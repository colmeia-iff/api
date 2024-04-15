package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.mod/entity"
	"go.mod/rest"
)

func HiveRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", hiveHandler)
	r.Get("/", hiveReturnHandler)
	return r
}

func hiveHandler(w http.ResponseWriter, r *http.Request) {
	var data entity.Hive
	if err := rest.ParseBody(w, r, &data); err != nil {
		return
	}
}

func hiveReturnHandler(w http.ResponseWriter, r *http.Request) {

}