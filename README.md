# weather-api

[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build Status](https://travis-ci.com/robertoduessmann/weather-api.svg?branch=master)](https://travis-ci.com/robertoduessmann/weather-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/robertoduessmann/weather-api)](https://goreportcard.com/report/github.com/robertoduessmann/weather-api)
[![GoDoc](https://godoc.org/github.com/robertoduessmann/weather-api?status.svg)](https://godoc.org/github.com/robertoduessmann/weather-api)

> A REST API to check the current weather.

> https://goweather.herokuapp.com/weather/Curitiba<br />
https://goweather.herokuapp.com/weather/{city}

## Build locally (Mac users)
```sh
brew install dep
dep ensure
go build
```
## Run
```sh
./weather-api
```
## Usage
```sh
curl http://localhost:3000/weather/{city}
```
## Example
#### Request
```sh
curl http://localhost:3000/weather/Curitiba
```
#### Response
```json
{  
   "temperature":"29 °C",
   "wind":"20 km/h",
   "description":"Partly cloudy",
   "forecast":[  
      {  
         "day":1,
         "temperature":"27 °C",
         "wind":"12 km/h"
      },
      {  
         "day":2,
         "temperature":"22 °C",
         "wind":"8 km/h"
      }
   ]
}
```
## Web version
A web client of the API is also available: https://reacttempo.netlify.app/ <br />
The project can be found in https://github.com/GabrielCampos99/appTempo.

## License
The MIT License ([MIT](https://github.com/robertoduessmann/weather-api/blob/master/LICENSE))
