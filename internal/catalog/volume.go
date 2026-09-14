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

	nistLiquid := model.Source{
		Kind:      model.SourceExact,
		Authority: "NIST",
		Reference: "NIST Handbook 133, Appendix E - General Tables of Units of Measurement",
		URL:       "https://www.nist.gov/system/files/documents/2023/02/10/2023%20NIST%20Handbook%20133.pdf",
		Notes:     "Exact U.S. liquid-volume relationships: 1 gal = 231 in³ = 4 qt = 8 pt = 128 fl oz.",
	}

	ukGov := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "UK Government",
		Reference: "Weights and Measures Act 1985, Schedule 1, Part IV",
		URL:       "https://www.legislation.gov.uk/ukpga/1985/72/schedule/1",
		Notes:     "Official imperial capacity definition: 1 imperial gallon = 4.54609 cubic decimetres.",
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

		{
			Name:      "galón estadounidense",
			Symbol:    "US gal",
			Aliases:   []string{"usgal", "galón US", "galon US"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.003785411784,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    nistLiquid,
		},
		{
			Name:      "cuarto líquido estadounidense",
			Symbol:    "US qt",
			Aliases:   []string{"usqt", "quart US"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.000946352946,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    nistLiquid,
		},
		{
			Name:      "pinta líquida estadounidense",
			Symbol:    "US pt",
			Aliases:   []string{"uspt", "pint US"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.000473176473,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    nistLiquid,
		},
		{
			Name:      "onza líquida estadounidense",
			Symbol:    "US fl oz",
			Aliases:   []string{"usfloz", "fl oz US"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.0000295735295625,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    nistLiquid,
		},
		{
			Name:      "galón imperial",
			Symbol:    "imp gal",
			Aliases:   []string{"impgal", "UK gal", "ukgal"},
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeVolume,
			Scale:     0.00454609,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    ukGov,
		},
	}
}
