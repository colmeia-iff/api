package core

import (
	"context"
	"log"

	"go.mod/db"
	"go.mod/entity"
)

type DataManager struct {
}

func DataManagerNew() *DataManager {
	return &DataManager{}
}

func (m *DataManager) DataCreateInfo(ctx context.Context, body entity.Data) error {
	if err := db.Moisture(ctx, body.Moisture); err != nil {
		log.Println("db.Moisture m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.OutsideTemperature(ctx, body.OutsideTemperature); err != nil {
		log.Println("db.OutsideTemperature m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.Temperature(ctx, body.Temperature); err != nil {
		log.Println("db.Temperature m.DataCreateInfo err: ", err)
		return err
	}
	if err := db.WeightHive(ctx, body.Weight); err != nil {
		log.Println("db.WeightHive m.DataCreateInfo err: ", err)
		return err
	}
	return nil
}

func (m *DataManager) ReturnInfos(ctx context.Context, id string) (*entity.DataInfo, error) {
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
	data := entity.DataInfo{
		Name:               name.Name,
		Slug:               name.Slug,
		Moisture:           mt,
		Weight:             w,
		Temperature:        t,
		OutsideTemperature: to,
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
