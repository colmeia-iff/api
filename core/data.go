package core

import (
	"context"

	"go.mod/db"
	"go.mod/entity"
	"go.mod/rest"
)

type DataManager struct {
}

func DataManagerNew() *DataManager {
	return &DataManager{}
}

func (m *DataManager) DataCreateInfo(ctx context.Context, body entity.Data) error {
	if err := db.CrateInfoDB(ctx, body); err != nil {
		return rest.LogError(err, "m.DataCreateInfo")
	}
	return nil
}
