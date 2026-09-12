// File: examples_viewmodel.go
//
// Responsibility:
//   - Transform supported unit definitions into presentation data
//     for the /examples web view.
//   - Group units by family and magnitude for readable display.
//
// Receives:
//   - []units.Unit provided by the application layer.
//
// Produces:
//   - ExamplesViewModel consumed by the examples HTML template.
//
// Previous logical stage:
//   - app.App.SupportedUnits().
//
// Next logical stage:
//   - handlers.go and the /examples HTML template.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not define scientific constants or formulas.
//   - Must not duplicate catalog data.
//   - Must not modify unit definitions.
//   - Must remain a presentation model only.
package web

import (
	"sort"

	"unit-converter/internal/units"
)

type ExamplesViewModel struct {
	ActivePage string
	Groups     []ExampleFamilyViewModel
}

type ExampleFamilyViewModel struct {
	Family     string
	Magnitudes []ExampleMagnitudeViewModel
}

type ExampleMagnitudeViewModel struct {
	Magnitude string
	Units     []ExampleUnitViewModel
}

type ExampleUnitViewModel struct {
	Name   string
	Symbol string
}

func NewExamplesViewModel(defs []units.Unit) ExamplesViewModel {
	families := make(map[string]map[string][]ExampleUnitViewModel)

	for _, unit := range defs {
		family := string(unit.Family)
		magnitude := string(unit.Magnitude)

		if _, ok := families[family]; !ok {
			families[family] = make(map[string][]ExampleUnitViewModel)
		}

		families[family][magnitude] = append(
			families[family][magnitude],
			ExampleUnitViewModel{
				Name:   unit.Name,
				Symbol: unit.Symbol,
			},
		)
	}

	familyNames := make([]string, 0, len(families))
	for family := range families {
		familyNames = append(familyNames, family)
	}
	sort.Strings(familyNames)

	groups := make([]ExampleFamilyViewModel, 0, len(familyNames))

	for _, family := range familyNames {
		magnitudesMap := families[family]

		magnitudeNames := make([]string, 0, len(magnitudesMap))
		for magnitude := range magnitudesMap {
			magnitudeNames = append(magnitudeNames, magnitude)
		}
		sort.Strings(magnitudeNames)

		magnitudes := make(
			[]ExampleMagnitudeViewModel,
			0,
			len(magnitudeNames),
		)

		for _, magnitude := range magnitudeNames {
			unitList := magnitudesMap[magnitude]

			sort.Slice(unitList, func(i, j int) bool {
				return unitList[i].Name < unitList[j].Name
			})

			magnitudes = append(
				magnitudes,
				ExampleMagnitudeViewModel{
					Magnitude: magnitude,
					Units:     unitList,
				},
			)
		}

		groups = append(
			groups,
			ExampleFamilyViewModel{
				Family:     family,
				Magnitudes: magnitudes,
			},
		)
	}

	return ExamplesViewModel{
		ActivePage: "examples",
		Groups:     groups,
	}
}
