package v2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/model"
)

const (
	timeFormat = "2006-01-02"
	wttrURL    = "https://wttr.in"
)

// CurrentWeather gets the current weather to show in JSON format
//
// This endpoint uses wttr API with JSON response under the hood to make it
// easier to handle with units and formats
func CurrentWeather(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]
	uri := fmt.Sprintf("%s/%s?format=j1", wttrURL, city)
	res, err := http.Get(uri)

	if err != nil {
		errJSON, _ := json.Marshal(model.ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(errJSON)
		return
	}

	if res.StatusCode != http.StatusOK {
		errJSON, _ := json.Marshal(model.ErrorMessage{Message: "NOT_FOUND"})
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write(errJSON)
		return
	}

	defer res.Body.Close()

	var wttr wttrResponse
	if err := json.NewDecoder(res.Body).Decode(&wttr); err != nil {
		errJSON, _ := json.Marshal(model.ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(errJSON)
		return
	}

	var (
		cc   = wttr.CurrentCondition[0]
		unit = mux.Vars(r)["unit"]
	)

	response := model.Weather{
		Description: cc.WeatherDesc[0].Value,
		Temperature: cc.Temp(unit),
		Wind:        cc.Windspeed(unit),
	}

	for i, weather := range wttr.Weather {
		day, err := time.Parse(timeFormat, weather.Date)
		if err != nil {
			errJSON, _ := json.Marshal(model.ErrorMessage{Message: err.Error()})
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(errJSON)
			return
		}

		response.Forecast[i] = model.Forecast{
			Day:         day.Weekday().String(),
			Temperature: weather.Hourly[0].Temp(unit),
			Wind:        weather.Hourly[0].Windspeed(unit),
		}
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		errJSON, _ := json.Marshal(model.ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(errJSON)
		return
	}
}
