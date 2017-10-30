package main

import (
	"log"
	"net/http"

	"weather-api/config"
	"weather-api/controller"

	"github.com/gorilla/mux"
)

func main() {

	weather := mux.NewRouter()
	weather.Path("/weather/{city}").Methods(http.MethodGet).HandlerFunc(controller.CurrentWeather)

	if err := http.ListenAndServe(":"+config.Get().Port, weather); err != nil {
		log.Fatal(err)
	}

}
