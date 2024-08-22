package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testTask2/internal/entity"
	"testTask2/internal/usecase/mock/mock_API"
	"testTask2/internal/usecase/mock/mock_storage"
	"testing"
)

func TestGetCoordinates(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание моков
	mockCityAPI := mock_API.NewMockCityAPI(ctrl)
	mockWeatherAPI := mock_API.NewMockWeatherAPI(ctrl)
	mockStorage := mock_storage.NewMockDatabase(ctrl)

	// Ожидания
	expectedLocation := entity.Location{
		City: "Rostov-on-Don",
		Lat:  "47.235",
		Lon:  "39.700",
	}

	// Настройка поведения моков
	mockCityAPI.EXPECT().GetСoordinates("Rostov-on-Don").Return(expectedLocation, nil)

	// Создание экземпляра useCase
	uc := NewUseCase(mockCityAPI, mockWeatherAPI, mockStorage)

	// Выполнение теста
	location, err := uc.GetСoordinates("Rostov-on-Don")
	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, location)
}

func TestFetchWeather(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание моков
	mockCityAPI := mock_API.NewMockCityAPI(ctrl)
	mockWeatherAPI := mock_API.NewMockWeatherAPI(ctrl)
	mockStorage := mock_storage.NewMockDatabase(ctrl)

	// Ожидания
	location := entity.Location{
		City: "Rostov-on-Don",
		Lat:  "47.235",
		Lon:  "39.700",
	}
	expectedWeatherInfo := entity.WeatherInfo{
		City:      "Rostov-on-Don",
		Temp:      18.5,
		Condition: "Cloudy",
	}

	// Настройка поведения моков
	mockWeatherAPI.EXPECT().FetchWeather(location).Return(expectedWeatherInfo, nil)

	// Создание экземпляра useCase
	uc := NewUseCase(mockCityAPI, mockWeatherAPI, mockStorage)

	// Выполнение теста
	weatherInfo, err := uc.FetchWeather(location)
	assert.NoError(t, err)
	assert.Equal(t, expectedWeatherInfo, weatherInfo)
}

func TestSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание моков
	mockCityAPI := mock_API.NewMockCityAPI(ctrl)
	mockWeatherAPI := mock_API.NewMockWeatherAPI(ctrl)
	mockStorage := mock_storage.NewMockDatabase(ctrl)

	// Ожидания
	weatherInfo := entity.WeatherInfo{
		City:      "Rostov-on-Don",
		Temp:      18.5,
		Condition: "Cloudy",
	}

	// Настройка поведения моков
	mockStorage.EXPECT().Save(&weatherInfo).Return(nil)

	// Создание экземпляра useCase
	uc := NewUseCase(mockCityAPI, mockWeatherAPI, mockStorage)

	// Выполнение теста
	err := uc.Save(&weatherInfo)
	assert.NoError(t, err)
}
