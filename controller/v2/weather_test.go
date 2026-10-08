package v2

import "testing"

// ============================================================
// Current condition
// ============================================================

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

	runUnitTests(t, cc.Temp, tests, "Temp")
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

	runUnitTests(t, cc.Windspeed, tests, "Windspeed")
}

// ============================================================
// Hourly
// ============================================================

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

	runUnitTests(t, h.Temp, tests, "Temp")
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

	runUnitTests(t, h.Windspeed, tests, "Windspeed")
}

// ============================================================
// Helpers
// ============================================================

type unitFunc func(string) string

type unitTest struct {
	name     string
	unit     string
	expected string
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
O que foi corrigido
Removida a duplicação entre tempFunc e windspeedFunc.
Removidos os dois helpers praticamente idênticos.
Criado um único unitTest.
Criado um único runUnitTests.
Mantido o t.Parallel().
Mantido o comportamento original dos testes.

