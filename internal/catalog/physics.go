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

	//ZONA DE FUENTES CIENTIFICAS PARA LOS CALCULOS, AQUI SE DEBEN AGREGAR MAS FUENTES CUANDO SEA
	bipm := model.Source{Kind: model.SourceOfficial, Authority: "BIPM", Reference: "SI Brochure, 9th ed., version 4.01", URL: "https://www.bipm.org/en/publications/si-brochure"}
	nist := model.Source{Kind: model.SourceExact, Authority: "NIST", Reference: "NIST Guide to the SI, Appendix B", URL: "https://www.nist.gov/pml/special-publication-811"}

	fujiPressure := model.Source{
		Kind:      model.SourceCustom,
		Authority: "Fuji Electric France",
		Reference: "Unidades de presión y conversión: todo lo que necesita saber",
		URL:       "https://www.fujielectric.fr/es/blog/unidades-de-presion-y-su-conversion-todo-lo-que-necesita-saber/",
		Notes:     "Referencia técnica e industrial secundaria para unidades de presión y equivalencias en pascales.",
	}
	//ZONA DE FUENTES CIENTIFICAS PARA LOS CALCULOS FINAL
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
		{Name: "atmósfera estándar", Symbol: "atm", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 101325, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
		{Name: "hectopascal", Symbol: "hPa", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 100, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "megapascal", Symbol: "MPa", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 1_000_000, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "atmósfera técnica", Symbol: "at", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 98066.5, Kind: units.TransformLinear, MinValue: &zero, Source: fujiPressure},
		{Name: "gigapascal", Symbol: "GPa", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 1_000_000_000, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "milibar", Symbol: "mbar", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 100, Kind: units.TransformLinear, MinValue: &zero, Source: fujiPressure},
		{Name: "torr", Symbol: "Torr", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 101325.0 / 760.0, Kind: units.TransformLinear, MinValue: &zero, Source: fujiPressure},
		{Name: "milímetro de mercurio", Symbol: "mmHg", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 133.3224, Kind: units.TransformLinear, MinValue: &zero, Source: fujiPressure},
		{Name: "pulgada de mercurio", Symbol: "inHg", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 3386.389, Kind: units.TransformLinear, MinValue: &zero, Source: fujiPressure},
		{Name: "kilogramo-fuerza por centímetro cuadrado", Symbol: "kgf/cm²", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 98066.5, Kind: units.TransformLinear, MinValue: &zero, Source: fujiPressure},
		{Name: "libra-fuerza por pulgada cuadrada", Symbol: "psi", Family: units.FamilyPhysics, Magnitude: units.MagnitudePressure, Scale: 6894.757293168361, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
		{Name: "joule", Symbol: "J", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Reference: true, Scale: 1, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "kilojoule", Symbol: "kJ", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 1000, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "caloría termoquímica", Symbol: "cal", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 4.184, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
		{Name: "kilocaloría termoquímica", Symbol: "kcal", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 4184, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
		{Name: "watt-hora", Symbol: "Wh", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 3600, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "kilowatt-hora", Symbol: "kWh", Family: units.FamilyPhysics, Magnitude: units.MagnitudeEnergy, Scale: 3.6e6, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
	}
}
