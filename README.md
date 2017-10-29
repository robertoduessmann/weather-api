# weather-api

[![Build Status](https://travis-ci.org/robertoduessmann/weather-api.svg?branch=master)](https://travis-ci.org/robertoduessmann/weather-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/robertoduessmann/weather-api)](https://goreportcard.com/report/github.com/robertoduessmann/weather-api)
[![GoDoc](https://godoc.org/github.com/robertoduessmann/weather-api?status.svg)](https://godoc.org/github.com/robertoduessmann/weather-api)

> A REST API to check the current weather.

> https://goweather.herokuapp.com/weather/Curitiba<br />
https://goweather.herokuapp.com/weather/SaoPaulo<br />
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
```sh
{"temperature":"14 °C","wind":"6 km/h"}
```
