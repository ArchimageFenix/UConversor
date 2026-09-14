// File: formatter.go
//
// Responsibility:
//   - Format numeric conversion values for web presentation.
//
// Receives:
//   - Raw float64 values produced by the conversion engine.
//
// Produces:
//   - Human-readable numeric strings for HTML templates.
//   - Scientific notation in mathematical form: a × 10ⁿ.
//
// Previous logical stage:
//   - model.ConversionResult.
//
// Next logical stage:
//   - viewmodel.go.
//
// Important restrictions:
//   - Must not perform unit conversions.
//   - Must not modify scientific values.
//   - Must not contain unit-specific logic.
//   - Formatting affects presentation only.
//   - Scientific notation must remain plain text and must not inject HTML.
package web

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// formatWebValue converts a raw numeric value into a
// presentation-friendly string for the web interface.
//
// Ordinary values use at most three decimal places.
// Unnecessary trailing zeros are removed.
//
// Very large or very small values preserve the existing
// scientific-notation thresholds, but are displayed using
// mathematical notation such as:
//
//	3.0857 × 10²⁵
func formatWebValue(value float64) string {
	abs := math.Abs(value)

	if abs >= 1e9 || (abs > 0 && abs < 1e-4) {
		return formatScientificWeb(value)
	}

	formatted := strconv.FormatFloat(value, 'f', 3, 64)

	formatted = strings.TrimRight(formatted, "0")
	formatted = strings.TrimRight(formatted, ".")

	return formatted
}

// formatScientificWeb converts Go's scientific representation:
//
//	3.0857e+25
//
// into a presentation-oriented mathematical representation:
//
//	3.0857 × 10²⁵
//
// This transformation changes only the displayed text.
func formatScientificWeb(value float64) string {
	raw := fmt.Sprintf("%.4e", value)

	parts := strings.SplitN(raw, "e", 2)
	if len(parts) != 2 {
		return raw
	}

	mantissa := parts[0]
	exponent := normalizeExponent(parts[1])

	return mantissa + " × 10" + toSuperscript(exponent)
}

// normalizeExponent removes syntax that is unnecessary
// in mathematical presentation.
//
// Examples:
//
//	+25  -> 25
//	+09  -> 9
//	-09  -> -9
func normalizeExponent(exponent string) string {
	sign := ""

	if strings.HasPrefix(exponent, "+") {
		exponent = strings.TrimPrefix(exponent, "+")
	} else if strings.HasPrefix(exponent, "-") {
		sign = "-"
		exponent = strings.TrimPrefix(exponent, "-")
	}

	exponent = strings.TrimLeft(exponent, "0")

	if exponent == "" {
		exponent = "0"
	}

	return sign + exponent
}

// toSuperscript transforms exponent characters into their
// Unicode superscript equivalents.
//
// Example:
//
//	-12 -> ⁻¹²
func toSuperscript(value string) string {
	replacer := strings.NewReplacer(
		"0", "⁰",
		"1", "¹",
		"2", "²",
		"3", "³",
		"4", "⁴",
		"5", "⁵",
		"6", "⁶",
		"7", "⁷",
		"8", "⁸",
		"9", "⁹",
		"-", "⁻",
	)

	return replacer.Replace(value)
}
