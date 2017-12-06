package model

// Weather entity
type Weather struct {
	Temperature string      `json:"temperature"`
	Wind        string      `json:"wind"`
	Description string      `json:"description"`
	Forecast    [2]Forecast `json:"forecast"`
}

// Forecast entity
type Forecast struct {
	Day         int    `json:"day"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}
