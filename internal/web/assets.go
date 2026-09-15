// File: assets.go
//
// Responsibility:
//   - Embed the web templates and static resources into the UConversor binary.
//   - Provide access to those resources to the HTTP server.
//
// Receives:
//   - Files located under templates/ and static/ at compile time.
//
// Produces:
//   - Embedded filesystem containing the web interface.
//
// Previous logical stage:
//   - Web source files.
//
// Next logical stage:
//   - server.go loads templates and serves static resources.
//
// Important restrictions:
//   - Contains no HTTP routing logic.
//   - Contains no conversion logic.
//   - Contains no scientific data.
package web

import "embed"

// assets contains all HTML templates and static web resources.
//
//go:embed templates/*.html templates/partials/*.html static/css/*.css static/css/visuals/*.css
var assets embed.FS
