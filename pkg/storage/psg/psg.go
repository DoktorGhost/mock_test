package psg

import (
	"database/sql"
	"fmt"
	"github.com/DoktorGhost/mock_test/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{DB: db}
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
