package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func ApiaryCreate(ctx context.Context, data entity.Apiary) error {
	_, err := infradb.Get().ExecContext(ctx, `
	INSERT INTO apiary (IdExterno, Name, Slug)
	VALUES
		$1,
		$2,
		$3;
	`, data.IdExterno, data.Name)
	if err != nil {
		return err
	}

	return nil
}

func ReturnApiary(ctx context.Context, data entity.Apiary) ([]entity.Apiary, error) {
	var body entity.Apiary
	rows, err := infradb.Get().QueryContext(ctx, `select IdExterno ,Name, Slug from apiary;`)
	if err != nil {
		return nil, err
	}

	var array []entity.Apiary
	for rows.Next() {
		rows.Scan(&body.IdExterno, &body.Name, &body.Slug)
		array = append(array, body)
	}
	return array, nil
}
