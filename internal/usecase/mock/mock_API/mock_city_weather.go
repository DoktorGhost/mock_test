// Code generated by MockGen. DO NOT EDIT.
// Source: weatherRepo.go

// Package mock_API is a generated GoMock package.
package mock_API

import (
	reflect "reflect"
	entity "testTask2/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockCityAPI is a mock of CityAPI interface.
type MockCityAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCityAPIMockRecorder
}

// MockCityAPIMockRecorder is the mock recorder for MockCityAPI.
type MockCityAPIMockRecorder struct {
	mock *MockCityAPI
}

// NewMockCityAPI creates a new mock instance.
func NewMockCityAPI(ctrl *gomock.Controller) *MockCityAPI {
	mock := &MockCityAPI{ctrl: ctrl}
	mock.recorder = &MockCityAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCityAPI) EXPECT() *MockCityAPIMockRecorder {
	return m.recorder
}

// GetСoordinates mocks base method.
func (m *MockCityAPI) GetСoordinates(city string) (entity.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetСoordinates", city)
	ret0, _ := ret[0].(entity.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetСoordinates indicates an expected call of GetСoordinates.
func (mr *MockCityAPIMockRecorder) GetСoordinates(city interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetСoordinates", reflect.TypeOf((*MockCityAPI)(nil).GetСoordinates), city)
}

// MockWeatherAPI is a mock of WeatherAPI interface.
type MockWeatherAPI struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherAPIMockRecorder
}

// MockWeatherAPIMockRecorder is the mock recorder for MockWeatherAPI.
type MockWeatherAPIMockRecorder struct {
	mock *MockWeatherAPI
}

// NewMockWeatherAPI creates a new mock instance.
func NewMockWeatherAPI(ctrl *gomock.Controller) *MockWeatherAPI {
	mock := &MockWeatherAPI{ctrl: ctrl}
	mock.recorder = &MockWeatherAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeatherAPI) EXPECT() *MockWeatherAPIMockRecorder {
	return m.recorder
}

// FetchWeather mocks base method.
func (m *MockWeatherAPI) FetchWeather(loc entity.Location) (entity.WeatherInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchWeather", loc)
	ret0, _ := ret[0].(entity.WeatherInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchWeather indicates an expected call of FetchWeather.
func (mr *MockWeatherAPIMockRecorder) FetchWeather(loc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchWeather", reflect.TypeOf((*MockWeatherAPI)(nil).FetchWeather), loc)
}
