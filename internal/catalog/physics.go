// File: physics.go
// Responsibility: Declare temperature, pressure, and energy conversion units.
// Receives: No runtime input.
// Produces: Physics-family unit definitions with SI reference units.
// Previous logical stage: normative/scientific source selection.
// Next logical stage: units/registry.go.
// Important restrictions: Temperature uses affine transforms; absolute temperature cannot be below 0 K.
package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func PhysicsUnits() []units.Unit {
	bipm := model.Source{Kind: model.SourceOfficial, Authority: "BIPM", Reference: "SI Brochure, 9th ed., version 4.01", URL: "https://www.bipm.org/en/publications/si-brochure"}
	nist := model.Source{Kind: model.SourceExact, Authority: "NIST", Reference: "NIST Guide to the SI, Appendix B", URL: "https://www.nist.gov/pml/special-publication-811"}
	zero := 0.0
	minus273 := -273.15
	minus459 := -459.67
	return []units.Unit{
		{Name: "kelvin", Symbol: "K", Family: units.FamilyPhysics, Magnitude: units.MagnitudeTemperature, Reference: true, Scale: 1, Offset: 0, Kind: units.TransformAffine, MinValue: &zero, Source: bipm},
		{Name: "grado Celsius", Symbol: "°C", Aliases: []string{"C", "celsius"}, Family: units.FamilyPhysics, Magnitude: units.MagnitudeTemperature, Scale: 1, Offset: 273.15, Kind: units.TransformAffine, MinValue: &minus273, Source: bipm},
		{Name: "grado Fahrenheit", Symbol: "°F", Aliases: []string{"F", "fahrenheit"}, Family: units.FamilyPhysics, Magnitude: units.MagnitudeTemperature, Scale: 5.0 / 9.0, Offset: 255.3722222222222, Kind: units.TransformAffine, MinValue: &minus459, Source: nist},

		{Name: "pascal", Symbol: "Pa", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Reference: true, Scale: 1, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "kilopascal", Symbol: "kPa", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 1000, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "bar", Symbol: "bar", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 100000, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "libra-fuerza por pulgada cuadrada", Symbol: "psi", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 6894.757293168361, Kind: units.TransformLinear, MinValue: &zero, Source: nist},

		{Name: "joule", Symbol: "J", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Reference: true, Scale: 1, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "kilojoule", Symbol: "kJ", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 1000, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "caloría termoquímica", Symbol: "cal", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 4.184, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
		{Name: "kilocaloría termoquímica", Symbol: "kcal", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 4184, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
		{Name: "watt-hora", Symbol: "Wh", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 3600, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "kilowatt-hora", Symbol: "kWh", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 3.6e6, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
	}
}
