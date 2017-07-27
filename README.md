# weather-api

[![Build Status](https://travis-ci.org/robertoduessmann/weather-api.svg?branch=master)](https://travis-ci.org/robertoduessmann/weather-api)

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
{"temperature":"17","wind":"4"}
```