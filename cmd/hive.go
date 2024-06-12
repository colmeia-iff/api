package main

import (
	"net/http"
	"strconv"

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
		Moisture           int `json:"uNinho"`
		Temperature        int `json:"tNinho"`
		OutsideTemperature int `json:"tExt"`
		Weight             int `json:"pNinho"`
		Volt               int `json:"vBat"`
		WeightMel          int `json:"pMelg"`
		TempVent           int `json:"tVent"`
		Tresist            int `json:"tResist"`
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
				Temp: strconv.Itoa(data.Moisture), // Conversão de int para string
			},
		},
		OutsideTemperature: entity.OutsideTemperature{
			Data: entity.Values{
				Temp: strconv.Itoa(data.OutsideTemperature), // Conversão de int para string
			},
		},
		Temperature: entity.Temperature{
			Data: entity.Values{
				Temp: strconv.Itoa(data.Temperature), // Conversão de int para string
			},
		},
		Weight: entity.Weight{
			Value: strconv.Itoa(data.Weight),
		},
		Melg: entity.Melg{
			Data: entity.ValuesNew{
				Values: strconv.Itoa(data.WeightMel), // Conversão de int para string
			},
		},
		Voltage:    entity.Voltage{Data: entity.ValuesNew{Values: strconv.Itoa(data.Volt)}},   // Conversão de int para string
		Resistance: entity.Resist{Data: entity.ValuesNew{Values: strconv.Itoa(data.Tresist)}}, // Conversão de int para string
		Vento:      entity.Vento{Data: entity.ValuesNew{Values: strconv.Itoa(data.TempVent)}}, // Conversão de int para string
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
