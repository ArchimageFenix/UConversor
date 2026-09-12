package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

// TimeUnits define las unidades pertenecientes a la familia Tiempo.
//
// Responsabilidad:
//   - Declarar las unidades de tiempo reconocidas por UConversor.
//   - Relacionar cada unidad con el segundo, unidad de referencia.
//   - Mantener las fuentes científicas asociadas a dichas relaciones.
//
// Recibe:
//   - No recibe datos externos.
//
// Produce:
//   - []units.Unit con el catálogo completo de unidades de tiempo.
//
// Etapa anterior:
//   - Definiciones de FamilyTime y MagnitudeTime.
//
// Etapa siguiente:
//   - Registry -> Validator -> Conversion Engine.
//
// Restricciones:
//   - No realiza conversiones.
//   - No contiene lógica de presentación.
//   - No modifica el motor.
//   - El año utilizado es el año gregoriano medio de 365.2425 días.
//   - No representa la duración de un año calendario concreto,
//     que puede contener 365 o 366 días.
//   - El mes utilizado es el mes gregoriano medio, equivalente
//     a 1/12 del año gregoriano medio.
//   - El sol representa el día solar medio marciano.
//   - No debe confundirse con el día sideral de Marte.
func TimeUnits() []units.Unit {
	bipmSI := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "The International System of Units (SI Brochure), 9th edition, version 4.01",
		URL:       "https://www.bipm.org/en/publications/si-brochure",
		Notes:     "El segundo es la unidad SI de tiempo. Los prefijos SI son potencias decimales de diez.",
	}

	bipmTraditionalTime := model.Source{
		Kind:      model.SourceExact,
		Authority: "BIPM",
		Reference: "SI Brochure, 9th edition, version 4.01 - Non-SI units",
		URL:       "https://www.bipm.org/en/publications/si-brochure",
		Notes:     "1 min = 60 s; 1 h = 3600 s; 1 d = 86400 s.",
	}

	weekSource := model.Source{
		Kind:      model.SourceDerived,
		Authority: "UConversor",
		Reference: "Derived from BIPM definition of the day",
		URL:       "https://www.bipm.org/en/publications/si-brochure",
		Notes:     "1 week = 7 days = 604800 seconds. Derived from the BIPM relation 1 d = 86400 s.",
	}

	gregorianYear := model.Source{
		Kind:      model.SourceDerived,
		Authority: "NIST",
		Reference: "Time and Frequency from A to Z - Year",
		URL:       "https://www.nist.gov/pml/time-and-frequency-division/popular-links/time-frequency-z/time-and-frequency-z-x-z",
		Notes:     "El año gregoriano medio contiene 365.2425 días. Usando 86400 s por día, equivale a 31556952 s.",
	}

	gregorianMonth := model.Source{
		Kind:      model.SourceDerived,
		Authority: "UConversor",
		Reference: "Derived from the mean Gregorian year documented by NIST",
		URL:       "https://www.nist.gov/pml/time-and-frequency-division/popular-links/time-frequency-z/time-and-frequency-z-x-z",
		Notes:     "1 mes gregoriano medio = 1/12 del año gregoriano medio = 30.436875 días = 2629746 segundos.",
	}

	marsSol := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "NASA/JPL",
		Reference: "Mars 2020 Perseverance Mission Overview",
		URL:       "https://www.jpl.nasa.gov/news/press_kits/mars_2020/launch/mission/",
		Notes:     "Un sol es el día solar medio marciano: 24 h 39 min 35.244 s, equivalente a 88775.244 segundos terrestres.",
	}

	return []units.Unit{
		{
			Name:      "Nanosegundo",
			Symbol:    "ns",
			Aliases:   []string{"nanosecond", "nanoseconds"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     1e-9,
			Source:    bipmSI,
		},
		{
			Name:      "Microsegundo",
			Symbol:    "µs",
			Aliases:   []string{"us", "microsecond", "microseconds"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     1e-6,
			Source:    bipmSI,
		},
		{
			Name:      "Milisegundo",
			Symbol:    "ms",
			Aliases:   []string{"millisecond", "milliseconds"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     1e-3,
			Source:    bipmSI,
		},
		{
			Name:      "Segundo",
			Symbol:    "s",
			Aliases:   []string{"sec", "seg", "second", "seconds"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Reference: true,
			Kind:      units.TransformLinear,
			Scale:     1,
			Source:    bipmSI,
		},
		{
			Name:      "Minuto",
			Symbol:    "min",
			Aliases:   []string{"minute", "minutes"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     60,
			Source:    bipmTraditionalTime,
		},
		{
			Name:      "Hora",
			Symbol:    "h",
			Aliases:   []string{"hr", "hour", "hours"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     3600,
			Source:    bipmTraditionalTime,
		},
		{
			Name:      "Día",
			Symbol:    "d",
			Aliases:   []string{"day", "days", "dia"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     86400,
			Source:    bipmTraditionalTime,
		},
		{
			Name:      "Sol marciano",
			Symbol:    "sol",
			Aliases:   []string{"mars sol", "martian sol"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     88775.244,
			Source:    marsSol,
		},
		{
			Name:      "Mes gregoriano medio",
			Symbol:    "mo",
			Aliases:   []string{"month", "months"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     2629746,
			Source:    gregorianMonth,
		},
		{
			Name:      "Semana",
			Symbol:    "wk",
			Aliases:   []string{"week", "weeks", "semana"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     604800,
			Source:    weekSource,
		},
		{
			Name:      "Año gregoriano medio",
			Symbol:    "yr",
			Aliases:   []string{"year", "yearavg"},
			Family:    units.FamilyTime,
			Magnitude: units.MagnitudeTime,
			Kind:      units.TransformLinear,
			Scale:     31556952,
			Source:    gregorianYear,
		},
	}
}
