// File: engine.go
// Responsibility: Convert a validated value to every compatible unit of the same magnitude.
// Receives: Numeric value, recognized source unit, and registry.
// Produces: Unrounded ConversionResult values.
// Previous logical stage: conversion/validator.go.
// Next logical stage: output/formatter.go.
// Important restrictions: Does not parse input, reject unknown units, print output, or round intermediate values.
package conversion

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func Convert(value float64, source units.Unit, registry *units.Registry) model.ConversionResult {
	referenceValue := source.ToReference(value)
	compatible := registry.UnitsForMagnitude(source.Magnitude)
	values := make([]model.ConvertedValue, 0, len(compatible))
	for _, target := range compatible {
		values = append(values, model.ConvertedValue{
			Name: target.Name, Symbol: target.Symbol, Value: target.FromReference(referenceValue),
		})
	}
	return model.ConversionResult{
		Family: string(source.Family), Magnitude: string(source.Magnitude),
		Input:  model.ConvertedValue{Name: source.Name, Symbol: source.Symbol, Value: value},
		Values: values,
	}
}
