// File: electronics.go
// Responsibility: Declare common electrical magnitudes and SI-prefixed units.
// Receives: No runtime input.
// Produces: Voltage, current, resistance, and power definitions.
// Previous logical stage: SI unit/source selection.
// Next logical stage: units/registry.go.
// Important restrictions: Magnitudes in this family remain mutually non-convertible.
package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func ElectronicsUnits() []units.Unit {
	si := model.Source{Kind: model.SourceOfficial, Authority: "BIPM", Reference: "SI Brochure, 9th ed., version 4.01", URL: "https://www.bipm.org/en/publications/si-brochure"}
	zero := 0.0
	return []units.Unit{
		{Name: "milivoltio", Symbol: "mV", Family: units.FamilyElectronics, Magnitude: units.MagnitudeVoltage, Scale: 1e-3, Kind: units.TransformLinear, Source: si},
		{Name: "voltio", Symbol: "V", Family: units.FamilyElectronics, Magnitude: units.MagnitudeVoltage, Reference: true, Scale: 1, Kind: units.TransformLinear, Source: si},
		{Name: "kilovoltio", Symbol: "kV", Family: units.FamilyElectronics, Magnitude: units.MagnitudeVoltage, Scale: 1e3, Kind: units.TransformLinear, Source: si},
		{Name: "microamperio", Symbol: "µA", Aliases: []string{"uA"}, Family: units.FamilyElectronics, Magnitude: units.MagnitudeCurrent, Scale: 1e-6, Kind: units.TransformLinear, Source: si},
		{Name: "miliamperio", Symbol: "mA", Family: units.FamilyElectronics, Magnitude: units.MagnitudeCurrent, Scale: 1e-3, Kind: units.TransformLinear, Source: si},
		{Name: "amperio", Symbol: "A", Family: units.FamilyElectronics, Magnitude: units.MagnitudeCurrent, Reference: true, Scale: 1, Kind: units.TransformLinear, Source: si},
		{Name: "ohmio", Symbol: "Ω", Aliases: []string{"ohm", "ohms"}, Family: units.FamilyElectronics, Magnitude: units.MagnitudeResistance, Reference: true, Scale: 1, Kind: units.TransformLinear, MinValue: &zero, Source: si},
		{Name: "kiloohmio", Symbol: "kΩ", Aliases: []string{"kohm", "kohms"}, Family: units.FamilyElectronics, Magnitude: units.MagnitudeResistance, Scale: 1e3, Kind: units.TransformLinear, MinValue: &zero, Source: si},
		{Name: "megaohmio", Symbol: "MΩ", Aliases: []string{"Mohm", "Mohms"}, Family: units.FamilyElectronics, Magnitude: units.MagnitudeResistance, Scale: 1e6, Kind: units.TransformLinear, MinValue: &zero, Source: si},
		{Name: "miliwatt", Symbol: "mW", Family: units.FamilyElectronics, Magnitude: units.MagnitudePower, Scale: 1e-3, Kind: units.TransformLinear, MinValue: &zero, Source: si},
		{Name: "watt", Symbol: "W", Family: units.FamilyElectronics, Magnitude: units.MagnitudePower, Reference: true, Scale: 1, Kind: units.TransformLinear, MinValue: &zero, Source: si},
		{Name: "kilowatt", Symbol: "kW", Family: units.FamilyElectronics, Magnitude: units.MagnitudePower, Scale: 1e3, Kind: units.TransformLinear, MinValue: &zero, Source: si},
	}
}
