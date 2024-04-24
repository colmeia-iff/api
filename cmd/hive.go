package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mod/core"
	"go.mod/entity"
	"go.mod/rest"
)

func HiveRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", hiveHandler)
	r.Post("/create/{idExterno}", updatedDataHandler)
	r.Get("/data/{id}", hiveReturnHandler)
	r.Get("/{slug}", returnBySlugHandler)
	return r
}

func hiveHandler(w http.ResponseWriter, r *http.Request) {
	var data entity.InfoData
	ctx := r.Context()
	if err := rest.ParseBody(w, r, &data); err != nil {
		return
	}
	manager := core.HiveManagerNew()
	err := manager.Create(ctx, data)
	if err != nil {
		rest.SendError(w, err)
		return
	}

}

func updatedDataHandler(w http.ResponseWriter, r *http.Request) {

	type Data struct {
		Moisture           string `json:"moisture"`
		Temperature        string `json:"temperature"`
		OutsideTemperature string `json:"outsidetemperature"`
		Weight             string `json:"weight"`
	}

	var data Data
	ctx := r.Context()
	if err := rest.ParseBody(w, r, &data); err != nil {
		rest.SendError(w, err)
		return
	}
	manager := core.DataManagerNew()
	send := entity.Data{
		Moisture: entity.Moisture{
			Data: entity.Values{
				Temp: data.Moisture,
			},
		},
		OutsideTemperature: entity.OutsideTemperature{
			Data: entity.Values{Temp: data.OutsideTemperature}},
		Temperature: entity.Temperature{
			Data: entity.Values{
				Temp: data.Temperature,
			},
		},
		Weight: entity.Weight{
			Value: data.Weight,
		},
	}
	err := manager.DataCreateInfo(ctx, send, chi.URLParam(r, "idExterno"))
	if err != nil {
		rest.SendError(w, err)
		return
	}

}

func hiveReturnHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	manager := core.DataManagerNew()
	data, err := manager.ReturnInfos(ctx, id)
	if err != nil {
		rest.SendError(w, err)
		return
	}

	rest.Send(w, data)
}

func returnBySlugHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := chi.URLParam(r, "slug")
	manager, err := core.DataManagerNew().HiveBySlug(ctx, slug)
	if err != nil {
		rest.SendError(w, err)
		return
	}

	rest.Send(w, manager)
}
