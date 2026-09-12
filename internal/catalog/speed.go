// File: speed.go
// Responsibility: Declare speed units with metre per second as reference.
// Receives: No runtime input.
// Produces: Speed unit definitions.
// Previous logical stage: normative source selection.
// Next logical stage: units/registry.go.
// Important restrictions: Negative speed is rejected; directional velocity is outside this initial scope.
package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func SpeedUnits() []units.Unit {
	bipm := model.Source{Kind: model.SourceOfficial, Authority: "BIPM", Reference: "SI Brochure, 9th ed., version 4.01; non-SI units table", URL: "https://www.bipm.org/en/publications/si-brochure"}
	nist := model.Source{Kind: model.SourceExact, Authority: "NIST", Reference: "International yard and pound definitions", URL: "https://www.nist.gov/pml/owm/si-units-length"}
	zero := 0.0
	return []units.Unit{
		{Name: "metro por segundo", Symbol: "m/s", Family: units.FamilySpeed, Magnitude: units.MagnitudeSpeed, Reference: true, Scale: 1, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "kilómetro por segundo", Symbol: "km/s", Family: units.FamilySpeed, Magnitude: units.MagnitudeSpeed, Scale: 1000.0, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "kilómetro por hora", Symbol: "km/h", Aliases: []string{"kph"}, Family: units.FamilySpeed, Magnitude: units.MagnitudeSpeed, Scale: 1000.0 / 3600.0, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "milla por hora", Symbol: "mph", Aliases: []string{"mi/h"}, Family: units.FamilySpeed, Magnitude: units.MagnitudeSpeed, Scale: 1609.344 / 3600.0, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
		{Name: "nudo", Symbol: "kn", Aliases: []string{"knot"}, Family: units.FamilySpeed, Magnitude: units.MagnitudeSpeed, Scale: 1852.0 / 3600.0, Kind: units.TransformLinear, MinValue: &zero, Source: bipm},
		{Name: "pie por segundo", Symbol: "ft/s", Family: units.FamilySpeed, Magnitude: units.MagnitudeSpeed, Scale: 0.3048, Kind: units.TransformLinear, MinValue: &zero, Source: nist},
	}
}
