// File: length.go
// Responsibility: Declare scientifically traceable length units.
// Receives: No runtime input.
// Produces: Length unit definitions referenced to the metre.
// Previous logical stage: normative source selection.
// Next logical stage: units/registry.go.
// Important restrictions: Factors must match documented scientific definitions.
package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

func LengthUnits() []units.Unit {
	bipm := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "SI Brochure, 9th ed., version 4.01",
		URL:       "https://www.bipm.org/en/publications/si-brochure",
	}

	nist := model.Source{
		Kind:      model.SourceExact,
		Authority: "NIST",
		Reference: "SI Units – Length / 1959 International Yard and Pound",
		URL:       "https://www.nist.gov/pml/owm/si-units-length",
	}

	nasa := model.Source{
		Kind:      model.SourceExact,
		Authority: "NASA Goddard Space Flight Center",
		Reference: "Imagine the Universe! Dictionary – Light Year",
		URL:       "https://imagine.gsfc.nasa.gov/resources/dictionary.html",
		Notes:     "NASA: 1 light-year = 9.46053 × 10^12 km",
	}

	iau := model.Source{
		Kind:      model.SourceExact,
		Authority: "IAU",
		Reference: "IAU 2012 Resolution B2 – Re-definition of the astronomical unit of length",
		URL:       "https://www.iau.org/static/resolutions/IAU2012_English.pdf",
		Notes:     "1 au = 149 597 870 700 m exactly",
	}

	return []units.Unit{
		{
			Name:      "milímetro",
			Symbol:    "mm",
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     0.001,
			Kind:      units.TransformLinear,
			Source:    bipm,
		},
		{
			Name:      "centímetro",
			Symbol:    "cm",
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     0.01,
			Kind:      units.TransformLinear,
			Source:    bipm,
		},
		{
			Name:      "metro",
			Symbol:    "m",
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Reference: true,
			Scale:     1,
			Kind:      units.TransformLinear,
			Source:    bipm,
		},
		{
			Name:      "kilómetro",
			Symbol:    "km",
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     1000,
			Kind:      units.TransformLinear,
			Source:    bipm,
		},
		{
			Name:      "pulgada",
			Symbol:    "in",
			Aliases:   []string{"inch"},
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     0.0254,
			Kind:      units.TransformLinear,
			Source:    nist,
		},
		{
			Name:      "pie",
			Symbol:    "ft",
			Aliases:   []string{"foot"},
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     0.3048,
			Kind:      units.TransformLinear,
			Source:    nist,
		},
		{
			Name:      "yarda",
			Symbol:    "yd",
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     0.9144,
			Kind:      units.TransformLinear,
			Source:    nist,
		},
		{
			Name:      "milla",
			Symbol:    "mi",
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     1609.344,
			Kind:      units.TransformLinear,
			Source:    nist,
		},
		{
			Name:      "año luz",
			Symbol:    "ly",
			Aliases:   []string{"lightyear", "light-year"},
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     9.46053e15,
			Kind:      units.TransformLinear,
			Source:    nasa,
		},

		{
			Name:      "unidad astronómica",
			Symbol:    "au",
			Aliases:   []string{"ua", "UA"},
			Family:    units.FamilyLength,
			Magnitude: units.MagnitudeLength,
			Scale:     149597870700,
			Kind:      units.TransformLinear,
			Source:    iau,
		},
	}
}
