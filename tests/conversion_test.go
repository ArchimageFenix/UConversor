// File: conversion_test.go
// Responsibility: Verify known conversion results and early rejection through the application flow.
// Receives: Known input cases.
// Produces: Automated evidence that calculations and rejection rules behave correctly.
// Previous logical stage: app/conversion implementation.
// Next logical stage: scientific metadata and presentation tests.
// Important restrictions: Scientific assertions use structured numeric results and must not depend on CLI formatting.
package tests

import (
	"errors"
	"math"
	"testing"

	"unit-converter/internal/app"
	"unit-converter/internal/catalog"
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func newApp() *app.App {
	return app.New(units.NewRegistry(catalog.All()))
}

func Test24MPH(t *testing.T) {
	result, err := newApp().Convert("24mph")
	if err != nil {
		t.Fatal(err)
	}

	assertConvertedValue(t, result.Values, "km/h", 38.624256, 1e-9)
	assertConvertedValue(t, result.Values, "m/s", 10.72896, 1e-9)
}

func TestUnknownStopsEarly(t *testing.T) {
	_, err := newApp().Convert("24xyz")
	if !errors.Is(err, app.ErrUnavailable) {
		t.Fatalf("expected unavailable, got %v", err)
	}
}

func TestTemperatureAffine(t *testing.T) {
	result, err := newApp().Convert("32°F")
	if err != nil {
		t.Fatal(err)
	}

	assertConvertedValue(t, result.Values, "K", 273.15, 1e-9)
	assertConvertedValue(t, result.Values, "°C", 0, 1e-9)
}

// TestTimeHours verifies that the Time family is correctly recognized
// and that ordinary hour-based conversions use the second as reference.
func TestTimeHours(t *testing.T) {
	result, err := newApp().Convert("48h")
	if err != nil {
		t.Fatal(err)
	}

	if result.Family != string(units.FamilyTime) {
		t.Fatalf(
			"expected family %q, got %q",
			units.FamilyTime,
			result.Family,
		)
	}

	if result.Magnitude != string(units.MagnitudeTime) {
		t.Fatalf(
			"expected magnitude %q, got %q",
			units.MagnitudeTime,
			result.Magnitude,
		)
	}

	assertConvertedValue(t, result.Values, "s", 172800, 1e-9)
	assertConvertedValue(t, result.Values, "min", 2880, 1e-9)
	assertConvertedValue(t, result.Values, "d", 2, 1e-9)
}

// TestGregorianAverageYear protects UConversor's explicit convention
// that "yr" represents the mean Gregorian year:
//
//	1 yr = 365.2425 d = 31,556,952 s
//
// It does not represent a specific Gregorian calendar year,
// which may contain 365 or 366 days.
func TestGregorianAverageYear(t *testing.T) {
	result, err := newApp().Convert("1yr")
	if err != nil {
		t.Fatal(err)
	}

	assertConvertedValue(t, result.Values, "s", 31556952, 1e-9)
	assertConvertedValue(t, result.Values, "d", 365.2425, 1e-9)
	assertConvertedValue(t, result.Values, "h", 8765.82, 1e-9)
}

// TestMilliseconds verifies a sub-second unit and ensures that
// the "ms" symbol is recognized as millisecond.
func TestMilliseconds(t *testing.T) {
	result, err := newApp().Convert("1500ms")
	if err != nil {
		t.Fatal(err)
	}

	assertConvertedValue(t, result.Values, "s", 1.5, 1e-9)
	assertConvertedValue(t, result.Values, "ms", 1500, 1e-9)
}

func assertConvertedValue(
	t *testing.T,
	values []model.ConvertedValue,
	symbol string,
	expected float64,
	tolerance float64,
) {
	t.Helper()

	for _, value := range values {
		if value.Symbol == symbol {
			if math.Abs(value.Value-expected) > tolerance {
				t.Fatalf(
					"%s: expected %.12f, got %.12f",
					symbol,
					expected,
					value.Value,
				)
			}
			return
		}
	}

	t.Fatalf("conversion for %s not found", symbol)
}
