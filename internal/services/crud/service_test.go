package crud_test

import (
	"database/sql"
	"fmt"
	"github.com/DoktorGhost/mock_test/internal/repositories"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"testing"

	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/suite"

	"github.com/DoktorGhost/mock_test/internal/entity"
	"github.com/DoktorGhost/mock_test/internal/services/crud"
)

type CrudServiceTestSuite struct {
	suite.Suite

	repo              crud.Repository
	db                *sql.DB
	postgresContainer testcontainers.Container
}

func (suite *CrudServiceTestSuite) SetupTest() {
	// Настройка PostgreSQL контейнера
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test_user",
			"POSTGRES_PASSWORD": "test_pas",
			"POSTGRES_DB":       "test_db",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
	}

	var err error
	suite.postgresContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	suite.Require().NoError(err)

	host, err := suite.postgresContainer.Host(ctx)
	suite.Require().NoError(err)

	port, err := suite.postgresContainer.MappedPort(ctx, "5432")
	suite.Require().NoError(err)

	dsn := fmt.Sprintf("postgres://test_user:test_pas@%s:%s/test_db?sslmode=disable", host, port.Port())
	suite.db, err = sql.Open("pgx", dsn)
	suite.Require().NoError(err)

	_, err = suite.db.Exec(`CREATE TABLE weather (
		id SERIAL PRIMARY KEY,
		city VARCHAR(255),
		temp FLOAT,
		condition VARCHAR(255)
	)`)
	suite.Require().NoError(err)

	suite.Require().NoError(err)

	suite.repo = repositories.NewPostgresRepository(suite.db)
}

func (suite *CrudServiceTestSuite) TearDownTest() {
	suite.db.Close()
	suite.postgresContainer.Terminate(context.Background())
}

func (suite *CrudServiceTestSuite) TestSave() {
	weatherInfo := &entity.WeatherInfo{
		City:      "TestCity",
		Temp:      25.5,
		Condition: "Sunny",
	}

	err := suite.repo.Save(weatherInfo)
	suite.NoError(err)

	var count int
	err = suite.db.QueryRow("SELECT COUNT(*) FROM weather WHERE city = $1", weatherInfo.City).Scan(&count)
	suite.NoError(err)
	suite.Equal(1, count)
}

func (suite *CrudServiceTestSuite) TestRead() {
	// Вставляем тестовые данные
	suite.db.Exec(`INSERT INTO weather (city, temp, condition) VALUES ('TestCity', 25.5, 'Sunny')`)

	weatherInfo, err := suite.repo.Read(1)
	suite.NoError(err)
	suite.NotNil(weatherInfo)
	suite.Equal("TestCity", weatherInfo.City)
	suite.Equal(25.5, weatherInfo.Temp)
	suite.Equal("Sunny", weatherInfo.Condition)
}

func (suite *CrudServiceTestSuite) TestUpdate() {
	// Вставляем тестовые данные
	suite.db.Exec(`INSERT INTO weather (city, temp, condition) VALUES ('TestCity', 25.5, 'Sunny')`)

	weatherInfo := &entity.WeatherInfo{
		City:      "UpdatedCity",
		Temp:      20.0,
		Condition: "Cloudy",
	}

	err := suite.repo.Update(1, weatherInfo)
	suite.NoError(err)

	updatedWeatherInfo, err := suite.repo.Read(1)
	suite.NoError(err)
	suite.Equal("UpdatedCity", updatedWeatherInfo.City)
	suite.Equal(20.0, updatedWeatherInfo.Temp)
	suite.Equal("Cloudy", updatedWeatherInfo.Condition)
}

func (suite *CrudServiceTestSuite) TestDelete() {
	// Вставляем тестовые данные
	suite.db.Exec(`INSERT INTO weather (city, temp, condition) VALUES ('TestCity', 25.5, 'Sunny')`)

	err := suite.repo.Delete(1)
	suite.NoError(err)

	var count int
	err = suite.db.QueryRow("SELECT COUNT(*) FROM weather WHERE id = 1").Scan(&count)
	suite.NoError(err)
	suite.Equal(0, count)
}

func TestCrudServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CrudServiceTestSuite))
}
