// File: unit.go
// Responsibility: Define one unit and its transformation to/from a magnitude reference unit.
// Receives: Catalog declarations.
// Produces: Unit definitions used by registry, validator, and conversion engine.
// Previous logical stage: catalog files.
// Next logical stage: units/registry.go and conversion/engine.go.
// Important restrictions: Transformations must be scientifically traceable and must not perform formatting.
package units

import "unit-converter/internal/model"

type TransformKind string

const (
	TransformLinear TransformKind = "linear"
	TransformAffine TransformKind = "affine"
)

type Unit struct {
	Name      string
	Symbol    string
	Aliases   []string
	Family    Family
	Magnitude Magnitude
	Reference bool
	Kind      TransformKind
	Scale     float64
	Offset    float64
	MinValue  *float64
	Source    model.Source
}

// ToReference converts a value into the reference unit of the same magnitude.
// Formula: reference = value*Scale + Offset.
func (u Unit) ToReference(value float64) float64 {
	return value*u.Scale + u.Offset
}

// FromReference converts from the reference unit back into this unit.
func (u Unit) FromReference(value float64) float64 {
	return (value - u.Offset) / u.Scale
}
