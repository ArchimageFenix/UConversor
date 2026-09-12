package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

// DataUnits define las unidades correspondientes a cantidad de información
// digital.
//
// Responsabilidad:
//   - Declarar las unidades de datos reconocidas por UConversor.
//   - Definir sus relaciones exactas respecto al bit, que actúa como
//     unidad de referencia interna.
//   - Mantener separados los prefijos decimales SI y los prefijos
//     binarios IEC.
//
// Recibe:
//   - No recibe datos externos. Devuelve definiciones estáticas de unidades.
//
// Produce:
//   - []units.Unit con todas las unidades de información digital.
//
// Etapa lógica anterior:
//   - units/family.go y units/magnitude.go definen FamilyData y MagnitudeData.
//
// Etapa lógica siguiente:
//   - catalog.go incorpora estas unidades al Registry.
//
// Restricciones:
//   - B (byte) y bit no son equivalentes: 1 B = 8 bit.
//   - MB y MiB NO son aliases.
//   - GB y GiB NO son aliases.
//   - Debe respetarse estrictamente mayúsculas y minúsculas de los símbolos.
//   - Todas las conversiones de esta familia son lineales.

func DataUnits() []units.Unit {
	bipm := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "The International System of Units (SI) – SI prefixes",
		URL:       "https://www.bipm.org/en/measurement-units/si-prefixes",
		Notes:     "Prefijos decimales SI: kilo=10^3, mega=10^6, giga=10^9, tera=10^12 y peta=10^15.",
	}

	iec := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "IEC",
		Reference: "IEC quantities and units for information technology",
		URL:       "https://styleguide.iec.ch/?docs=iec%2Ftypographic%2Funits-and-symbols",
		Notes:     "Fuente IEC para bit, byte y prefijos binarios. Los prefijos binarios siguen potencias de 2: kibi=2^10, mebi=2^20, gibi=2^30, tebi=2^40 y pebi=2^50.",
	}

	return []units.Unit{
		// ---------------------------------------------------------
		// Unidad de referencia
		// ---------------------------------------------------------

		{
			Name:      "bit",
			Symbol:    "bit",
			Aliases:   []string{"bits"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Reference: true,
			Kind:      units.TransformLinear,
			Scale:     1,
			Source:    iec,
		},

		// ---------------------------------------------------------
		// Byte
		// ---------------------------------------------------------

		{
			Name:      "byte",
			Symbol:    "B",
			Aliases:   []string{"byte", "bytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8,
			Source:    iec,
		},

		// ---------------------------------------------------------
		// Bits — prefijos decimales
		// ---------------------------------------------------------

		{
			Name:      "kilobit",
			Symbol:    "kbit",
			Aliases:   []string{"kilobit", "kilobits"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     1e3,
			Source:    bipm,
		},
		{
			Name:      "megabit",
			Symbol:    "Mbit",
			Aliases:   []string{"megabit", "megabits"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     1e6,
			Source:    bipm,
		},
		{
			Name:      "gigabit",
			Symbol:    "Gbit",
			Aliases:   []string{"gigabit", "gigabits"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     1e9,
			Source:    bipm,
		},
		{
			Name:      "terabit",
			Symbol:    "Tbit",
			Aliases:   []string{"Tb", "terabit", "terabits"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     1e12,
			Source:    bipm,
		},
		{
			Name:      "petabit",
			Symbol:    "Pbit",
			Aliases:   []string{"petabit", "petabits"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     1e15,
			Source:    bipm,
		},

		// ---------------------------------------------------------
		// Bytes — prefijos decimales
		// ---------------------------------------------------------

		{
			Name:      "kilobyte",
			Symbol:    "kB",
			Aliases:   []string{"kilobyte", "kilobytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8e3,
			Source:    bipm,
		},
		{
			Name:      "megabyte",
			Symbol:    "MB",
			Aliases:   []string{"megabyte", "megabytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8e6,
			Source:    bipm,
		},
		{
			Name:      "gigabyte",
			Symbol:    "GB",
			Aliases:   []string{"gigabyte", "gigabytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8e9,
			Source:    bipm,
		},
		{
			Name:      "terabyte",
			Symbol:    "TB",
			Aliases:   []string{"terabyte", "terabytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8e12,
			Source:    bipm,
		},
		{
			Name:      "petabyte",
			Symbol:    "PB",
			Aliases:   []string{"petabyte", "petabytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8e15,
			Source:    bipm,
		},

		// ---------------------------------------------------------
		// Bytes — prefijos binarios IEC
		// ---------------------------------------------------------

		{
			Name:      "kibibyte",
			Symbol:    "KiB",
			Aliases:   []string{"kibibyte", "kibibytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8 * 1024,
			Source:    iec,
		},
		{
			Name:      "mebibyte",
			Symbol:    "MiB",
			Aliases:   []string{"mebibyte", "mebibytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8 * 1024 * 1024,
			Source:    iec,
		},
		{
			Name:      "gibibyte",
			Symbol:    "GiB",
			Aliases:   []string{"gibibyte", "gibibyte"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8 * 1024 * 1024 * 1024,
			Source:    iec,
		},
		{
			Name:      "tebibyte",
			Symbol:    "TiB",
			Aliases:   []string{"tebibyte", "tebibytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8 * 1024 * 1024 * 1024 * 1024,
			Source:    iec,
		},
		{
			Name:      "pebibyte",
			Symbol:    "PiB",
			Aliases:   []string{"pebibyte", "pebibytes"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeData,
			Kind:      units.TransformLinear,
			Scale:     8 * 1024 * 1024 * 1024 * 1024 * 1024,
			Source:    iec,
		},

		{
			Name:      "megabyte por segundo",
			Symbol:    "MB/s",
			Aliases:   []string{"MBps"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeDataRate,
			Kind:      units.TransformLinear,
			Scale:     8e6,
			Source:    bipm,
		},

		{
			Name:      "byte por segundo",
			Symbol:    "B/s",
			Aliases:   []string{"Bps"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeDataRate,
			Kind:      units.TransformLinear,
			Scale:     8,
			Source:    iec,
		},

		{
			Name:      "kilobyte por segundo",
			Symbol:    "kB/s",
			Aliases:   []string{"kBps"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeDataRate,
			Kind:      units.TransformLinear,
			Scale:     8e3,
			Source:    bipm,
		},

		{
			Name:      "gigabyte por segundo",
			Symbol:    "GB/s",
			Aliases:   []string{"GBps"},
			Family:    units.FamilyData,
			Magnitude: units.MagnitudeDataRate,
			Kind:      units.TransformLinear,
			Scale:     8e9,
			Source:    bipm,
		},
	}

}
