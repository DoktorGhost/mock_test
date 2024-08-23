package crud

import (
	"github.com/DoktorGhost/mock_test/internal/entity"
)

//go:generate mockgen -source=$GOFILE -destination=./mock_repository.go -package=${GOPACKAGE}

type Repository interface {
	Save(info *entity.WeatherInfo) error
	Read(id int) (*entity.WeatherInfo, error)
	Update(id int, info *entity.WeatherInfo) error
	Delete(id int) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(info *entity.WeatherInfo) error {
	return s.repo.Save(info)
}
