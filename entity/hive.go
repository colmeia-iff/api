package entity

import "time"

type Hive struct {
	IdExterno          string `json:"idExterno"`
	Name               string `json:"name"`
	Moisture           Moisture
	OutsideTemperature OutsideTemperature
	Temperature        Temperature
	Weight             Weight
}

type HiveInitial struct {
	IdExterno   string `json:"idExterno"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
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

type ValuesNew struct {
	Values string    `json:"values"`
	Time   time.Time `json:"time"`
}
type Voltage struct {
	Data ValuesNew
}

type Vento struct {
	Data ValuesNew
}

type Melg struct {
	Data ValuesNew
}

type Resist struct {
	Data ValuesNew
}
type HiveReturn struct {
	IdExterno          string `json:"id_externo"`
	Name               string `json:"name"`
	Moisture           []Moisture
	OutsideTemperature []OutsideTemperature
	Temperature        []Temperature
	Weight             []Weight
	Resist             []Resist
	Voltage            []Voltage
	Vento              []Vento
	Melg               []Melg
}
