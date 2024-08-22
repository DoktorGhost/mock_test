package usecase

import (
	"testTask2/internal/API"
	"testTask2/internal/entity"
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

func (uc *useCase) GetСoordinates(city string) (entity.Location, error) {
	return uc.city.GetСoordinates(city)
}

func (uc *useCase) FetchWeather(loc entity.Location) (entity.WeatherInfo, error) {
	return uc.weather.FetchWeather(loc)
}

func (uc *useCase) Save(info *entity.WeatherInfo) error {
	return uc.storage.Save(info)
}

func (uc *useCase) GetAndSave(city string) error {
	loc, err := uc.GetСoordinates(city)
	if err != nil {
		return err
	}
	weather, err := uc.FetchWeather(loc)
	if err != nil {
		return err
	}
	err = uc.Save(&weather)
	if err != nil {
		return err
	}
	return nil
}
