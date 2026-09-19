// File: frequency.go
//
// Responsibility:
//   - Define the frequency units supported by UConversor.
//   - Associate each frequency unit with its SI scale and scientific source.
//
// Receives:
//   - No runtime input. This file declares static catalog data.
//
// Produces:
//   - A collection of frequency units used by the central catalog.
//
// Previous logical stage:
//   - units/magnitude.go and units/family.go define the classifications used here.
//
// Next logical stage:
//   - catalog/catalog.go aggregates these units.
//   - units/registry.go later exposes them for recognition and conversion.
//
// Important restrictions:
//   - All units declared here must represent the same physical magnitude.
//   - Frequency uses the hertz (Hz) as the reference unit.
//   - Conversions in this catalog are linear.
//   - This file must not implement relationships between frequency and
//     other magnitudes, such as wavelength.

package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

// Frequency returns the frequency units supported by UConversor.
func Frequency() []units.Unit {
	bipm := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "SI Brochure, 9th ed., version 4.01",
		URL:       "https://www.bipm.org/en/publications/si-brochure",
	}

	zero := 0.0

	return []units.Unit{
		{
			Name:      "hertz",
			Symbol:    "Hz",
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeFrequency,
			Reference: true,
			Scale:     1,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "kilohertz",
			Symbol:    "kHz",
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeFrequency,
			Scale:     1e3,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "megahertz",
			Symbol:    "MHz",
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeFrequency,
			Scale:     1e6,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "gigahertz",
			Symbol:    "GHz",
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeFrequency,
			Scale:     1e9,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
		{
			Name:      "terahertz",
			Symbol:    "THz",
			Family:    units.FamilyPhysics,
			Magnitude: units.MagnitudeFrequency,
			Scale:     1e12,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipm,
		},
	}
}
