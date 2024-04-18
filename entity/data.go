package entity

type Data struct {
	Moisture           Moisture
	OutsideTemperature OutsideTemperature
	Temperature        Temperature
	Weight             Weight
}

type InfoData struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	IdExterno   string `json:"idExterno"`
	Description string `json:"description"`
}

type DataInfo struct {
	Name               string `json:"name"`
	Slug               string `json:"slug"`
	Moisture           []Moisture
	OutsideTemperature []OutsideTemperature
	Temperature        []Temperature
	Weight             []Weight
}
