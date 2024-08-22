package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testTask2/internal/entity"
	"testTask2/internal/usecase"
	"testTask2/internal/usecase/mock/mock_API"
	"testTask2/internal/usecase/mock/mock_storage"
)

func TestUseCase_GetAndSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем моки
	mockCityAPI := mock_API.NewMockCityAPI(ctrl)
	mockWeatherAPI := mock_API.NewMockWeatherAPI(ctrl)
	mockStorage := mock_storage.NewMockDatabase(ctrl)

	// Создаем экземпляр useCase
	uc := usecase.NewUseCase(mockCityAPI, mockWeatherAPI, mockStorage)

	// 1. Успешный сценарий
	t.Run("success", func(t *testing.T) {
		city := "Rostov-on-Don"
		location := entity.Location{City: city, Lat: "47.235", Lon: "39.700"}
		weatherInfo := entity.WeatherInfo{City: city, Temp: 18.5, Condition: "Cloudy"}

		// Настройка поведения моков
		mockCityAPI.EXPECT().GetСoordinates(city).Return(location, nil)
		mockWeatherAPI.EXPECT().FetchWeather(location).Return(weatherInfo, nil)
		mockStorage.EXPECT().Save(&weatherInfo).Return(nil)

		// Выполнение теста
		err := uc.GetAndSave(city)
		assert.NoError(t, err)
	})

	// 2. Ошибка при получении координат города
	t.Run("error getting coordinates", func(t *testing.T) {
		city := "InvalidCity"
		expectedErr := errors.New("could not find city")

		// Настройка поведения моков
		mockCityAPI.EXPECT().GetСoordinates(city).Return(entity.Location{}, expectedErr)

		// Выполнение теста
		err := uc.GetAndSave(city)
		assert.Equal(t, expectedErr, err)
	})

	// 3. Ошибка при получении данных о погоде
	t.Run("error fetching weather", func(t *testing.T) {
		city := "Rostov-on-Don"
		location := entity.Location{City: city, Lat: "47.235", Lon: "39.700"}
		expectedErr := errors.New("failed to fetch weather")

		// Настройка поведения моков
		mockCityAPI.EXPECT().GetСoordinates(city).Return(location, nil)
		mockWeatherAPI.EXPECT().FetchWeather(location).Return(entity.WeatherInfo{}, expectedErr)

		// Выполнение теста
		err := uc.GetAndSave(city)
		assert.Equal(t, expectedErr, err)
	})

	// 4. Ошибка при сохранении данных
	t.Run("error saving data", func(t *testing.T) {
		city := "Rostov-on-Don"
		location := entity.Location{City: city, Lat: "47.235", Lon: "39.700"}
		weatherInfo := entity.WeatherInfo{City: city, Temp: 18.5, Condition: "Cloudy"}
		expectedErr := errors.New("failed to save data")

		// Настройка поведения моков
		mockCityAPI.EXPECT().GetСoordinates(city).Return(location, nil)
		mockWeatherAPI.EXPECT().FetchWeather(location).Return(weatherInfo, nil)
		mockStorage.EXPECT().Save(&weatherInfo).Return(expectedErr)

		// Выполнение теста
		err := uc.GetAndSave(city)
		assert.Equal(t, expectedErr, err)
	})
}
