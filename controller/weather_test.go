```go
package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/model"
)

// setupWeatherRouter cria um router isolado para cada teste.
func setupWeatherRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(
		"/weather/{city}",
		CurrentWeather,
	).Methods(http.MethodGet)

	return router
}

// TestCurrentWeather verifica se uma cidade válida retorna
// uma resposta HTTP 200 contendo informações meteorológicas.
func TestCurrentWeather(t *testing.T) {
	t.Parallel()

	router := setupWeatherRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/weather/Curitiba",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	var weather model.Weather

	if err := json.NewDecoder(rec.Body).Decode(&weather); err != nil {
		t.Fatalf(
			"failed to decode weather response: %v",
			err,
		)
	}

	if weather == (model.Weather{}) {
		t.Fatal("expected weather information, got empty response")
	}
}

// TestNotFoundWeather verifica se uma cidade inexistente
// retorna HTTP 404.
func TestNotFoundWeather(t *testing.T) {
	t.Parallel()

	router := setupWeatherRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/weather/SapPaulo",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
			rec.Code,
		)
	}
}

// TestCurrentWeatherMethodNotAllowed verifica se métodos
// diferentes de GET são rejeitados.
func TestCurrentWeatherMethodNotAllowed(t *testing.T) {
	t.Parallel()

	router := setupWeatherRouter()

	req := httptest.NewRequest(
		http.MethodPost,
		"/weather/Curitiba",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusMethodNotAllowed,
			rec.Code,
		)
	}
}

// TestCurrentWeatherInvalidRoute verifica se uma rota
// inexistente retorna HTTP 404.
func TestCurrentWeatherInvalidRoute(t *testing.T) {
	t.Parallel()

	router := setupWeatherRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/invalid-route",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
			rec.Code,
		)
	}
}

// BenchmarkCurrentWeather mede apenas o processamento
// da requisição pelo router e controller.
//
// A rota é registrada uma única vez, antes do loop,
// evitando distorção do benchmark.
func BenchmarkCurrentWeather(b *testing.B) {
	router := setupWeatherRouter()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(
			http.MethodGet,
			"/weather/Berlin",
			nil,
		)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)
	}
}
```
