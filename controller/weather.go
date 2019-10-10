package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
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

	resp := getExternalWeather(getCity(r))
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
		fmt.Fprintf(w, string(toJSON(weather)))
	}

}

func getCity(r *http.Request) string {
	return mux.Vars(r)["city"]
}

func getExternalWeather(city string) *http.Response {
	resp, err := http.Get("http://wttr.in/" + city + "?m")
	if err != nil {
		log.Fatal("Cannot open url: ", err)
	}
	return resp
}

func parse(resp *http.Response, weather *model.Weather) error {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	weather.Description = parser.Parse(doc, descriptionTags)
	weather.Temperature = parser.Parse(doc, temperatureTags)
	if len(weather.Temperature) > 0 {
		weather.Temperature = weather.Temperature + " °C"
	}

	weather.Wind = parser.Parse(doc, windTags)
	if len(weather.Wind) > 0 {
		weather.Wind = weather.Wind + " km/h"
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
	if len(weather.Description) == 0 {
		return true
	}

	return false
}

func toJSON(object interface{}) []byte {
	respose, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
	}
	return respose
}
