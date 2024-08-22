package psg

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
	"testTask2/internal/config"
	"testTask2/internal/entity"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func InitStorage(conf *config.Config) (*PostgresStorage, error) {

	login := conf.DB_LOGIN
	password := conf.DB_PASS
	host := conf.DB_HOST
	port := conf.DB_PORT
	dbname := conf.DB_NAME

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", login, password, host, port, dbname)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Чтение schema.sql
	schema, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения schema.sql: %v", err)
	}

	// Выполнение SQL-запросов
	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения schema.sql: %v", err)
	}

	return NewPostgresStorage(db), nil
}

func (s *PostgresStorage) Save(info *entity.WeatherInfo) error {
	var id int
	query := `INSERT INTO weather (city, temp, condition) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRow(query, info.City, info.Temp, info.Condition).Scan(&id)
	if err != nil {
		return fmt.Errorf("Error save mock_API: %v", err)
	}
	log.Println("Добавлена запись. ID: ", id)
	return nil
}

func (s *PostgresStorage) Read(id int) (*entity.WeatherInfo, error) {
	var info entity.WeatherInfo
	query := `SELECT city, temp, condition FROM weather WHERE id=$1`
	err := s.db.QueryRow(query, id).Scan(&info.City, &info.Temp, &info.Condition)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (s *PostgresStorage) Update(id int, info *entity.WeatherInfo) error {
	query := `UPDATE weather SET city=$1, temp=$2, condition=$3 WHERE id=$4`
	_, err := s.db.Exec(query, info.City, info.Temp, info.Condition, id)
	return err
}

func (s *PostgresStorage) Delete(id int) error {
	query := `DELETE FROM weather WHERE id=$1`
	_, err := s.db.Exec(query, id)
	return err
}
