
package main

import "testing"

type unitFunc func(string) string

type unitTest struct {
	name     string
	unit     string
	expected string
}

func TestCurrentConditionTemp(t *testing.T) {
	t.Parallel()

	cc := currentCondition{
		TempC: "17",
		TempF: "62",
	}

	tests := []unitTest{
		{name: "metric (SI)", unit: "m", expected: "17 °C"},
		{name: "USCS", unit: "u", expected: "62 °F"},
		{name: "unknown unit", unit: "unknown", expected: "17 °C"},
		{name: "empty unit", unit: "", expected: "17 °C"},
		{name: "whitespace unit", unit: " ", expected: "17 °C"},
	}

	runUnitTests(t, cc.Temp, tests, "CurrentCondition.Temp")
}

func TestCurrentConditionWindspeed(t *testing.T) {
	t.Parallel()

	cc := currentCondition{
		WindspeedKmph:  "19",
		WindspeedMiles: "11",
	}

	tests := []unitTest{
		{name: "metric (SI)", unit: "m", expected: "19 km/h"},
		{name: "USCS", unit: "u", expected: "11 mph"},
		{name: "unknown unit", unit: "unknown", expected: "19 km/h"},
		{name: "empty unit", unit: "", expected: "19 km/h"},
		{name: "whitespace unit", unit: " ", expected: "19 km/h"},
	}

	runUnitTests(t, cc.Windspeed, tests, "CurrentCondition.Windspeed")
}

func TestHourlyTemp(t *testing.T) {
	t.Parallel()

	h := hourly{
		TempC: "30",
		TempF: "86",
	}

	tests := []unitTest{
		{name: "metric (SI)", unit: "m", expected: "30 °C"},
		{name: "USCS", unit: "u", expected: "86 °F"},
		{name: "unknown unit", unit: "unknown", expected: "30 °C"},
		{name: "empty unit", unit: "", expected: "30 °C"},
		{name: "whitespace unit", unit: " ", expected: "30 °C"},
	}

	runUnitTests(t, h.Temp, tests, "Hourly.Temp")
}

func TestHourlyWindspeed(t *testing.T) {
	t.Parallel()

	h := hourly{
		WindspeedKmph:  "25",
		WindspeedMiles: "15",
	}

	tests := []unitTest{
		{name: "metric (SI)", unit: "m", expected: "25 km/h"},
		{name: "USCS", unit: "u", expected: "15 mph"},
		{name: "unknown unit", unit: "unknown", expected: "25 km/h"},
		{name: "empty unit", unit: "", expected: "25 km/h"},
		{name: "whitespace unit", unit: " ", expected: "25 km/h"},
	}

	runUnitTests(t, h.Windspeed, tests, "Hourly.Windspeed")
}

func runUnitTests(
	t *testing.T,
	fn unitFunc,
	tests []unitTest,
	functionName string,
) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := fn(tt.unit)
			if got != tt.expected {
				t.Errorf(
					"%s(%q) = %q; want %q",
					functionName,
					tt.unit,
					got,
					tt.expected,
				)
			}
		})
	}
}