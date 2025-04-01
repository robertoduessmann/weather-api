package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/cache"
	"github.com/robertoduessmann/weather-api/model"
	"github.com/robertoduessmann/weather-api/parser"
)

var temperatureTags = []string{"body > pre > span:nth-child(3)", "body > pre > span:nth-child(2)"}
var windTags = []string{"body > pre > span:nth-child(6)", "body > pre > span:nth-child(7)"}
var descriptionTags = []string{"body > pre"}
var temperatureForecastTags = [3][]string{{"body > pre >span:nth-child(17)", "body > pre > span:nth-child(16)"},
	{"body > pre >span:nth-child(55)", "body > pre > span:nth-child(54)"},
	{"body > pre >span:nth-child(91)", "body > pre > span:nth-child(90)"}}
var windForecastTags = [3][]string{{"body > pre >span:nth-child(31)", "body > pre > span:nth-child(30)", "body > pre >span:nth-child(32)"},
	{"body > pre >span:nth-child(68)", "body > pre > span:nth-child(67)", "body > pre > span:nth-child(69)"},
	{"body > pre >span:nth-child(105)", "body > pre > span:nth-child(104)", "body > pre > span:nth-child(106)"}}

// CurrentWeather gets the current weather to show in JSON format
func CurrentWeather(w http.ResponseWriter, r *http.Request) {
	var weather model.Weather
	var err error

	city := getCity(r)
	cacheKey := fmt.Sprintf("html-%s", city)

	cacheManager := cache.NewCacheManager()
	weatherCache := cacheManager.NewCache("weather-html", 2*time.Minute)

	if cached, found := weatherCache.Get(cacheKey); found {
		log.Printf("[CACHE HIT] key=%s", cacheKey)
		w.Header().Set("Content-Type", "application/json")
		w.Write(cached.([]byte))
		return
	}

	log.Printf("[CACHE MISS] key=%s", cacheKey)
	resp := getExternalWeather(city)
	if resp == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, string(toJSON(model.ErrorMessage{Message: "NOT_FOUND"})))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		err = parse(resp, &weather)
	}

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, string(toJSON(model.ErrorMessage{Message: "NOT_FOUND"})))
	} else {
		w.WriteHeader(http.StatusOK)
		result := toJSON(weather)
		weatherCache.Put(cacheKey, result)
		w.Write(result)
	}
}

func getCity(r *http.Request) string {
	return mux.Vars(r)["city"]
}

func getExternalWeather(city string) *http.Response {
	url := "https://wttr.in/" + city + "?m"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Warning: failed to fetch weather for %s: %v", city, err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Warning: weather API for %s returned status %d", city, resp.StatusCode)
		resp.Body.Close()
		return nil
	}
	return resp
}

func parse(resp *http.Response, weather *model.Weather) error {
	if resp == nil {
		return errors.New("response is nil")
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing response body: %v", err)
		return err
	}

	weather.Description = parser.Parse(doc, descriptionTags)
	weather.Temperature = parser.Parse(doc, temperatureTags)
	if len(weather.Temperature) > 0 {
		weather.Temperature += " °C"
	}

	weather.Wind = parser.Parse(doc, windTags)
	if len(weather.Wind) > 0 {
		weather.Wind += " km/h"
	}

	if notFound(weather) {
		return errors.New("NOT_FOUND")
	}

	for i := range weather.Forecast {
		weather.Forecast[i].Day = strconv.Itoa(i + 1)
		weather.Forecast[i].Temperature = parser.Parse(doc, temperatureForecastTags[i]) + " °C"
		weather.Forecast[i].Wind = parser.Parse(doc, windForecastTags[i]) + " km/h"
	}

	return nil
}

func notFound(weather *model.Weather) bool {
	return len(weather.Description) == 0
}

func toJSON(object interface{}) []byte {
	response, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
	}
	return response
}
