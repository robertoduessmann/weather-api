package v2

import "testing"

func TestCurrentConditionTemp(t *testing.T) {
	t.Parallel()

	cc := currentCondition{
		TempC: "17",
		TempF: "62",
	}

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
		{
			name:     "whitespace unit defaults to metric",
			unit:     " ",
			expected: "17 °C",
		},
	}

	runTempTests(t, cc.Temp, tests)
}

func TestCurrentConditionWindspeed(t *testing.T) {
	t.Parallel()

	cc := currentCondition{
		WindspeedKmph:  "19",
		WindspeedMiles: "11",
	}

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
		{
			name:     "whitespace unit defaults to metric",
			unit:     " ",
			expected: "19 km/h",
		},
	}

	runWindspeedTests(t, cc.Windspeed, tests)
}

func TestHourlyTemp(t *testing.T) {
	t.Parallel()

	h := hourly{
		TempC: "30",
		TempF: "86",
	}

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
		{
			name:     "whitespace unit defaults to metric",
			unit:     " ",
			expected: "30 °C",
		},
	}

	runTempTests(t, h.Temp, tests)
}

func TestHourlyWindspeed(t *testing.T) {
	t.Parallel()

	h := hourly{
		WindspeedKmph:  "25",
		WindspeedMiles: "15",
	}

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
		{
			name:     "whitespace unit defaults to metric",
			unit:     " ",
			expected: "25 km/h",
		},
	}

	runWindspeedTests(t, h.Windspeed, tests)
}

// Helpers

type tempFunc func(string) string

type windspeedFunc func(string) string

func runTempTests(
	t *testing.T,
	fn tempFunc,
	tests []struct {
		name     string
		unit     string
		expected string
	},
) {
	t.Helper()

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := fn(tt.unit)

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

func runWindspeedTests(
	t *testing.T,
	fn windspeedFunc,
	tests []struct {
		name     string
		unit     string
		expected string
	},
) {
	t.Helper()

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := fn(tt.unit)

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