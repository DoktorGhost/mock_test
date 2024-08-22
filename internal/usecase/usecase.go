package usecase

import (
	"testTask2/internal/API"
	"testTask2/internal/storage"
)

type useCase struct {
	city    API.CityAPI
	weather API.WeatherAPI
	storage storage.Database
}

func NewUseCase(city API.CityAPI, weather API.WeatherAPI, storage storage.Database) *useCase {
	return &useCase{city: city, weather: weather, storage: storage}
}

func (uc *useCase) GetAndSave(city string) error {
	loc, err := uc.city.GetСoordinates(city)
	if err != nil {
		return err
	}
	weather, err := uc.weather.FetchWeather(loc)
	if err != nil {
		return err
	}
	err = uc.storage.Save(&weather)
	if err != nil {
		return err
	}
	return nil
}
