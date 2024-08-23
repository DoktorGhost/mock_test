package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DoktorGhost/mock_test/internal/entity"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (s *PostgresRepository) Save(info *entity.WeatherInfo) error {
	var id int
	query := `INSERT INTO weather (city, temp, condition) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRow(query, info.City, info.Temp, info.Condition).Scan(&id)
	if err != nil {
		return fmt.Errorf("Error save mock_API: %v", err)
	}
	log.Println("Добавлена запись. ID: ", id)
	return nil
}

func (s *PostgresRepository) Read(id int) (*entity.WeatherInfo, error) {
	var info entity.WeatherInfo
	query := `SELECT city, temp, condition FROM weather WHERE id=$1`
	err := s.db.QueryRow(query, id).Scan(&info.City, &info.Temp, &info.Condition)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (s *PostgresRepository) Update(id int, info *entity.WeatherInfo) error {
	query := `UPDATE weather SET city=$1, temp=$2, condition=$3 WHERE id=$4`
	_, err := s.db.Exec(query, info.City, info.Temp, info.Condition, id)
	return err
}

func (s *PostgresRepository) Delete(id int) error {
	query := `DELETE FROM weather WHERE id=$1`
	_, err := s.db.Exec(query, id)
	return err
}
