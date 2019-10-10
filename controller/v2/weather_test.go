package v2

import "testing"

func TestCurrentConditionTemp(t *testing.T) {
	tt := []struct {
		description string
		unit        string
		expected    string
	}{
		// used by default everywhere except US
		{description: "metric (SI)", unit: "m", expected: "17 °C"},
		// used by default in US
		{description: "USCS", unit: "u", expected: "62 °F"},
		{description: "unknown metric", unit: "unknown", expected: "17 °C"},
	}

	var cc = currentCondition{
		TempC: "17",
		TempF: "62",
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if actual := cc.Temp(tc.unit); actual != tc.expected {
				t.Errorf("expected %s; got %s", tc.expected, actual)
			}
		})
	}
}

func TestCurrentConditionWindspeed(t *testing.T) {
	tt := []struct {
		description string
		unit        string
		expected    string
	}{
		// used by default everywhere except US
		{description: "metric (SI)", unit: "m", expected: "19 km/h"},
		// used by default in US
		{description: "USCS", unit: "u", expected: "11 mph"},
		{description: "unknown metric", unit: "unknown", expected: "19 km/h"},
	}

	var cc = currentCondition{
		WindspeedKmph:  "19",
		WindspeedMiles: "11",
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if actual := cc.Windspeed(tc.unit); actual != tc.expected {
				t.Errorf("expected %s; got %s", tc.expected, actual)
			}
		})
	}
}

func TestHourlyTemp(t *testing.T) {
	tt := []struct {
		description string
		unit        string
		expected    string
	}{
		// used by default everywhere except US
		{description: "metric (SI)", unit: "m", expected: "30 °C"},
		// used by default in US
		{description: "USCS", unit: "u", expected: "86 °F"},
		{description: "unknown metric", unit: "unknown", expected: "30 °C"},
	}

	var h = hourly{
		TempC: "30",
		TempF: "86",
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if actual := h.Temp(tc.unit); actual != tc.expected {
				t.Errorf("expected %s; got %s", tc.expected, actual)
			}
		})
	}
}

func TestHourlyWindspeed(t *testing.T) {
	tt := []struct {
		description string
		unit        string
		expected    string
	}{
		// used by default everywhere except US
		{description: "metric (SI)", unit: "m", expected: "25 km/h"},
		// used by default in US
		{description: "USCS", unit: "u", expected: "15 mph"},
		{description: "unknown metric", unit: "unknown", expected: "25 km/h"},
	}

	var h = hourly{
		WindspeedKmph:  "25",
		WindspeedMiles: "15",
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if actual := h.Windspeed(tc.unit); actual != tc.expected {
				t.Errorf("expected %s; got %s", tc.expected, actual)
			}
		})
	}
}
