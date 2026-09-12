// File: scientific_validation_test.go
// Responsibility: Ensure every registered unit has traceable scientific metadata.
// Receives: Complete catalog definitions.
// Produces: Automated evidence that no anonymous factor enters production.
// Previous logical stage: catalog definitions.
// Next logical stage: maintenance/release validation.
// Important restrictions: Missing authority/reference is a test failure.
package tests

import (
	"testing"
	"unit-converter/internal/catalog"
)

func TestAllUnitsHaveScientificSource(t *testing.T) {
	for _, u := range catalog.All() {
		if u.Source.Authority == "" || u.Source.Reference == "" || u.Source.URL == "" {
			t.Fatalf("unit %s has incomplete source metadata", u.Symbol)
		}
	}
}
