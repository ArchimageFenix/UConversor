// File: mass.go
// Responsibility: Declare mass units with kilogram as reference.
// Receives: No runtime input.
// Produces: Mass unit definitions.
// Previous logical stage: normative source selection.
// Next logical stage: units/registry.go.
// Important restrictions: This catalog models mass, not force/physical weight.
package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func MassUnits() []units.Unit {
	bipm := model.Source{Kind: model.SourceOfficial, Authority: "BIPM", Reference: "SI Brochure, 9th ed., version 4.01", URL: "https://www.bipm.org/en/publications/si-brochure"}
	nist := model.Source{Kind: model.SourceExact, Authority: "NIST", Reference: "International yard and pound / SI conversion values", URL: "https://www.nist.gov/pml/special-publication-811"}
	zero := 0.0
	return []units.Unit{
		{Name: "microgramo", Symbol: "µg", Aliases: []string{"ug"}, Family: units.FamilyMass, Magnitude: units.MagnitudeMass, Scale: 1e-9, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "miligramo", Symbol: "mg", Family: units.FamilyMass, Magnitude: units.MagnitudeMass, Scale: 1e-6, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "gramo", Symbol: "g", Family: units.FamilyMass, Magnitude: units.MagnitudeMass, Scale: 0.001, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "kilogramo", Symbol: "kg", Family: units.FamilyMass, Magnitude: units.MagnitudeMass, Reference: true, Scale: 1, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "onza avoirdupois", Symbol: "oz", Family: units.FamilyMass, Magnitude: units.MagnitudeMass, Scale: 0.028349523125, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
		{Name: "libra avoirdupois", Symbol: "lb", Family: units.FamilyMass, Magnitude: units.MagnitudeMass, Scale: 0.45359237, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
	}
}
