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
package web

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// formatWebValue converts a raw numeric value into
// a presentation-friendly string for the web interface.
//
// Ordinary values use at most two decimal places.
// Unnecessary trailing zeros are removed.
//
// Very large or very small values preserve the existing
// scientific-notation policy with four decimal places
// in the mantissa.
func formatWebValue(value float64) string {
	abs := math.Abs(value)

	if abs >= 1e9 || (abs > 0 && abs < 1e-4) {
		return fmt.Sprintf("%.4e", value)
	}

	formatted := strconv.FormatFloat(value, 'f', 3, 64)

	formatted = strings.TrimRight(formatted, "0")
	formatted = strings.TrimRight(formatted, ".")

	return formatted
}
