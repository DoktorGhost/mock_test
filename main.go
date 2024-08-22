package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"testTask2/internal/API/cityAPI"
	"testTask2/internal/API/weatherAPI"
	"testTask2/internal/config"
	"testTask2/internal/storage/psg"
	"testTask2/internal/usecase"
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

	db, err := psg.InitStorage(conf)
	if err != nil {
		log.Fatal("Ошибка подключения к бд:", err)
	}
	log.Println("База данных запущена")

	coordinate := usecase.NewUseCase(cityAPI.Location{}, weatherAPI.Weather{}, db)

	citys := []string{"Rostov-on-Don", "Moscow", "Bergen", "Amsterdam", ""}
	for _, city := range citys {
		err = coordinate.GetAndSave(city)
		if err != nil {
			fmt.Println(err)
		}
	}
}
