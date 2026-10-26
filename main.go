package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Estrutura do retorno da API OpenWeatherMap
type OpenWeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

type WeatherService struct {
	APIKey     string
	HTTPClient *http.Client
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second, // Timeout para evitar requisições presas
		},
	}
}

func (s *WeatherService) FetchWeather(ctx context.Context, city, unit string) (*OpenWeatherResponse, error) {
	// Mapeia unidades para o padrão da OpenWeatherMap (metric / imperial)
	unitsParam := "metric"
	if unit == "fahrenheit" {
		unitsParam = "imperial"
	}

	endpoint := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s",
		url.QueryEscape(city), unitsParam, s.APIKey,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar API externa: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("cidade não encontrada")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta inesperada da API: status %d", resp.StatusCode)
	}

	var data OpenWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return &data, nil
}
