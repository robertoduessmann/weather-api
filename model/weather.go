package model

type Weather struct {
	Temperature string      `json:"temperature"`
	Wind        string      `json:"wind"`
	Description string      `json:"description"`
	Forecast    [2]Forecast `json:"forecast"`
}

type Forecast struct {
	Day         int    `json:"day"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}
