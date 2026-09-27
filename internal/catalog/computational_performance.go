// File: computational_performance.go
// Responsibility: Declare computational performance units with FLOP per second as reference.
// Receives: No runtime input.
// Produces: Computational performance unit definitions.
// Previous logical stage: normative source selection.
// Next logical stage: units/registry.go.
// Important restrictions:
//   - Negative computational performance values are rejected.
//   - Units express floating-point operations per second.
//   - Conversions do not compare hardware architectures.
//   - Precision formats such as FP16, FP32 and FP64 are outside this conversion scope.
//   - Theoretical and measured performance are not treated as equivalent claims.
package catalog

import (
	"unit-converter/internal/model"
	"unit-converter/internal/units"
)

// ComputationalPerformanceUnits returns the supported units
// for floating-point computational performance.
func ComputationalPerformanceUnits() []units.Unit {
	top500 := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "TOP500",
		Reference: "Frequently Asked Questions — Mflop/s, Gflop/s and Tflop/s",
		URL:       "https://top500.org/resources/frequently-asked-questions/",
		Notes:     "FLOP/s represents a rate of floating-point operations executed per second.",
	}

	bipmPrefixes := model.Source{
		Kind:      model.SourceOfficial,
		Authority: "BIPM",
		Reference: "The International System of Units — SI prefixes",
		URL:       "https://www.bipm.org/en/measurement-units/si-prefixes",
		Notes:     "Decimal prefix factors are exact: kilo 10^3, mega 10^6, giga 10^9, tera 10^12, peta 10^15 and exa 10^18.",
	}

	zero := 0.0

	return []units.Unit{
		{
			Name:      "FLOP por segundo",
			Symbol:    "FLOP/s",
			Aliases:   []string{"FLOPS", "flop/s"},
			Family:    units.FamilyComputing,
			Magnitude: units.MagnitudeComputationalPerformance,
			Reference: true,
			Scale:     1,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    top500,
		},
		{
			Name:      "kiloFLOP por segundo",
			Symbol:    "kFLOP/s",
			Aliases:   []string{"kFLOPS", "kflop/s"},
			Family:    units.FamilyComputing,
			Magnitude: units.MagnitudeComputationalPerformance,
			Scale:     1e3,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipmPrefixes,
		},
		{
			Name:      "megaFLOP por segundo",
			Symbol:    "MFLOP/s",
			Aliases:   []string{"MFLOPS", "Mflop/s"},
			Family:    units.FamilyComputing,
			Magnitude: units.MagnitudeComputationalPerformance,
			Scale:     1e6,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipmPrefixes,
		},
		{
			Name:      "gigaFLOP por segundo",
			Symbol:    "GFLOP/s",
			Aliases:   []string{"GFLOPS", "Gflop/s"},
			Family:    units.FamilyComputing,
			Magnitude: units.MagnitudeComputationalPerformance,
			Scale:     1e9,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipmPrefixes,
		},
		{
			Name:      "teraFLOP por segundo",
			Symbol:    "TFLOP/s",
			Aliases:   []string{"TFLOPS", "Tflop/s"},
			Family:    units.FamilyComputing,
			Magnitude: units.MagnitudeComputationalPerformance,
			Scale:     1e12,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipmPrefixes,
		},
		{
			Name:      "petaFLOP por segundo",
			Symbol:    "PFLOP/s",
			Aliases:   []string{"PFLOPS", "Pflop/s"},
			Family:    units.FamilyComputing,
			Magnitude: units.MagnitudeComputationalPerformance,
			Scale:     1e15,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipmPrefixes,
		},
		{
			Name:      "exaFLOP por segundo",
			Symbol:    "EFLOP/s",
			Aliases:   []string{"EFLOPS", "Eflop/s"},
			Family:    units.FamilyComputing,
			Magnitude: units.MagnitudeComputationalPerformance,
			Scale:     1e18,
			Kind:      units.TransformLinear,
			MinValue:  &zero,
			Source:    bipmPrefixes,
		},
	}
}
