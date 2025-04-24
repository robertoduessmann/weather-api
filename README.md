# weather-api

[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/robertoduessmann/weather-api)](https://goreportcard.com/report/github.com/robertoduessmann/weather-api)
[![GoDoc](https://godoc.org/github.com/robertoduessmann/weather-api?status.svg)](https://godoc.org/github.com/robertoduessmann/weather-api)

> A REST API to check the current weather.

> http://goweather.xyz/weather/Berlin<br /> > http://goweather.xyz/weather/{city}

## Build locally (Mac users)

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
  "temperature": "29 °C",
  "wind": "20 km/h",
  "description": "Partly cloudy",
  "forecast": [
    {
      "day": 1,
      "temperature": "27 °C",
      "wind": "12 km/h"
    },
    {
      "day": 2,
      "temperature": "22 °C",
      "wind": "8 km/h"
    }
  ]
}
```

## Web Version

Few web clients of the API and their Projects:

1. **Client:** [https://reacttempo.netlify.app/](https://reacttempo.netlify.app/)  
   **Project:** [https://github.com/GabrielCampos99/appTempo](https://github.com/GabrielCampos99/appTempo)

2. **Client:** [https://emaniaditya.github.io/weather-app](https://emaniaditya.github.io/weather-app)  
   **Project:** [https://github.com/emaniaditya/weather-app/](https://github.com/emaniaditya/weather-app/)

3. **Client:** [https://weather-react-tsx.netlify.app/](https://weather-react-tsx.netlify.app/)  
   **Project:** [https://github.com/keissiant/weather-api-react.tsx](https://github.com/keissiant/weather-api-react.tsx)

## License

The MIT License ([MIT](https://github.com/robertoduessmann/weather-api/blob/master/LICENSE))
