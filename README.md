# Написание unit тестов с моками на Go

## Описание задачи

Компания разрабатывает сервис на Go, который взаимодействует с внешними API и базой данных.
Необходимо написать модульные тесты для бизнес-логики сервиса с использованием моков.
Это позволит изолировать тестируемый код от внешних зависимостей и проверять логику работы сервиса в различных
сценариях.

## Требования

### 1. Реализовать простой сервис

- Реализовать сервис на Go, который выполняет следующие действия:
    - Запрашивает данные у внешнего API (например, информацию о погоде или курсе валют).
    - Обрабатывает полученные данные и сохраняет их в базу данных (или возвращает пользователю).

- Пример сущности, которую обрабатывает сервис:

```go
package main

type WeatherInfo struct {
	City      string
	Temp      float64
	Condition string
}

```

- Реализовать функции для работы с внешним API и базой данных:
    - FetchWeather(city string) (WeatherInfo, error)
    - SaveWeather(weather WeatherInfo) error

### 2. Настроить мокирование зависимостей

- Использовать библиотеку [gomock](https://github.com/uber-go/mock) для создания моков зависимостей сервиса.
- Создать интерфейсы для взаимодействия с внешним API и базой данных, чтобы их можно было замокать в тестах.

```go
package main

type WeatherInfo struct {
	City      string
	Temp      float64
	Condition string
}

type WeatherAPI interface {
	FetchWeather(city string) (WeatherInfo, error)
}

type Database interface {
	Save(info WeatherInfo) error
}

```

- Внедрить зависимости в сервис через конструктор. (Dependency Injection)

### 3. Написать модульные тесты

- Написать тесты, используя моки, чтобы проверить поведение сервиса в разных сценариях:
    - Успешное получение данных и их сохранение.
    - Ошибка при запросе данных у API.
    - Ошибка при сохранении данных в базу.

### 4. Обработать различные сценарии

- Покрыть тестами разные ситуации:
    - Внешний API возвращает ошибку.
    - База данных не может сохранить данные.
    - Корректность передачи данных между слоями.

## Ожидаемый результат

Предоставить код сервиса и тестов, которые корректно работают с поднятым контейнером PostgreSQL.
Все тесты должны успешно выполняться при запуске через команду go test.
Код должен быть написан с учетом хороших практик Go (структурирование пакетов, использование контекстов и т.д.).

### Дополнительные требования (необязательно)

Настроить CI (например, GitHub Actions) для автоматического запуска тестов.