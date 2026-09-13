// File: volume.go
//
// Responsibility: Declare volume conversion units.
//
// Receives: No runtime input.
//
// Produces: Volume unit definitions belonging to the Physics family.
//
// Previous logical stage: normative/scientific source selection.
//
// Next logical stage: units/registry.go through catalog.All().
//
// Important restrictions:
//   - Volume units are convertible only within MagnitudeVolume.
//   - The reference unit is the cubic metre (m³).
//   - All current volume transformations are linear.
//   - Physical volume cannot be negative.

package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func VolumeUnits() []units.Unit {

	bipm := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "SI Brochure, 9th ed., version 4.01",
		URL:       "https://www.bipm.org/en/publications/si-brochure",
		Notes:     "Reference for SI volume units and the litre as a non-SI unit accepted for use with the SI.",
	}

	nist := model.Source{
		Kind:      model.SourceExact,
		Authority: "NIST",
		Reference: "NIST Guide to the SI, Appendix B",
		URL:       "https://www.nist.gov/pml/special-publication-811",
		Notes:     "Reference for exact relationships derived from the international inch and foot.",
	}

	zero := 0.0

	return []units.Unit{
		{
			Name:      "metro cúbico",
			Symbol:    "m³",
			Aliases:   []string{"m3"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Reference: true,
			Scale:     1,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "litro",
			Symbol:    "L",
			Aliases:   []string{"l", "litro", "litros"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.001,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "decilitro",
			Symbol:    "dL",
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.0001,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "centilitro",
			Symbol:    "cL",
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.00001,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "mililitro",
			Symbol:    "mL",
			Aliases:   []string{"ml"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.000001,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "decímetro cúbico",
			Symbol:    "dm³",
			Aliases:   []string{"dm3"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.001,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "centímetro cúbico",
			Symbol:    "cm³",
			Aliases:   []string{"cm3", "cc"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.000001,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "pie cúbico",
			Symbol:    "ft³",
			Aliases:   []string{"ft3"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.028316846592,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    nist,
		},
		{
			Name:      "pulgada cúbica",
			Symbol:    "in³",
			Aliases:   []string{"in3"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.000016387064,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    nist,
		},
	}
}
