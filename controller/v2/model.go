package v2

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
	UvIndex         string `json:"uvIndex"`
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
