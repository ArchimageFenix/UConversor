// File: force.go
// Responsibility: Declare force units and their conversion relations.
// Receives: No runtime input.
// Produces: Force unit definitions for the central unit registry.
// Previous logical stage: units/family.go and units/magnitude.go.
// Next logical stage: catalog/catalog.go and units/registry.go.
// Important restrictions:
//   - Force units must use scientifically documented conversion relations.
//   - The newton is the reference unit for this magnitude.
//   - This file must not contain conversion engine logic.
//   - Adding force units here must not require changes to the conversion engine.

package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func ForceUnits() []units.Unit {
	si := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "SI Brochure, 9th ed., version 4.01",
		URL:       "https://www.bipm.org/en/publications/si-brochure",
		Notes:     "The newton (N) is the coherent SI derived unit of force: kg·m/s².",
	}

	nist := model.Source{
		Kind:      model.SourceExact,
		Authority: "NIST",
		Reference: "NIST Guide to the SI, SP 811 - Footnotes",
		URL:       "https://www.nist.gov/pml/special-publication-811/nist-guide-si-footnotes",
		Notes:     "1 lbf = 4.4482216152605 N using the standard acceleration gn = 9.80665 m/s².",
	}

	return []units.Unit{
		{
			Name:      "newton",
			Symbol:    "N",
			Aliases:   []string{"newton"},
			Family:    units.FamilyForce,
			Magnitude: units.MagnitudeForce,
			Reference: true,
			Kind:      units.TransformLinear,
			Scale:     1,
			Source:    si,
		},
		{
			Name:      "kilonewton",
			Symbol:    "kN",
			Aliases:   []string{"kilonewton"},
			Family:    units.FamilyForce,
			Magnitude: units.MagnitudeForce,
			Kind:      units.TransformLinear,
			Scale:     1000,
			Source:    si,
		},

		{
			Name:      "meganewton",
			Symbol:    "MN",
			Aliases:   []string{"meganewton"},
			Family:    units.FamilyForce,
			Magnitude: units.MagnitudeForce,
			Kind:      units.TransformLinear,
			Scale:     1e6,
			Source:    si,
		},

		{
			Name:      "libra-fuerza",
			Symbol:    "lbf",
			Aliases:   []string{"pound-force"},
			Family:    units.FamilyForce,
			Magnitude: units.MagnitudeForce,
			Kind:      units.TransformLinear,
			Scale:     4.4482216152605,
			Source:    nist,
		},
	}
}
