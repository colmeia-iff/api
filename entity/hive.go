package entity

import "time"

type Hive struct {
	IdExterno          string `json:"id_externo"`
	Name               string `json:"name"`
	Moisture           Moisture
	OutsideTemperature OutsideTemperature
	Temperature        Temperature
	Weight             Weight
}

type Moisture struct {
	Data Values
}
type OutsideTemperature struct {
	Data Values
}

type Temperature struct {
	Data Values
}

type Weight struct {
	Value string    `json:"value"`
	Time  time.Time `json:"time"`
}
type Values struct {
	Temp string    `json:"temp"`
	Time time.Time `json:"time"`
}

type HiveReturn struct {
	IdExterno          string `json:"id_externo"`
	Name               string `json:"name"`
	Moisture           []Moisture
	OutsideTemperature []OutsideTemperature
	Temperature        []Temperature
	Weight             []Weight
}
