// File: validator.go
//
// Responsibility:
//   - Validate scientific metadata associated with a recognized unit.
//   - Validate physical value constraints before conversion.
//
// Receives:
//   - A recognized unit.
//   - Numeric input value.
//
// Produces:
//   - nil when the value and unit metadata are valid for conversion.
//   - A typed validation error when the conversion cannot continue.
//
// Previous logical stage:
//   - units/registry.go.
//
// Next logical stage:
//   - conversion/engine.go.
//
// Important restrictions:
//   - Unknown units must already have been rejected before this stage.
//   - Validator must never recognize or guess units.
//   - Validator must never invent missing scientific sources.
//   - Validator must not perform conversions.
//   - Validation errors must remain identifiable with errors.Is()
//     so presentation layers can classify them safely.
package conversion

import (
	"errors"
	"fmt"

	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

// ErrMissingSource indicates that a recognized unit does not have
// the scientific metadata required to guarantee a validated conversion.
//
// This represents an internal catalog/configuration problem and should
// not be exposed directly to the end user.
var ErrMissingSource = errors.New(
	"la unidad no posee una fuente científica registrada",
)

// ErrOutOfRange indicates that the numeric value violates a physical
// constraint defined by the recognized unit.
//
// This is considered a user-input validation error and may be presented
// through a controlled, friendly message in the user interface.
var ErrOutOfRange = errors.New(
	"valor fuera del rango permitido",
)

// Validate verifies that the recognized unit has valid scientific
// metadata and that the numeric value satisfies its physical constraints.
//
// The function assumes that unit recognition has already been completed
// successfully by the registry/application layer.
func Validate(unit units.Unit, value float64) error {
	if unit.Source.Authority == "" ||
		unit.Source.Reference == "" {
		return ErrMissingSource
	}

	if unit.Source.Kind != model.SourceOfficial &&
		unit.Source.Kind != model.SourceExact &&
		unit.Source.Kind != model.SourceDerived &&
		unit.Source.Kind != model.SourceCustom {
		return ErrMissingSource
	}

	if unit.MinValue != nil &&
		value < *unit.MinValue {
		return fmt.Errorf(
			"%w: %s no admite valores menores que %.4f %s",
			ErrOutOfRange,
			unit.Name,
			*unit.MinValue,
			unit.Symbol,
		)
	}

	return nil
}
