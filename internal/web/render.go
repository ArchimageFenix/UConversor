// File: render.go
//
// Responsibility:
//   - Render HTML templates safely for HTTP responses.
//   - Centralize HTTP status handling for rendered web pages.
//   - Provide dedicated rendering paths for the Home page
//     and the Examples page.
//
// Receives:
//   - HTTP response writer.
//   - HTTP status code.
//   - Presentation view models prepared by handlers.go.
//
// Produces:
//   - Complete HTML responses ready to be sent to the browser.
//
// Previous logical stage:
//   - handlers.go.
//
// Next logical stage:
//   - HTML templates.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not classify domain errors.
//   - Must not contain scientific logic.
//   - Must not inspect catalog definitions.
//   - Must not modify application or registry state.
//   - Must not prepare presentation view models.
//   - Template rendering must complete successfully before
//     writing the HTTP status and response body.
package web

import (
	"bytes"
	"net/http"
)

// renderPage renders the main UConversor page.
//
// Receives:
//   - HTTP response writer.
//   - HTTP status code.
//   - PageViewModel prepared by handlers.go.
//
// Produces:
//   - Complete HTML response based on the "layout" template.
//
// Previous logical stage:
//   - handleHome() or handleConvert().
//
// Next logical stage:
//   - Main HTML layout and its partial templates.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not classify application errors.
//   - Must not modify the PageViewModel.
//   - Must render the complete template into memory before
//     writing the HTTP response.
func (h *Handler) renderPage(
	w http.ResponseWriter,
	status int,
	viewModel PageViewModel,
) {
	var buffer bytes.Buffer

	if err := h.templates.ExecuteTemplate(
		&buffer,
		"layout",
		viewModel,
	); err != nil {
		http.Error(
			w,
			"Error interno al generar la interfaz",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)

	w.WriteHeader(status)

	_, _ = w.Write(buffer.Bytes())
}

// renderExamplesPage renders the UConversor supported-units page.
//
// Receives:
//   - HTTP response writer.
//   - HTTP status code.
//   - ExamplesViewModel prepared by handleExamples().
//
// Produces:
//   - Complete HTML response based on the "examples" template.
//
// Previous logical stage:
//   - handlers.go -> handleExamples().
//
// Next logical stage:
//   - Examples HTML template.
//
// Important restrictions:
//   - Must not inspect catalog definitions directly.
//   - Must not perform conversions.
//   - Must not recognize units.
//   - Must not modify the ExamplesViewModel.
//   - Must not contain presentation data that belongs in
//     examples_viewmodel.go.
//   - Must render the complete template into memory before
//     writing the HTTP response.
func (h *Handler) renderExamplesPage(
	w http.ResponseWriter,
	status int,
	viewModel ExamplesViewModel,
) {
	var buffer bytes.Buffer

	if err := h.templates.ExecuteTemplate(
		&buffer,
		"examples",
		viewModel,
	); err != nil {
		http.Error(
			w,
			"Error interno al generar la interfaz",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)

	w.WriteHeader(status)

	_, _ = w.Write(buffer.Bytes())
}

// renderLearningPage renders the UConversor Learning page.
//
// Receives:
//   - HTTP response writer.
//   - HTTP status code.
//   - LearningViewModel prepared by handleLearning().
//
// Produces:
//   - Complete HTML response based on the "learning" template.
//
// Previous logical stage:
//   - handlers.go -> handleLearning().
//
// Next logical stage:
//   - Learning HTML template.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not access app.App.
//   - Must not inspect the unit Registry or catalog.
//   - Must not contain scientific formulas or calculations.
//   - Must not modify the LearningViewModel.
//   - Must render the complete template into memory before
//     writing the HTTP response.
func (h *Handler) renderLearningPage(
	w http.ResponseWriter,
	status int,
	viewModel LearningViewModel,
) {
	var buffer bytes.Buffer

	if err := h.templates.ExecuteTemplate(
		&buffer,
		"learning",
		viewModel,
	); err != nil {
		http.Error(
			w,
			"Error interno al generar la interfaz",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)

	w.WriteHeader(status)

	_, _ = w.Write(buffer.Bytes())
}

// renderTutorialPage renders one static thematic tutorial from the
// UConversor Learning area.
//
// Responsibility:
//   - Execute the common tutorial template.
//   - Write the generated HTML response.
//
// Receives:
//   - HTTP response writer.
//   - HTTP status code.
//   - TutorialViewModel previously validated by the Web handler.
//
// Produces:
//   - Complete HTML response for a Learning tutorial.
//
// Previous logical stage:
//   - handleTutorial() in handlers.go.
//
// Next logical stage:
//   - tutorial.html and the selected thematic tutorial template.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not call App.Convert().
//   - Must not inspect Registry or Catalog.
//   - Must not decide which tutorial is valid.
//   - Tutorial validation belongs to handleTutorial().
func (h *Handler) renderTutorialPage(
	w http.ResponseWriter,
	status int,
	viewModel TutorialViewModel,
) {
	var buffer bytes.Buffer

	if err := h.templates.ExecuteTemplate(
		&buffer,
		"tutorial",
		viewModel,
	); err != nil {
		http.Error(
			w,
			"Error interno al generar la interfaz",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)

	w.WriteHeader(status)
	_, _ = w.Write(buffer.Bytes())
}
