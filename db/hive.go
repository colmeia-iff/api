package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func HiveCreate(ctx context.Context, data entity.Hive) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Hive (IdExterno, Name) VALUES ($1, $2);`)
	if err != nil {
		return err
	}
	if err := Moisture(ctx, data.Moisture); err != nil {
		return err
	}
	if err := OutsideTemperature(ctx, data.OutsideTemperature); err != nil {
		return err
	}
	if err := Temperature(ctx, data.Temperature); err != nil {
		return err
	}
	if err := WeightHive(ctx, data.Weight); err != nil {
		return err
	}
	return nil

}

func Moisture(ctx context.Context, data entity.Moisture) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Moisture (IdExterno, Temp, Time) VALUES ($1, $2, $3);`)
	if err != nil {
		return err
	}
	return nil
}

func OutsideTemperature(ctx context.Context, data entity.OutsideTemperature) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO OutsideTemperature (IdExterno, Temp, Time) VALUES ($1, $2, $3);`)
	if err != nil {
		return err
	}
	return nil
}

func Temperature(ctx context.Context, data entity.Temperature) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Temperature (IdExterno, Temp, Time) VALUES ($1, $2, $3);`)
	if err != nil {
		return err
	}
	return nil
}

func WeightHive(ctx context.Context, data entity.Weight) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO WeightHive (IdExterno, Value, Time) VALUES ($1, $2,$3);`)
	if err != nil {
		return err
	}
	return nil
}
