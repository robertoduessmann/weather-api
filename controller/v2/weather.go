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

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(model.ErrorMessage{Message: message})
}

func CurrentWeather(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]
	unit := mux.Vars(r)["unit"]
	uri := fmt.Sprintf("%s/%s?format=j1", wttrURL, city)

	res, err := http.Get(uri)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		writeError(w, http.StatusNotFound, "NOT_FOUND")
		return
	}

	var wttr wttrResponse
	if err := json.NewDecoder(res.Body).Decode(&wttr); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Safety checks
	if len(wttr.CurrentCondition) == 0 || len(wttr.CurrentCondition[0].WeatherDesc) == 0 {
		writeError(w, http.StatusInternalServerError, "Unexpected API response")
		return
	}

	cc := wttr.CurrentCondition[0]
	response := model.Weather{
		Description: cc.WeatherDesc[0].Value,
		Temperature: cc.Temp(unit),
		Wind:        cc.Windspeed(unit),
		Forecast:    make([]model.Forecast, len(wttr.Weather)),
	}

	for i, weather := range wttr.Weather {
		if len(weather.Hourly) == 0 {
			writeError(w, http.StatusInternalServerError, "Incomplete forecast data")
			return
		}

		day, err := time.Parse(timeFormat, weather.Date)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.Forecast[i] = model.Forecast{
			Day:         day.Weekday().String(),
			Temperature: weather.Hourly[0].Temp(unit),
			Wind:        weather.Hourly[0].Windspeed(unit),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}
