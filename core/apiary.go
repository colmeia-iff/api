package core

import (
	"context"
	"log"

	"go.mod/db"
	"go.mod/entity"
)

type ApiaryManager struct {
}

func ApiaryManagerNew() *ApiaryManager {
	return &ApiaryManager{}
}

func (m *ApiaryManager) Create(ctx context.Context, data entity.Apiary) error {
	err := db.ApiaryCreate(ctx, data)
	if err != nil {
		log.Println("ApiaryManager.Create db.ApiaryCreate error err: ", err)
		return err
	}
	return nil
}

func (m *ApiaryManager) ReturnApiarys(ctx context.Context) ([]entity.Apiary, error) {
	data, err := db.ReturnApiary(ctx)
	if err != nil {
		log.Println("ApiaryManager.Create db.ReturnApiary error err: ", err)
		return nil, err
	}

	return data, nil
}
