// File: handlers.go
//
// Responsibility:
//   - Handle HTTP requests from the browser.
//   - Coordinate requests with the UConversor application layer.
//   - Apply request-level HTTP limits defined by Web security configuration.
//   - Prepare presentation data for HTML templates.
//
// Receives:
//   - HTTP requests.
//   - Application service.
//   - Parsed HTML templates.
//   - Validated Web security configuration.
//
// Produces:
//   - Rendered HTML responses.
//   - Controlled HTTP protocol errors for invalid or oversized requests.
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
//   - Oversized HTTP request bodies must be rejected before
//     reaching app.Convert().
package web

import (
	"errors"
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
//
// It also applies request-level limits defined by SecurityConfig.
type Handler struct {
	templates   *template.Template
	application *app.App
	config      SecurityConfig
}

// NewHandler creates a configured web request handler.
//
// Receives:
//   - Parsed HTML templates.
//   - UConversor application service.
//   - Validated Web security configuration.
//
// Produces:
//   - Handler ready to be registered by server.go.
//
// Important restrictions:
//   - Does not create the application service.
//   - Does not load templates from disk or embedded assets.
//   - Does not load or validate security.yaml directly.
//   - Dependency construction belongs to cmd/web/main.go.
func NewHandler(
	templates *template.Template,
	application *app.App,
	config SecurityConfig,
) *Handler {
	return &Handler{
		templates:   templates,
		application: application,
		config:      config,
	}
}

// handleHome renders the initial UConversor page.
//
// Receives:
//   - GET request for "/".
//
// Produces:
//   - Initial HTML page with an empty PageViewModel.
//   - ActivePage set explicitly to "home" for navigation state.
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

	viewModel := PageViewModel{
		ActivePage: "home",
	}

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

// handleLearning renders the UConversor Learning page.
//
// Receives:
//   - GET request for "/aprendizaje".
//
// Produces:
//   - LearningViewModel prepared for the educational Web page.
//
// Previous logical stage:
//   - server.go route registration.
//
// Next logical stages:
//   - learning_viewmodel.go.
//   - render.go.
//
// Important restrictions:
//   - Must not call app.App.
//   - Must not perform conversions.
//   - Must not inspect the unit Registry or catalog.
//   - Must not define or calculate scientific formulas.
//   - Must not render templates directly.
//   - Only GET is accepted for this route.
func (h *Handler) handleLearning(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.URL.Path != "/aprendizaje" {
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

	viewModel := NewLearningViewModel()

	h.renderLearningPage(
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
//   - ActivePage set explicitly to "home" because conversion
//     results belong to the main page.
//   - HTTP 413 when the request body exceeds the configured limit.
//   - HTTP 400 when the form cannot be parsed.
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
//   - Request size must be validated before calling app.Convert().
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

	r.Body = http.MaxBytesReader(
		w,
		r.Body,
		h.config.HTTP.MaxBodyBytes,
	)

	if err := r.ParseForm(); err != nil {
		var maxBytesError *http.MaxBytesError

		if errors.As(err, &maxBytesError) {
			http.Error(
				w,
				"Payload Too Large",
				http.StatusRequestEntityTooLarge,
			)
			return
		}

		http.Error(
			w,
			"Solicitud inválida",
			http.StatusBadRequest,
		)
		return
	}

	expression := r.Form.Get("expression")

	// The original text is preserved for presentation.
	// TrimSpace is used only to detect an empty browser submission.
	if strings.TrimSpace(expression) == "" {
		status, errorViewModel := classifyConversionError(
			input.ErrEmptyInput,
		)

		viewModel := PageViewModel{
			ActivePage: "home",
			Input:      expression,
			HasError:   true,
			Error:      &errorViewModel,
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
			ActivePage: "home",
			Input:      expression,
			HasError:   true,
			Error:      &errorViewModel,
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
		ActivePage: "home",
		Input:      expression,
		HasResult:  true,
		Result:     &resultViewModel,
	}

	h.renderPage(
		w,
		http.StatusOK,
		viewModel,
	)
}

// handleHealth reports whether the Web HTTP service is alive.
//
// Receives:
//   - GET request for "/health".
//
// Produces:
//   - HTTP 200 with a minimal plain-text response.
//
// Previous logical stage:
//   - server.go route registration.
//
// Next logical stage:
//   - HTTP client or Render health-check system.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not query the catalog.
//   - Must not call app.Convert().
//   - Must remain lightweight.
//   - Only GET is accepted for this route.
func (h *Handler) handleHealth(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"text/plain; charset=utf-8",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	_, _ = w.Write(
		[]byte("OK"),
	)
}

// handleTutorial resolves and serves the static thematic tutorials
// available in the UConversor Learning area.
//
// Responsibility:
//   - Receive requests under /aprendizaje/{tema}.
//   - Validate the requested tutorial against a controlled Web-side
//     list.
//   - Build the corresponding TutorialViewModel.
//   - Delegate HTML generation to the tutorial renderer.
//
// Receives:
//   - HTTP GET requests for a supported Learning tutorial.
//
// Produces:
//   - A rendered static tutorial page.
//   - HTTP 404 when the requested tutorial does not exist.
//   - HTTP 405 when the request method is not GET.
//
// Previous logical stage:
//   - server.go.
//
// Next logical stage:
//   - TutorialViewModel and renderTutorialPage().
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not call App.Convert().
//   - Must not inspect Registry or Catalog.
//   - Must not send arbitrary URL values directly to the template
//     system.
//   - Every tutorial must be explicitly allowed by this Web layer.
func (h *Handler) handleTutorial(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var tutorial string

	switch r.URL.Path {

	case "/aprendizaje/longitud":
		tutorial = "length"

	default:
		http.NotFound(w, r)
		return
	}

	viewModel := NewTutorialViewModel(tutorial)

	h.renderTutorialPage(
		w,
		http.StatusOK,
		viewModel,
	)
}
