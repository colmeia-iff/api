package core

import (
	"context"
	"log"

	"go.mod/db"
	"go.mod/entity"
)

type HiveManager struct {
}

func HiveManagerNew() *HiveManager {
	return &HiveManager{}
}

func (m *HiveManager) Create(ctx context.Context, data entity.InfoData) error {
	if err := db.HiveCreate(ctx, data); err != nil {
		log.Println("m *HiveManager db.HiveCreate err: ", err)
		return err
	}
	return nil
}
