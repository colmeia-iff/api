package db

import (
	"context"

	infradb "go.mod/connect"
	"go.mod/entity"
)

func HiveCreate(ctx context.Context, data entity.InfoData) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Hive (IdExterno, Name, Slug, Description) VALUES ($1, $2, $3, $4);`, data.IdExterno, data.Name, data.Slug, data.Description)
	if err != nil {
		return err
	}
	// if err := Moisture(ctx, data.Moisture); err != nil {
	// 	return err
	// }
	// if err := OutsideTemperature(ctx, data.OutsideTemperature); err != nil {
	// 	return err
	// }
	// if err := Temperature(ctx, data.Temperature); err != nil {
	// 	return err
	// }
	// if err := WeightHive(ctx, data.Weight); err != nil {
	// 	return err
	// }
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

func ReturnDataInfo(ctx context.Context, id string) (*entity.InfoData, error) {
	var data entity.InfoData
	err := infradb.Get().QueryRowContext(ctx, `select Name, Slug from hive where idExterno = $1 `, id).Scan(&data.Name, &data.Slug)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetTemperature(ctx context.Context, id string) ([]entity.Temperature, error) {
	rows, err := infradb.Get().QueryContext(ctx, `
		SELECT  Value, Time 
		FROM Temperature 
		WHERE idExterno = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weights []entity.Temperature
	for rows.Next() {
		var temp entity.Temperature
		if err := rows.Scan(&temp.Data.Temp, &temp.Data.Time); err != nil {
			return nil, err
		}
		weights = append(weights, temp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return weights, nil
}

func GetOutsideTemperature(ctx context.Context, id string) ([]entity.OutsideTemperature, error) {
	rows, err := infradb.Get().QueryContext(ctx, `
		SELECT  Value, Time 
		FROM OutsideTemperature 
		WHERE idExterno = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []entity.OutsideTemperature
	for rows.Next() {
		var dataifo entity.OutsideTemperature
		if err := rows.Scan(&dataifo.Data.Temp, &dataifo.Data.Time); err != nil {
			return nil, err
		}
		data = append(data, dataifo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
func GetWeightHive(ctx context.Context, id string) ([]entity.Weight, error) {
	rows, err := infradb.Get().QueryContext(ctx, `
		SELECT  Value, Time 
		FROM WeightHive 
		WHERE idExterno = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weights []entity.Weight
	for rows.Next() {
		var weight entity.Weight
		if err := rows.Scan(&weight.Value, &weight.Time); err != nil {
			return nil, err
		}
		weights = append(weights, weight)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return weights, nil
}

func GetMoisture(ctx context.Context, id string) ([]entity.Moisture, error) {
	rows, err := infradb.Get().QueryContext(ctx, `
		SELECT  Value, Time 
		FROM Moisture 
		WHERE idExterno = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Moistures []entity.Moisture
	for rows.Next() {
		var Moisture entity.Moisture
		if err := rows.Scan(&Moisture.Data.Temp, &Moisture.Data.Time); err != nil {
			return nil, err
		}
		Moistures = append(Moistures, Moisture)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Moistures, nil
}

func ReturnDataInfoBySlug(ctx context.Context, slug string) ([]entity.HiveInitial, error) {
	var data entity.HiveInitial
	var dataArray []entity.HiveInitial
	rows, err := infradb.Get().QueryContext(ctx, `select idexterno, Name, Slug, Description  from hive where slug = $1 `, slug)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&data.IdExterno, &data.Name, &data.Slug, &data.Description)
		dataArray = append(dataArray, data)
	}
	return dataArray, nil
}
