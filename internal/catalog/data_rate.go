// File: data_rate.go
//
// Responsibility:
//   - Define the data transfer rate units recognized by UConversor.
//   - Establish their linear relationships using bit per second
//     as the internal reference unit.
//   - Keep data transfer rate separate from data storage.
//
// Receives:
//   - No external runtime input.
//   - Returns static unit definitions.
//
// Produces:
//   - []units.Unit containing supported data transfer rate units.
//
// Previous logical stage:
//   - units/family.go defines FamilyData.
//   - units/magnitude.go defines MagnitudeDataRate.
//
// Next logical stage:
//   - catalog.go incorporates these definitions into the Registry.
//
// Important restrictions:
//   - Data transfer rate is not the same magnitude as data storage.
//   - Mbit/s and MB/s are not equivalent.
//   - SI decimal prefixes are used for network data rates.
//   - Symbols and aliases must preserve their intended capitalization.
//   - All conversions in this catalog are linear.
package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

// DataRateUnits defines the supported data transfer rate units.
//
// The internal reference unit for this magnitude is bit per second (bit/s).
func DataRateUnits() []units.Unit {
	bipm := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "The International System of Units (SI) – SI prefixes",
		URL:       "https://www.bipm.org/en/measurement-units/si-prefixes",
		Notes:     "Prefijos decimales SI utilizados en tasas de datos: kilo=10^3, mega=10^6 y giga=10^9.",
	}

	iec := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "IEC",
		Reference: "IEC quantities and units for information technology",
		URL:       "https://styleguide.iec.ch/?docs=iec%2Ftypographic%2Funits-and-symbols",
		Notes:     "Referencia para la unidad bit y su símbolo en tecnología de la información.",
	}

	return []units.Unit{
		// ---------------------------------------------------------
		// Unidad de referencia
		// ---------------------------------------------------------

		{
			Name:      "bit por segundo",
			Symbol:    "bit/s",
			Aliases:   []string{"bps", "bitps"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeDataRate,
			Reference: true,
			Kind:      units.TransformLinear,
			Scale:     1,
			Source:    iec,
		},

		// ---------------------------------------------------------
		// Tasas con prefijos decimales SI
		// ---------------------------------------------------------

		{
			Name:      "kilobit por segundo",
			Symbol:    "kbit/s",
			Aliases:   []string{"kbps", "kbitps"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeDataRate,
			Kind:      units.TransformLinear,
			Scale:     1e3,
			Source:    bipm,
		},
		{
			Name:      "megabit por segundo",
			Symbol:    "Mbit/s",
			Aliases:   []string{"Mbps", "Mbitps"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeDataRate,
			Kind:      units.TransformLinear,
			Scale:     1e6,
			Source:    bipm,
		},
		{
			Name:      "gigabit por segundo",
			Symbol:    "Gbit/s",
			Aliases:   []string{"Gbps", "Gbitps"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeDataRate,
			Kind:      units.TransformLinear,
			Scale:     1e9,
			Source:    bipm,
		},
	}
}
