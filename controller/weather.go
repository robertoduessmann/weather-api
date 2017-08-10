package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

	weather.Temperature = parseTemperature(doc)
	weather.Wind = parseWind(doc)
}

func parseTemperature(doc *goquery.Document) string {
	var temperature string
	doc.Find("body > pre > span:nth-child(3)").Each(func(i int, s *goquery.Selection) {
		if !strings.Contains(s.Text(), "_ - _") {
			temperature = s.Text()
		} else {
			doc.Find("body > pre > span:nth-child(2)").Each(func(i int, s *goquery.Selection) {
				temperature = s.Text()
			})
		}
	})
	return temperature
}

func parseWind(doc *goquery.Document) string {
	var wind string
	doc.Find("body > pre > span:nth-child(6)").Each(func(i int, s *goquery.Selection) {
		if !strings.Contains(s.Text(), "↑") && !strings.Contains(s.Text(), "←") {
			wind = s.Text()
		} else {
			doc.Find("body > pre > span:nth-child(7)").Each(func(i int, s *goquery.Selection) {
				wind = s.Text()
			})
		}
	})
	return wind
}

func toJSON(weather model.Weather) []byte {
	respose, err := json.Marshal(weather)
	if err != nil {
		fmt.Println(err)
	}
	return respose
}
