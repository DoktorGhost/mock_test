package entity

type WeatherInfo struct {
	City      string  `json:"city"`
	Temp      float64 `json:"temperature"`
	Condition string  `json:"condition"`
}

type Location struct {
	City string `json:"city"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}
