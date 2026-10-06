```go
package v2

import "testing"

func TestCurrentConditionTemp(t *testing.T) {
	tests := []struct {
		name     string
		unit     string
		expected string
	}{
		{
			name:     "metric",
			unit:     "m",
			expected: "17 °C",
		},
		{
			name:     "USCS",
			unit:     "u",
			expected: "62 °F",
		},
		{
			name:     "unknown unit defaults to metric",
			unit:     "unknown",
			expected: "17 °C",
		},
		{
			name:     "empty unit defaults to metric",
			unit:     "",
			expected: "17 °C",
		},
	}

	cc := currentCondition{
		TempC: "17",
		TempF: "62",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := cc.Temp(tt.unit)

			if got != tt.expected {
				t.Errorf(
					"Temp(%q): expected %q, got %q",
					tt.unit,
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestCurrentConditionWindspeed(t *testing.T) {
	tests := []struct {
		name     string
		unit     string
		expected string
	}{
		{
			name:     "metric",
			unit:     "m",
			expected: "19 km/h",
		},
		{
			name:     "USCS",
			unit:     "u",
			expected: "11 mph",
		},
		{
			name:     "unknown unit defaults to metric",
			unit:     "unknown",
			expected: "19 km/h",
		},
		{
			name:     "empty unit defaults to metric",
			unit:     "",
			expected: "19 km/h",
		},
	}

	cc := currentCondition{
		WindspeedKmph:  "19",
		WindspeedMiles: "11",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := cc.Windspeed(tt.unit)

			if got != tt.expected {
				t.Errorf(
					"Windspeed(%q): expected %q, got %q",
					tt.unit,
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestHourlyTemp(t *testing.T) {
	tests := []struct {
		name     string
		unit     string
		expected string
	}{
		{
			name:     "metric",
			unit:     "m",
			expected: "30 °C",
		},
		{
			name:     "USCS",
			unit:     "u",
			expected: "86 °F",
		},
		{
			name:     "unknown unit defaults to metric",
			unit:     "unknown",
			expected: "30 °C",
		},
		{
			name:     "empty unit defaults to metric",
			unit:     "",
			expected: "30 °C",
		},
	}

	h := hourly{
		TempC: "30",
		TempF: "86",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := h.Temp(tt.unit)

			if got != tt.expected {
				t.Errorf(
					"Temp(%q): expected %q, got %q",
					tt.unit,
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestHourlyWindspeed(t *testing.T) {
	tests := []struct {
		name     string
		unit     string
		expected string
	}{
		{
			name:     "metric",
			unit:     "m",
			expected: "25 km/h",
		},
		{
			name:     "USCS",
			unit:     "u",
			expected: "15 mph",
		},
		{
			name:     "unknown unit defaults to metric",
			unit:     "unknown",
			expected: "25 km/h",
		},
		{
			name:     "empty unit defaults to metric",
			unit:     "",
			expected: "25 km/h",
		},
	}

	h := hourly{
		WindspeedKmph:  "25",
		WindspeedMiles: "15",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := h.Windspeed(tt.unit)

			if got != tt.expected {
				t.Errorf(
					"Windspeed(%q): expected %q, got %q",
					tt.unit,
					tt.expected,
					got,
				)
			}
		})
	}
}
```
