package storage

import "testTask2/internal/entity"

type Database interface {
	Save(info *entity.WeatherInfo) error
}
