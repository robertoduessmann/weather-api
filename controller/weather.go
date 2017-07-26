package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/model"
)

func CurrentWeather(w http.ResponseWriter, r *http.Request) {

	var weather model.Weather

	resp := getExternalWeather(getCity(r))
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		parse(resp, &weather)
	}

	fmt.Fprintf(w, string(toJSON(weather)))
}

func getCity(r *http.Request) string {
	return mux.Vars(r)["city"]
}

func getExternalWeather(city string) *http.Response {
	resp, err := http.Get("http://wttr.in/" + city)
	if err != nil {
		log.Fatal("Cannot open url: ", err)
	}
	return resp
}

func parse(resp *http.Response, weather *model.Weather) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("body > pre > span:nth-child(3)").Each(func(i int, s *goquery.Selection) {
		weather.Temperature = s.Text()
	})
	doc.Find("body > pre > span:nth-child(6)").Each(func(i int, s *goquery.Selection) {
		weather.Wind = s.Text()
	})
}

func toJSON(weather model.Weather) []byte {
	respose, err := json.Marshal(weather)
	if err != nil {
		fmt.Println(err)
	}
	return respose
}
