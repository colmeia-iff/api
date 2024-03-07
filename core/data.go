package core

import (
	"context"

	"go.mod/entity"
)

type DataManager struct {
}

func DataManagerNew() *DataManager {
	return &DataManager{}
}

func (m *DataManager) DataCreateInfo(ctx context.Context, body entity.Data) error {
	return nil
}
