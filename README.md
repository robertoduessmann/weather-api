# weather-api

[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build Status](https://travis-ci.org/robertoduessmann/weather-api.svg?branch=master)](https://travis-ci.org/robertoduessmann/weather-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/robertoduessmann/weather-api)](https://goreportcard.com/report/github.com/robertoduessmann/weather-api)
[![GoDoc](https://godoc.org/github.com/robertoduessmann/weather-api?status.svg)](https://godoc.org/github.com/robertoduessmann/weather-api)

> A REST API to check the current weather.

> https://goweather.herokuapp.com/weather/Curitiba<br />
https://goweather.herokuapp.com/weather/{city}

## Build
```sh
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
