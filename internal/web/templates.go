// File: templates.go
//
// Responsibility:
//   - Load and parse the HTML templates embedded in UConversor Web.
//
// Receives:
//   - Embedded template files provided by assets.go.
//
// Produces:
//   - Parsed template set ready for use by the web handlers.
//
// Previous logical stage:
//   - assets.go.
//
// Next logical stage:
//   - handlers.go.
//
// Important restrictions:
//   - Must not contain HTTP routing logic.
//   - Must not perform conversions.
//   - Must not contain scientific or catalog data.
package web

import (
	"fmt"
	"html/template"
)

// LoadTemplates parses the embedded HTML templates
// required by the UConversor web presentation layer.
func LoadTemplates() (*template.Template, error) {
	templates, err := template.ParseFS(
		assets,
		"templates/*.html",
		"templates/partials/*.html",
	)
	if err != nil {
		return nil, fmt.Errorf(
			"cargar plantillas web: %w",
			err,
		)
	}

	return templates, nil
}
