// File: app.go
// Responsibility: Orchestrate the complete CLI-independent conversion workflow.
// Receives: Raw user expression plus constructed registry.
// Produces: Structured conversion result or a user-facing domain error.
// Previous logical stage: cmd/converter/main.go (CLI boundary).
// Next logical stage: output presenter / future web adapter.
// Important restrictions: Contains no scientific constants; unknown units are rejected before validation/conversion.
package app

import (
	"errors"
	"fmt"
	"unit-converter/internal/conversion"
	"unit-converter/internal/input"
	"unit-converter/internal/model"
	"unit-converter/internal/output"
	"unit-converter/internal/units"
)

var ErrUnavailable = errors.New("unidad o familia no disponible")

type App struct{ registry *units.Registry }

func New(registry *units.Registry) *App { return &App{registry: registry} }

// SupportedUnits returns the units currently available to
// presentation layers.
//
// The application exposes this information so interfaces such
// as the web layer do not need direct access to Registry.
func (a *App) SupportedUnits() []units.Unit {
	return a.registry.AllUnits()
}

// Convert executes the domain workflow and returns structured data. This is the
// preferred entry point for presentation layers such as the current CLI or a future web server.
func (a *App) Convert(raw string) (model.ConversionResult, error) {
	normalized := input.Normalize(raw)
	parsed, err := input.Parse(normalized)
	if err != nil {
		return model.ConversionResult{}, err
	}

	unit, ok := a.registry.Find(parsed.Unit)
	if !ok {
		return model.ConversionResult{}, fmt.Errorf("%w: %s", ErrUnavailable, parsed.Unit)
	}

	req := model.ConversionRequest{
		Value: parsed.Value, InputUnit: parsed.Unit, UnitSymbol: unit.Symbol,
		Magnitude: string(unit.Magnitude), Family: string(unit.Family),
	}
	if err := conversion.Validate(unit, req.Value); err != nil {
		return model.ConversionResult{}, err
	}

	return conversion.Convert(req.Value, unit, a.registry), nil
}

// Process remains as a plain-text compatibility path for tests and future
// non-terminal consumers. Terminal styling must not be introduced here.
func (a *App) Process(raw string) (string, error) {
	result, err := a.Convert(raw)
	if err != nil {
		return "", err
	}
	return output.Format(result), nil
}
