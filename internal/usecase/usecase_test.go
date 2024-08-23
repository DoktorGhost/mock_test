package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/DoktorGhost/mock_test/internal/entity"
	"github.com/DoktorGhost/mock_test/internal/usecase"
)

func TestUseCase_GetAndSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer t.Cleanup(ctrl.Finish)

	// Создаем моки
	mockCityAPI := usecase.NewMockCityAPI(ctrl)
	mockWeatherAPI := usecase.NewMockWeatherAPI(ctrl)
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
