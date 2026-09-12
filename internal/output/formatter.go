// File: formatter.go
// Responsibility: Format unrounded conversion results into terminal-ready table data.
// Receives: model.ConversionResult.
// Produces: Text table with exactly four decimal places and no ANSI color codes.
// Previous logical stage: conversion/engine.go.
// Next logical stage: output/console.go.
// Important restrictions: Rounding occurs only here; it must not alter engine values or add terminal-specific colors.
package output

import (
	"fmt"
	"math"
	"strings"
	"unit-converter/internal/model"
)

const (
	unitColumnWidth   = 24
	symbolColumnWidth = 10
	valueColumnWidth  = 16
)

// Format preserves the original plain-text formatter for non-terminal consumers
// and automated tests. The visible precision is always four decimal places.
func Format(result model.ConversionResult) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Familia: %s\nMagnitud: %s\n\n", result.Family, result.Magnitude)
	b.WriteString(FormatTable(result))
	return strings.TrimRight(b.String(), "\n")
}

// FormatTable produces an aligned ASCII table without colors. Keeping ANSI
// sequences outside this function prevents escape codes from breaking column widths.
func FormatTable(result model.ConversionResult) string {
	var b strings.Builder

	writeBorder(&b)
	fmt.Fprintf(
		&b,
		"| %-*s | %-*s | %*s |\n",
		unitColumnWidth, "UNIDAD",
		symbolColumnWidth, "SÍMBOLO",
		valueColumnWidth, "VALOR",
	)
	writeBorder(&b)

	for _, v := range result.Values {
		value := formatValue(v.Value)

		fmt.Fprintf(
			&b,
			"| %-*s | %-*s | %*s |\n",
			unitColumnWidth, v.Name,
			symbolColumnWidth, v.Symbol,
			valueColumnWidth, value,
		)
	}

	writeBorder(&b)
	return strings.TrimRight(b.String(), "\n")
}

func writeBorder(b *strings.Builder) {
	fmt.Fprintf(
		b,
		"+-%s-+-%s-+-%s-+\n",
		strings.Repeat("-", unitColumnWidth),
		strings.Repeat("-", symbolColumnWidth),
		strings.Repeat("-", valueColumnWidth),
	)
}
func formatValue(value float64) string {
	abs := math.Abs(value)

	if abs >= 1e9 || (abs > 0 && abs < 1e-4) {
		return fmt.Sprintf("%.4e", value)
	}

	return fmt.Sprintf("%.4f", value)
}
