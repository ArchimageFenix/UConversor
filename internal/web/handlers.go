// File: handlers.go
//
// Responsibility:
//   - Handle HTTP requests from the browser.
//   - Coordinate requests with the UConversor application layer.
//   - Prepare presentation data for HTML templates.
//
// Receives:
//   - HTTP requests.
//   - Application service.
//   - Parsed HTML templates.
//
// Produces:
//   - Rendered HTML responses.
//
// Previous logical stage:
//   - server.go routing.
//
// Next logical stage:
//   - app.App, errors.go, viewmodel.go and render.go.
//
// Important restrictions:
//   - Must not implement conversion formulas.
//   - Must not recognize units independently.
//   - Must not duplicate catalog rules.
//   - Scientific processing must be delegated to app.Convert().
//   - Error classification must be delegated to errors.go.
//   - Template rendering must be delegated to render.go.
//   - Internal errors and HTTP status codes must not be exposed
//     directly to the user.
package web

import (
	"html/template"
	"net/http"
	"strings"

	"unit-converter/internal/app"
	"unit-converter/internal/input"
)

// Handler coordinates HTTP requests with the
// UConversor application layer and the HTML presentation layer.
//
// It receives requests from server.go and delegates:
//   - scientific processing to app.App;
//   - error classification to errors.go;
//   - presentation adaptation to viewmodel.go;
//   - HTML rendering to render.go.
type Handler struct {
	templates   *template.Template
	application *app.App
}

// NewHandler creates a configured web request handler.
//
// Receives:
//   - Parsed HTML templates.
//   - UConversor application service.
//
// Produces:
//   - Handler ready to be registered by server.go.
//
// Important restrictions:
//   - Does not create the application service.
//   - Does not load templates from disk or embedded assets.
//   - Dependency construction belongs to cmd/web/main.go.
func NewHandler(
	templates *template.Template,
	application *app.App,
) *Handler {
	return &Handler{
		templates:   templates,
		application: application,
	}
}

// handleHome renders the initial UConversor page.
//
// Receives:
//   - GET request for "/".
//
// Produces:
//   - Initial HTML page with an empty PageViewModel.
//
// Previous logical stage:
//   - server.go route registration.
//
// Next logical stage:
//   - render.go.
//
// Important restrictions:
//   - Does not perform a conversion.
//   - Does not inspect or recognize units.
//   - Does not render templates directly.
//   - Only GET is accepted for this route.
func (h *Handler) handleHome(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	viewModel := PageViewModel{}

	h.renderPage(
		w,
		http.StatusOK,
		viewModel,
	)
}

// handleExamples renders the page that lists the units currently
// supported by UConversor.
//
// Receives:
//   - GET request for "/examples".
//
// Produces:
//   - ExamplesViewModel built from the units exposed by app.App.
//
// Previous logical stage:
//   - server.go route registration.
//
// Next logical stages:
//   - app.App.SupportedUnits().
//   - examples_viewmodel.go.
//   - render.go.
//
// Important restrictions:
//   - Must not inspect catalog files directly.
//   - Must not recognize units independently.
//   - Must not duplicate supported-unit lists.
//   - Must not perform conversions.
//   - Must not render templates directly.
//   - Only GET is accepted for this route.
func (h *Handler) handleExamples(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.URL.Path != "/examples" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	defs := h.application.SupportedUnits()
	viewModel := NewExamplesViewModel(defs)

	h.renderExamplesPage(
		w,
		http.StatusOK,
		viewModel,
	)
}

// handleConvert processes a conversion request submitted by the browser.
//
// Receives:
//   - POST request for "/convert".
//   - Form field "expression" containing the user's numeric value
//     and unit.
//
// Produces:
//   - A PageViewModel containing either:
//   - a conversion result, or
//   - a controlled presentation error.
//
// Previous logical stage:
//   - server.go route registration.
//
// Next logical stages:
//   - app.App for conversion processing.
//   - errors.go when conversion fails.
//   - viewmodel.go when conversion succeeds.
//   - render.go for the final HTML response.
//
// Important restrictions:
//   - Must not parse units independently.
//   - Must not perform calculations or conversions.
//   - Must not inspect catalog definitions.
//   - Must not expose raw internal errors.
//   - Must preserve the user's original expression in the form.
//   - HTTP status codes are protocol information only and must not
//     be rendered visually to the user.
func (h *Handler) handleConvert(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.URL.Path != "/convert" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	expression := r.FormValue("expression")

	// The original text is preserved for presentation.
	// TrimSpace is used only to detect an empty browser submission.
	if strings.TrimSpace(expression) == "" {
		status, errorViewModel := classifyConversionError(
			input.ErrEmptyInput,
		)

		viewModel := PageViewModel{
			Input:    expression,
			HasError: true,
			Error:    &errorViewModel,
		}

		h.renderPage(
			w,
			status,
			viewModel,
		)
		return
	}

	result, err := h.application.Convert(expression)
	if err != nil {
		status, errorViewModel := classifyConversionError(err)

		viewModel := PageViewModel{
			Input:    expression,
			HasError: true,
			Error:    &errorViewModel,
		}

		h.renderPage(
			w,
			status,
			viewModel,
		)
		return
	}

	resultViewModel := NewResultViewModel(result)

	viewModel := PageViewModel{
		Input:     expression,
		HasResult: true,
		Result:    &resultViewModel,
	}

	h.renderPage(
		w,
		http.StatusOK,
		viewModel,
	)
}
