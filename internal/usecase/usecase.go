package usecase

import (
	"github.com/DoktorGhost/mock_test/internal/entity"
	"github.com/DoktorGhost/mock_test/internal/services/crud"
)

//go:generate mockgen -source=$GOFILE -destination=./mock_clients.go -package=${GOPACKAGE}

type CityAPI interface {
	GetCoordinates(city string) (entity.Location, error)
}

type WeatherAPI interface {
	FetchWeather(loc entity.Location) (entity.WeatherInfo, error)
}

type UseCase struct {
	city        CityAPI
	weather     WeatherAPI
	crudService *crud.Service
}

func NewUseCase(
	city CityAPI,
	weather WeatherAPI,
	crudService *crud.Service,
) *UseCase {
	return &UseCase{
		city:        city,
		weather:     weather,
		crudService: crudService,
	}
}

func (uc *UseCase) GetAndSave(city string) error {
	loc, err := uc.city.GetCoordinates(city)
	if err != nil {
		return err
	}
	weather, err := uc.weather.FetchWeather(loc)
	if err != nil {
		return err
	}
	err = uc.crudService.Create(&weather)
	if err != nil {
		return err
	}
	return nil
}
