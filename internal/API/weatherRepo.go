package API

import "testTask2/internal/entity"

type CityAPI interface {
	GetСoordinates(city string) (entity.Location, error)
}

type WeatherAPI interface {
	FetchWeather(loc entity.Location) (entity.WeatherInfo, error)
}
