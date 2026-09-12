package catalog

import (
	"math"

	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

// AngleUnits define las unidades correspondientes al ángulo plano.
//
// Recibe:
//   - No recibe datos externos.
//
// Produce:
//   - El catálogo de unidades angulares reconocidas por UConversor.
//
// Etapa anterior:
//   - Definición de FamilyAngle y MagnitudePlaneAngle.
//
// Etapa siguiente:
//   - Registro global mediante catalog.go.
//
// Restricciones:
//   - El radián es la unidad de referencia.
//   - Todas las conversiones son lineales.
//   - Las relaciones basadas en π deben expresarse con math.Pi.
func AngleUnits() []units.Unit {

	bipm := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "SI Brochure, 9th edition, version 4.01",
		URL:       "https://www.bipm.org/en/publications/si-brochure",
		Notes:     "El radián es la unidad SI coherente del ángulo plano.",
	}

	nist := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "NIST",
		Reference: "NIST Guide to the SI, Appendix B.9 - Angle",
		URL:       "https://www.nist.gov/pml/special-publication-811/nist-guide-si-appendix-b-conversion-factors/nist-guide-si-appendix-b9",
		Notes:     "Relaciones de conversión para grado, gon y revolución.",
	}

	return []units.Unit{
		{
			Name:      "radián",
			Symbol:    "rad",
			Aliases:   []string{"radian", "radianes"},
			Family:    units.FamilyAngle,
			Magnitude: units.MagnitudePlaneAngle,
			Reference: true,
			Kind:      units.TransformLinear,
			Scale:     1,
			Source:    bipm,
		},
		{
			Name:      "grado",
			Symbol:    "°",
			Aliases:   []string{"deg", "grado", "grados"},
			Family:    units.FamilyAngle,
			Magnitude: units.MagnitudePlaneAngle,
			Kind:      units.TransformLinear,
			Scale:     math.Pi / 180,
			Source:    bipm,
		},
		{
			Name:      "gon",
			Symbol:    "gon",
			Aliases:   []string{"grad", "gradian"},
			Family:    units.FamilyAngle,
			Magnitude: units.MagnitudePlaneAngle,
			Kind:      units.TransformLinear,
			Scale:     math.Pi / 200,
			Source:    nist,
		},
		{
			Name:      "revolución",
			Symbol:    "rev",
			Aliases:   []string{"revolution", "vuelta", "vueltas"},
			Family:    units.FamilyAngle,
			Magnitude: units.MagnitudePlaneAngle,
			Kind:      units.TransformLinear,
			Scale:     2 * math.Pi,
			Source:    nist,
		},
	}
}
