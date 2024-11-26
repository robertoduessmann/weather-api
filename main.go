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
	log.Println("Starting application")

	weather := mux.NewRouter()

	log.Println("Creating routes")
	weather.
		Path("/weather/{city}").
		Methods(http.MethodGet).
		HandlerFunc(controller.CurrentWeather)

	weather.
		Path("/v2/weather/{city}").
		Queries("unit", "{unit}").
		Methods(http.MethodGet).
		HandlerFunc(v2.CurrentWeather)

	weather.
		Path("/v2/weather/{city}").
		Methods(http.MethodGet).
		HandlerFunc(v2.CurrentWeather)

	log.Println("Created routes")
	port := config.Get().Port
	log.Println("Initialize server on port: " + port)
	if err := http.ListenAndServe(":"+port, handlers.CORS()(weather)); err != nil {
		log.Fatal(err)
	}
}
