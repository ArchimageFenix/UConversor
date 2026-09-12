// File: parser.go
// Responsibility: Split a normalized expression into numeric value and unit token.
// Receives: Normalized text such as "24mph" or "24 mph".
// Produces: Parsed numeric value and raw unit token.
// Previous logical stage: input/normalizer.go.
// Next logical stage: units/registry.go.
// Important restrictions: Does not recognize units, select formulas, or perform conversions.
package input

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type Parsed struct {
	Value float64
	Unit  string
}

var (
	ErrEmptyInput   = errors.New("entrada vacía")
	ErrMissingValue = errors.New("falta el valor numérico")
	ErrMissingUnit  = errors.New("falta la unidad")
	ErrInvalidValue = errors.New("valor numérico inválido")
)

func Parse(normalized string) (Parsed, error) {
	s := strings.TrimSpace(normalized)
	if s == "" {
		return Parsed{}, ErrEmptyInput
	}

	// Find the longest prefix that can represent a float: sign, digits, decimal point, exponent.
	cut := 0
	seenDigit := false
	for i, r := range s {
		allowed := unicode.IsDigit(r) || r == '+' || r == '-' || r == '.' || r == 'e' || r == 'E'
		if !allowed {
			cut = i
			break
		}
		if unicode.IsDigit(r) {
			seenDigit = true
		}
		cut = i + len(string(r))
	}
	if !seenDigit || cut == 0 {
		return Parsed{}, ErrMissingValue
	}

	numPart := strings.TrimSpace(s[:cut])
	unitPart := strings.TrimSpace(s[cut:])
	if unitPart == "" {
		return Parsed{}, ErrMissingUnit
	}

	value, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		return Parsed{}, ErrInvalidValue
	}
	return Parsed{Value: value, Unit: unitPart}, nil
}
