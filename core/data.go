package core

import (
	"context"
	"fmt"
	"log"

	"go.mod/db"
	"go.mod/entity"
)

type DataManager struct {
}

func DataManagerNew() *DataManager {
	return &DataManager{}
}

func (m *DataManager) DataCreateInfo(ctx context.Context, body entity.Data, id string) error {

	if err := db.Moisture(ctx, body.Moisture, id); err != nil {
		log.Println("db.Moisture m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.OutsideTemperature(ctx, body.OutsideTemperature, id); err != nil {
		log.Println("db.OutsideTemperature m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.Temperature(ctx, body.Temperature, id); err != nil {
		log.Println("db.Temperature m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.WeightHive(ctx, body.Weight, id); err != nil {
		log.Println("db.WeightHive m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.Vento(ctx, body.Vento, id); err != nil {
		log.Println("db.Moisture m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.Melg(ctx, body.Melg, id); err != nil {
		log.Println("db.Moisture m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.Volt(ctx, body.Voltage, id); err != nil {
		log.Println("db.Moisture m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.Resist(ctx, body.Resistance, id); err != nil {
		log.Println("db.Moisture m.DataCreateInfo err: ", err)
		return err
	}
	return nil
}

func (m *DataManager) ReturnInfos(ctx context.Context, id string) (*entity.DataInfo, error) {
	fmt.Println("teste", id)
	log.Println("teste", id)
	name, err := db.ReturnDataInfo(ctx, id)
	if err != nil {
		log.Println("db.ReturnDataInfo err: ", err)
		return nil, err
	}

	w, err := db.GetWeightHive(ctx, id)
	if err != nil {
		log.Println("db.GetWeightHive err: ", err)
		return nil, err
	}
	mt, err := db.GetMoisture(ctx, id)
	if err != nil {
		log.Println("db.GetMoisture err: ", err)
		return nil, err
	}
	t, err := db.GetTemperature(ctx, id)
	if err != nil {
		log.Println("db.GetTemperature err: ", err)
		return nil, err
	}
	to, err := db.GetOutsideTemperature(ctx, id)
	if err != nil {
		log.Println("db.GetOutsideTemperature err: ", err)
		return nil, err
	}
	vt, err := db.GetVento(ctx, id)
	if err != nil {
		log.Println("db.GetMoisture err: ", err)
		return nil, err
	}
	vg, err := db.GetVolt(ctx, id)
	if err != nil {
		log.Println("db.GetMoisture err: ", err)
		return nil, err
	}
	pm, err := db.GetMelg(ctx, id)
	if err != nil {
		log.Println("db.GetMoisture err: ", err)
		return nil, err
	}
	r, err := db.GetResist(ctx, id)
	if err != nil {
		log.Println("db.GetMoisture err: ", err)
		return nil, err
	}
	data := entity.DataInfo{
		Name:               name.Name,
		Slug:               name.Slug,
		Moisture:           mt,
		Weight:             w,
		Temperature:        t,
		OutsideTemperature: to,
		Melg:               pm,
		Vento:              vt,
		Voltage:            vg,
		Resist:             r,
	}

	return &data, nil

}

func (m *DataManager) HiveBySlug(ctx context.Context, slug string) ([]entity.HiveInitial, error) {
	data, err := db.ReturnDataInfoBySlug(ctx, slug)
	if err != nil {
		log.Println("db.ReturnDataInfoBySlug err: ", err)
		return nil, err
	}
	return data, nil
}
