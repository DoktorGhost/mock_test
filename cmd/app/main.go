package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/DoktorGhost/mock_test/internal/config"
	"github.com/DoktorGhost/mock_test/internal/repositories"
	"github.com/DoktorGhost/mock_test/internal/services/citiAPI"
	"github.com/DoktorGhost/mock_test/internal/services/crud"
	"github.com/DoktorGhost/mock_test/internal/services/weatherAPI"
	"github.com/DoktorGhost/mock_test/internal/usecase"
	"github.com/DoktorGhost/mock_test/pkg/storage/psg"
)

func main() {

	//считываем .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env", err)
	}

	//парсим переменные окружения в conf
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	pgsqlConnector, err := psg.InitStorage(conf)
	if err != nil {
		log.Fatal("Ошибка подключения к бд:", err)
	}
	log.Println("База данных запущена")

	crudRepo := repositories.NewPostgresRepository(pgsqlConnector.DB)
	crudService := crud.NewService(crudRepo)
	coordinate := usecase.NewUseCase(cityAPI.New(), weatherAPI.New(), crudService)

	citys := []string{"Rostov-on-Don", "Moscow", "Bergen", "Amsterdam", ""}
	for _, city := range citys {
		err = coordinate.GetAndSave(city)
		if err != nil {
			fmt.Println(err)
		}
	}
}
