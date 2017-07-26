package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/controller"
)

func main() {

	weather := mux.NewRouter()
	weather.Path("/weather/{city}").Methods(http.MethodGet).HandlerFunc(controller.CurrentWeather)

	if err := http.ListenAndServe(":3000", weather); err != nil {
		log.Fatal(err)
	}

}
