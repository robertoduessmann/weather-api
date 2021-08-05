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

type wttrResponse struct {
	CurrentCondition []currentCondition `json:"current_condition"`
	Request          []struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	} `json:"request"`
	Weather []struct {
		Astronomy []struct {
			MoonIllumination string `json:"moon_illumination"`
			MoonPhase        string `json:"moon_phase"`
			Moonrise         string `json:"moonrise"`
			Moonset          string `json:"moonset"`
			Sunrise          string `json:"sunrise"`
			Sunset           string `json:"sunset"`
		} `json:"astronomy"`
		Date        string   `json:"date"`
		Hourly      []hourly `json:"hourly"`
		MaxtempC    string   `json:"maxtempC"`
		MaxtempF    string   `json:"maxtempF"`
		MintempC    string   `json:"mintempC"`
		MintempF    string   `json:"mintempF"`
		SunHour     string   `json:"sunHour"`
		TotalSnowCm string   `json:"totalSnow_cm"`
		UvIndex     string   `json:"uvIndex"`
	} `json:"weather"`
}

type currentCondition struct {
	FeelsLikeC      string `json:"FeelsLikeC"`
	FeelsLikeF      string `json:"FeelsLikeF"`
	Cloudcover      string `json:"cloudcover"`
	Humidity        string `json:"humidity"`
	ObservationTime string `json:"observation_time"`
	PrecipMM        string `json:"precipMM"`
	Pressure        string `json:"pressure"`
	TempC           string `json:"temp_C"`
	TempF           string `json:"temp_F"`
	UvIndex         int    `json:"uvIndex"`
	Visibility      string `json:"visibility"`
	WeatherCode     string `json:"weatherCode"`
	WeatherDesc     []struct {
		Value string `json:"value"`
	} `json:"weatherDesc"`
	WeatherIconURL []struct {
		Value string `json:"value"`
	} `json:"weatherIconUrl"`
	Winddir16Point string `json:"winddir16Point"`
	WinddirDegree  string `json:"winddirDegree"`
	WindspeedKmph  string `json:"windspeedKmph"`
	WindspeedMiles string `json:"windspeedMiles"`
}

func (cc currentCondition) Temp(unit string) string {
	switch unit {
	case "u":
		return cc.TempF + " °F"
	default:
		return cc.TempC + " °C"
	}
}

func (cc currentCondition) Windspeed(unit string) string {
	switch unit {
	case "u":
		return cc.WindspeedMiles + " mph"
	default:
		return cc.WindspeedKmph + " km/h"
	}
}

type hourly struct {
	DewPointC        string `json:"DewPointC"`
	DewPointF        string `json:"DewPointF"`
	FeelsLikeC       string `json:"FeelsLikeC"`
	FeelsLikeF       string `json:"FeelsLikeF"`
	HeatIndexC       string `json:"HeatIndexC"`
	HeatIndexF       string `json:"HeatIndexF"`
	WindChillC       string `json:"WindChillC"`
	WindChillF       string `json:"WindChillF"`
	WindGustKmph     string `json:"WindGustKmph"`
	WindGustMiles    string `json:"WindGustMiles"`
	Chanceoffog      string `json:"chanceoffog"`
	Chanceoffrost    string `json:"chanceoffrost"`
	Chanceofhightemp string `json:"chanceofhightemp"`
	Chanceofovercast string `json:"chanceofovercast"`
	Chanceofrain     string `json:"chanceofrain"`
	Chanceofremdry   string `json:"chanceofremdry"`
	Chanceofsnow     string `json:"chanceofsnow"`
	Chanceofsunshine string `json:"chanceofsunshine"`
	Chanceofthunder  string `json:"chanceofthunder"`
	Chanceofwindy    string `json:"chanceofwindy"`
	Cloudcover       string `json:"cloudcover"`
	Humidity         string `json:"humidity"`
	PrecipMM         string `json:"precipMM"`
	Pressure         string `json:"pressure"`
	TempC            string `json:"tempC"`
	TempF            string `json:"tempF"`
	Time             string `json:"time"`
	UvIndex          string `json:"uvIndex"`
	Visibility       string `json:"visibility"`
	WeatherCode      string `json:"weatherCode"`
	WeatherDesc      []struct {
		Value string `json:"value"`
	} `json:"weatherDesc"`
	WeatherIconURL []struct {
		Value string `json:"value"`
	} `json:"weatherIconUrl"`
	Winddir16Point string `json:"winddir16Point"`
	WinddirDegree  string `json:"winddirDegree"`
	WindspeedKmph  string `json:"windspeedKmph"`
	WindspeedMiles string `json:"windspeedMiles"`
}

func (h hourly) Temp(unit string) string {
	switch unit {
	case "u":
		return h.TempF + " °F"
	default:
		return h.TempC + " °C"
	}
}

func (h hourly) Windspeed(unit string) string {
	switch unit {
	case "u":
		return h.WindspeedMiles + " mph"
	default:
		return h.WindspeedKmph + " km/h"
	}
}

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
