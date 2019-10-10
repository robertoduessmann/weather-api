package v2

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/model"
)

const wttrURL = "http://wttr.in"

type wttrResponse struct {
	CurrentCondition []struct {
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
	} `json:"current_condition"`
	Request []struct {
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
		Date   string `json:"date"`
		Hourly []struct {
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
		} `json:"hourly"`
		MaxtempC    string `json:"maxtempC"`
		MaxtempF    string `json:"maxtempF"`
		MintempC    string `json:"mintempC"`
		MintempF    string `json:"mintempF"`
		SunHour     string `json:"sunHour"`
		TotalSnowCm string `json:"totalSnow_cm"`
		UvIndex     string `json:"uvIndex"`
	} `json:"weather"`
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
		cc = wttr.CurrentCondition[0]
	)

	response := model.Weather{
		Description: cc.WeatherDesc[0].Value,
		Temperature: cc.FeelsLikeC + " °C",
		Wind:        cc.WindspeedKmph + " km/h",
		Forecast: [3]model.Forecast{
			{
				Day:         wttr.Weather[0].Date,
				Temperature: wttr.Weather[0].MaxtempC + " °C",
				Wind:        wttr.Weather[0].Hourly[0].WindspeedKmph + " km/h",
			},
			{
				Day:         wttr.Weather[1].Date,
				Temperature: wttr.Weather[1].MaxtempC + " °C",
				Wind:        wttr.Weather[2].Hourly[0].WindspeedKmph + " km/h",
			},
			{
				Day:         wttr.Weather[2].Date,
				Temperature: wttr.Weather[2].MaxtempC + " °C",
				Wind:        wttr.Weather[2].Hourly[0].WindspeedKmph + " km/h",
			},
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		errJSON, _ := json.Marshal(model.ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(errJSON)
		return
	}
}
