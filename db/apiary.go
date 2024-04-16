package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func ApiaryCreate(ctx context.Context, data entity.Apiary) error {
	_, err := infradb.Get().ExecContext(ctx, ``, data.IdExterno, data.Name,)
	if err != nil {
		return err
	}

	return nil
}
