package models

type Rate struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ValuteID string
	NumCode  string  `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Nominal  int     `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}
