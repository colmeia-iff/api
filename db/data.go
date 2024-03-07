package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func CrateInfoDB(ctx context.Context, body entity.Data) error {
	_, err := infradb.DB.ExecContext(ctx, `insert `)
	if err != nil {
		return err
	}
	return nil
}
