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
