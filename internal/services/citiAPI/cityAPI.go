package cityAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/DoktorGhost/mock_test/internal/entity"
)

type Location struct{}

func New() Location {
	return Location{}
}

func (l Location) GetCoordinates(city string) (entity.Location, error) {

	apiURL := "https://nominatim.openstreetmap.org/search"

	// Подготовка параметров запроса
	params := url.Values{}
	params.Add("q", city)
	params.Add("format", "json")
	params.Add("limit", "1")

	// Создание URL с параметрами
	requestURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	// Выполнение HTTP-запроса
	resp, err := http.Get(requestURL)
	if err != nil {
		return entity.Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.Location{}, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	// Чтение и декодирование ответа
	var locations []entity.Location
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return entity.Location{}, fmt.Errorf("error decoding response: %v", err)
	}

	if len(locations) > 0 {
		locations[0].City = city
		return locations[0], nil
	} else {
		return entity.Location{}, fmt.Errorf("no results found for %s", city)
	}

}
