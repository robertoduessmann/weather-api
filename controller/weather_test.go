package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/model"
)

// stubFetcher implements a controllable response for tests.
type stubFetcher struct{ }

func (stubFetcher) Fetch(city string) (*http.Response, error) {
	if city == "Curitiba" {
		html := `<!doctype html><html><body><pre>Header\n<span>25</span>\nDescription for Curitiba\n<span>10</span></pre></body></html>`
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(html)),
		}, nil
	}
	// Simulate not found / upstream error for other cities used in tests
	return nil, nil
}

func TestCurrentWeather(t *testing.T) {
	// inject stub fetcher
	old := weatherFetcher
	weatherFetcher = stubFetcher{}
	defer func() { weatherFetcher = old }()

	var weather model.Weather
	var blankWeather model.Weather

	router := new(mux.Router)
	router.HandleFunc("/weather/{city}", CurrentWeather)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/weather/Curitiba", nil)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected request to be OK. got=%d body=%s", w.Code, w.Body.String())
		t.FailNow()
	}

	if err := json.Unmarshal(w.Body.Bytes(), &weather); err != nil {
		t.Errorf("Unexpected error while unmarshal response: %v", err)
		t.FailNow()
	}

	if weather == blankWeather {
		t.Errorf("Expected weather information NOT to be empty.")
		t.FailNow()
	}
}

func TestNotFoundWeather(t *testing.T) {
	// inject stub fetcher
	old := weatherFetcher
	weatherFetcher = stubFetcher{}
	defer func() { weatherFetcher = old }()

	router := new(mux.Router)
	router.HandleFunc("/weather/{city}", CurrentWeather)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/weather/SapPaulo", nil)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected request to be NotFound. got=%d body=%s", w.Code, w.Body.String())
		t.FailNow()
	}
}

func BenchmarkCurrentWeather(b *testing.B) {
	router := new(mux.Router)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.HandleFunc("/weather/{city}", CurrentWeather)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/weather/Berlin", nil)
		router.ServeHTTP(w, r)
	}
}
