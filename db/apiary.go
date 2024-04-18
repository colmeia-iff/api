package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func ApiaryCreate(ctx context.Context, data entity.Apiary) error {
	_, err := infradb.Get().ExecContext(ctx, `
		INSERT INTO apiary (Name, Slug, Local)
		VALUES ($1, $2, $3)
	`, data.Name, data.Slug, data.Local)
	if err != nil {
		return err
	}

	return nil
}

func ReturnApiary(ctx context.Context) ([]entity.Apiary, error) {
	var body entity.Apiary
	rows, err := infradb.Get().QueryContext(ctx, `select  name, slug , local from apiary;`)
	if err != nil {
		return nil, err
	}

	var array []entity.Apiary
	for rows.Next() {
		rows.Scan(&body.Name, &body.Slug, &body.Local)
		array = append(array, body)
	}
	return array, nil
}
