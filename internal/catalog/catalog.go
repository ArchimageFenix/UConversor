// File: catalog.go
// Responsibility: Aggregate all initial unit catalogs into one registry input.
// Receives: No runtime user input.
// Produces: Complete initial []units.Unit catalog.
// Previous logical stage: individual catalog files.
// Next logical stage: units/registry.go during app construction.
// Important restrictions: Aggregates only; contains no conversion formulas of its own.
package catalog

import "unit-converter/internal/units"

func All() []units.Unit {
	var all []units.Unit
	all = append(all, LengthUnits()...)
	all = append(all, MassUnits()...)
	all = append(all, SpeedUnits()...)
	all = append(all, ElectronicsUnits()...)
	all = append(all, PhysicsUnits()...)
	all = append(all, ForceUnits()...)
	all = append(all, DataUnits()...)
	all = append(all, AngleUnits()...)
	all = append(all, TimeUnits()...)
	all = append(all, DataRateUnits()...)
	return all
}
