package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/config"
	"github.com/robertoduessmann/weather-api/controller"
	v2 "github.com/robertoduessmann/weather-api/controller/v2"
)

func main() {

	weather := mux.NewRouter()
	weather.Path("/weather/{city}").Methods(http.MethodGet).HandlerFunc(controller.CurrentWeather)

	weather.
		Path("/v2/weather/{city}").
		Queries("unit", "{unit}").
		Methods(http.MethodGet).
		HandlerFunc(v2.CurrentWeather)

	weather.
		Path("/v2/weather/{city}").
		Methods(http.MethodGet).
		HandlerFunc(v2.CurrentWeather)

	if err := http.ListenAndServe(":"+config.Get().Port, handlers.CORS()(weather)); err != nil {
		log.Fatal(err)
	}

}
