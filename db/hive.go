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

func Moisture(ctx context.Context, data entity.Moisture, idExterno string) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Moisture (IdExterno, Temp, Time) VALUES ($1, $2, NOW());`, idExterno, data.Data.Temp)
	if err != nil {
		return err
	}
	return nil
}

func Vento(ctx context.Context, data entity.Vento, idExterno string) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Vento (IdExterno, Temp, Time) VALUES ($1, $2, NOW());`, idExterno, data.Data.Time)
	if err != nil {
		return err
	}
	return nil
}
func Melg(ctx context.Context, data entity.Melg, idExterno string) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Melg (IdExterno, Temp, Time) VALUES ($1, $2, NOW());`, idExterno, data.Data.Time)
	if err != nil {
		return err
	}
	return nil
}
func Volt(ctx context.Context, data entity.Voltage, idExterno string) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Voltage (IdExterno, Temp, Time) VALUES ($1, $2, NOW());`, idExterno, data.Data.Time)
	if err != nil {
		return err
	}
	return nil
}
func Resist(ctx context.Context, data entity.Resist, idExterno string) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Resist (IdExterno, Temp, Time) VALUES ($1, $2, NOW());`, idExterno, data.Data.Time)
	if err != nil {
		return err
	}
	return nil
}
func OutsideTemperature(ctx context.Context, data entity.OutsideTemperature, id string) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO OutsideTemperature (IdExterno, Temp, Time) VALUES ($1, $2, NOW());`, id, data.Data.Temp)
	if err != nil {
		return err
	}
	return nil
}

func Temperature(ctx context.Context, data entity.Temperature, id string) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO Temperature (IdExterno, Temp, Time) VALUES ($1, $2, NOW());`, id, data.Data.Temp)
	if err != nil {
		return err
	}
	return nil
}

func WeightHive(ctx context.Context, data entity.Weight, id string) error {
	_, err := infradb.Get().ExecContext(ctx, `INSERT INTO WeightHive (IdExterno, Value, Time) VALUES ($1, $2,NOW());`, id, data.Value)
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
		SELECT  Temp, Time 
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
		SELECT  Temp, Time 
		FROM OutsideTemperature 
		WHERE idexterno = $1;
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
	SELECT Value, Time 
	FROM WeightHive 
	WHERE IdExterno = $1;
	
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
		SELECT  Temp, Time 
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

func GetVento(ctx context.Context, id string) ([]entity.Vento, error) {
	rows, err := infradb.Get().QueryContext(ctx, `
		SELECT  Temp, Time 
		FROM Vento 
		WHERE idExterno = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vento []entity.Vento
	for rows.Next() {
		var ventos entity.Vento
		if err := rows.Scan(&ventos.Data.Time, &ventos.Data.Time); err != nil {
			return nil, err
		}
		vento = append(vento, ventos)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vento, nil
}

func GetMelg(ctx context.Context, id string) ([]entity.Melg, error) {
	rows, err := infradb.Get().QueryContext(ctx, `
		SELECT  Temp, Time 
		FROM Melg 
		WHERE idExterno = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var melgs []entity.Melg
	for rows.Next() {
		var melg entity.Melg
		if err := rows.Scan(&melg.Data.Time, &melg.Data.Time); err != nil {
			return nil, err
		}
		melgs = append(melgs, melg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return melgs, nil
}

func GetVolt(ctx context.Context, id string) ([]entity.Voltage, error) {
	rows, err := infradb.Get().QueryContext(ctx, `
		SELECT  Temp, Time 
		FROM Voltage 
		WHERE idExterno = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vento []entity.Voltage
	for rows.Next() {
		var ventos entity.Voltage
		if err := rows.Scan(&ventos.Data.Time, &ventos.Data.Time); err != nil {
			return nil, err
		}
		vento = append(vento, ventos)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vento, nil
}

func GetResist(ctx context.Context, id string) ([]entity.Resist, error) {
	rows, err := infradb.Get().QueryContext(ctx, `
		SELECT  Temp, Time 
		FROM Resist 
		WHERE idExterno = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vento []entity.Resist
	for rows.Next() {
		var ventos entity.Resist
		if err := rows.Scan(&ventos.Data.Time, &ventos.Data.Time); err != nil {
			return nil, err
		}
		vento = append(vento, ventos)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vento, nil
}
