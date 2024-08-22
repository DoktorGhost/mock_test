package weatherAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testTask2/internal/entity"
)

type Weather struct{}

func (w Weather) FetchWeather(loc entity.Location) (entity.WeatherInfo, error) {

	accessKey := os.Getenv("ACCESS_KEY")

	// Заголовки запроса
	headers := map[string]string{
		"Content-Type":         "application/json",
		"X-Yandex-Weather-Key": accessKey,
	}

	// GraphQL-запрос
	city := loc.City
	lat, err := strconv.ParseFloat(loc.Lat, 32)
	if err != nil {
		return entity.WeatherInfo{}, err
	}

	lon, err := strconv.ParseFloat(loc.Lon, 32)
	if err != nil {
		return entity.WeatherInfo{}, err
	}

	query := fmt.Sprintf(`{
		weatherByPoint(request: { lat: %f, lon: %f }) {
			now {
				temperature
				condition
			}
		}
	}`, lat, lon)

	// Подготовка данных для отправки запроса
	body := map[string]string{
		"query": query,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return entity.WeatherInfo{}, fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Создание HTTP POST-запроса
	req, err := http.NewRequest("POST", "https://api.weather.yandex.ru/graphql/query", bytes.NewBuffer(jsonData))
	if err != nil {
		return entity.WeatherInfo{}, fmt.Errorf("error creating request: %v", err)
	}

	// Добавление заголовков
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return entity.WeatherInfo{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Чтение и вывод ответа
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.WeatherInfo{}, fmt.Errorf("error reading response: %v", err)
	}

	var weatherResp WeatherResponse

	// Декодирование JSON в структуру
	err = json.Unmarshal([]byte(responseData), &weatherResp)
	if err != nil {
		return entity.WeatherInfo{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	result := entity.WeatherInfo{
		City:      city,
		Temp:      weatherResp.Data.WeatherByPoint.Now.Temp,
		Condition: weatherResp.Data.WeatherByPoint.Now.Condition,
	}
	return result, nil
}

// Вспомогательная структура для маппинга вложенного JSON
type WeatherResponse struct {
	Data struct {
		WeatherByPoint struct {
			Now entity.WeatherInfo `json:"now"`
		} `json:"weatherByPoint"`
	} `json:"data"`
}
