// File: viewmodel.go
//
// Responsibility:
//   - Define the data exposed by the web layer to HTML templates.
//   - Adapt conversion results to presentation-friendly structures.
//
// Receives:
//   - Conversion results produced by the application layer.
//   - User input and web presentation state.
//
// Produces:
//   - Structured data consumed by HTML templates.
//
// Previous logical stage:
//   - handlers.go.
//
// Next logical stage:
//   - HTML templates.
//
// Important restrictions:
//   - Must not perform unit conversions.
//   - Must not contain scientific formulas.
//   - Must not duplicate catalog information.
//   - Must remain a presentation model, not a domain model.
package web

import "unit-converter/internal/model"

// PageViewModel represents the complete state
// required to render the main UConversor web page.
type PageViewModel struct {
	Input      string
	ActivePage string
	HasResult  bool
	HasError   bool

	Result *ResultViewModel
	Error  *ErrorViewModel
}

// ResultViewModel represents one successful conversion
// prepared for HTML presentation.
type ResultViewModel struct {
	Family       string
	Magnitude    string
	VisualFamily string
	InputName    string
	InputSymbol  string
	InputValue   string

	Values []ConvertedValueViewModel
}

// ConvertedValueViewModel represents one result card.
type ConvertedValueViewModel struct {
	Name     string
	Symbol   string
	Value    string
	IsOrigin bool
}

// ErrorViewModel represents an error shown by
// the web presentation layer.
type ErrorViewModel struct {
	Title   string
	Message string
}

// NewResultViewModel adapts a domain conversion result
// into data prepared for HTML presentation.
//
// It does not perform calculations. Numeric formatting is
// delegated to formatter.go.
func NewResultViewModel(
	result model.ConversionResult,
) ResultViewModel {
	values := make(
		[]ConvertedValueViewModel,
		0,
		len(result.Values),
	)

	for _, value := range result.Values {
		values = append(
			values,
			ConvertedValueViewModel{
				Name:     value.Name,
				Symbol:   value.Symbol,
				Value:    formatWebValue(value.Value),
				IsOrigin: value.Symbol == result.Input.Symbol,
			},
		)
	}

	return ResultViewModel{
		Family:       result.Family,
		Magnitude:    result.Magnitude,
		VisualFamily: visualFamily(result.Family, result.Magnitude),
		InputName:    result.Input.Name,
		InputSymbol:  result.Input.Symbol,
		InputValue:   formatWebValue(result.Input.Value),
		Values:       values,
	}
}

// visualFamily maps an already classified scientific result
// to a stable visual identifier used exclusively by the Web presentation.
//
// Family and magnitude are received from the domain result.
// This function does not determine or modify their scientific classification.
func visualFamily(family, magnitude string) string {
	switch {
	case family == "Longitud / distancia":
		return "length"

	case family == "Física" && magnitude == "Temperatura":
		return "temperature"

	case family == "Física" && magnitude == "Presión":
		return "pressure"

	case family == "Física" && magnitude == "Energía":
		return "energy"

	case family == "Electrónica" && magnitude == "Voltaje":
		return "voltage"

	case family == "Electrónica" && magnitude == "Corriente":
		return "current"

	default:
		return "generic"
	}
}
