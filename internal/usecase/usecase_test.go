package usecase_test

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/mock/gomock"

	"github.com/DoktorGhost/mock_test/internal/entity"
	"github.com/DoktorGhost/mock_test/internal/services/crud"
	"github.com/DoktorGhost/mock_test/internal/usecase"
	"github.com/testcontainers/testcontainers-go"
)

type CrudServiceTestSuite struct {
	suite.Suite

	ctrl              *gomock.Controller
	db                *sql.DB
	postgresContainer testcontainers.Container
	mockCityAPI       *usecase.MockCityAPI
	mockWeatherAPI    *usecase.MockWeatherAPI
	mockStorage       *crud.MockRepository
	crudService       *crud.Service
	useCase           *usecase.UseCase
}

func (suite *CrudServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())

	suite.mockCityAPI = usecase.NewMockCityAPI(suite.ctrl)
	suite.mockWeatherAPI = usecase.NewMockWeatherAPI(suite.ctrl)
	suite.mockStorage = crud.NewMockRepository(suite.ctrl)
	suite.crudService = crud.NewService(suite.mockStorage)
	suite.useCase = usecase.NewUseCase(suite.mockCityAPI, suite.mockWeatherAPI, suite.crudService)
}

func (suite *CrudServiceTestSuite) TestGetAndSave_Success() {
	city := "Rostov-on-Don"
	location := entity.Location{City: city, Lat: "47.235", Lon: "39.700"}
	weatherInfo := entity.WeatherInfo{City: city, Temp: 18.5, Condition: "Cloudy"}

	suite.mockCityAPI.EXPECT().GetCoordinates(city).Return(location, nil)
	suite.mockWeatherAPI.EXPECT().FetchWeather(location).Return(weatherInfo, nil)
	suite.mockStorage.EXPECT().Save(&weatherInfo).Return(nil)

	err := suite.useCase.GetAndSave(city)
	suite.NoError(err)
}

func (suite *CrudServiceTestSuite) TestGetAndSave_ErrorGettingCoordinates() {
	city := "InvalidCity"
	expectedErr := errors.New("could not find city")

	suite.mockCityAPI.EXPECT().GetCoordinates(city).Return(entity.Location{}, expectedErr)

	err := suite.useCase.GetAndSave(city)
	suite.Equal(expectedErr, err)
}

func (suite *CrudServiceTestSuite) TestGetAndSave_ErrorFetchingWeather() {
	city := "Rostov-on-Don"
	location := entity.Location{City: city, Lat: "47.235", Lon: "39.700"}
	expectedErr := errors.New("failed to fetch weather")

	suite.mockCityAPI.EXPECT().GetCoordinates(city).Return(location, nil)
	suite.mockWeatherAPI.EXPECT().FetchWeather(location).Return(entity.WeatherInfo{}, expectedErr)

	err := suite.useCase.GetAndSave(city)
	suite.Equal(expectedErr, err)
}

func (suite *CrudServiceTestSuite) TestGetAndSave_ErrorSavingData() {
	city := "Rostov-on-Don"
	location := entity.Location{City: city, Lat: "47.235", Lon: "39.700"}
	weatherInfo := entity.WeatherInfo{City: city, Temp: 18.5, Condition: "Cloudy"}
	expectedErr := errors.New("failed to save data")

	suite.mockCityAPI.EXPECT().GetCoordinates(city).Return(location, nil)
	suite.mockWeatherAPI.EXPECT().FetchWeather(location).Return(weatherInfo, nil)
	suite.mockStorage.EXPECT().Save(&weatherInfo).Return(expectedErr)

	err := suite.useCase.GetAndSave(city)
	suite.Equal(expectedErr, err)
}

func TestCrudServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CrudServiceTestSuite))
}
