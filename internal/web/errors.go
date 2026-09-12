// File: errors.go
//
// Responsibility:
//   - Classify application/domain errors for the web presentation layer.
//   - Translate internal errors into safe, user-facing messages.
//
// Receives:
//   - Errors returned by app.Convert() and lower application layers.
//
// Produces:
//   - HTTP status code for internal protocol handling.
//   - ErrorViewModel with a controlled message for the browser.
//
// Previous logical stage:
//   - handlers.go receives an error from app.Convert().
//
// Next logical stage:
//   - viewmodel.go / render.go.
//
// Important restrictions:
//   - Must not expose internal implementation details.
//   - Must not expose raw error messages to the browser.
//   - Must not expose stack traces, file paths, configuration details,
//     scientific-source failures, or internal HTTP/server information.
//   - Must classify errors with errors.Is(), never by comparing strings.
//   - User-visible messages must remain simple and related to the input domain.
package web

import (
	"errors"
	"net/http"

	"unit-converter/internal/app"
	"unit-converter/internal/conversion"
	"unit-converter/internal/input"
)

// classifyConversionError converts an internal/application error into:
//
//   - an HTTP status code used internally by the server;
//   - a safe ErrorViewModel intended for browser presentation.
//
// The HTTP status is not intended to be shown visually to the user.
func classifyConversionError(err error) (int, ErrorViewModel) {
	switch {
	case errors.Is(err, input.ErrEmptyInput):
		return http.StatusBadRequest, ErrorViewModel{
			Title:   "Entrada requerida",
			Message: "Introduce un valor numérico junto con su unidad.",
		}

	case errors.Is(err, input.ErrMissingValue):
		return http.StatusBadRequest, ErrorViewModel{
			Title:   "Falta el valor",
			Message: "Debes indicar un valor numérico antes de la unidad.",
		}

	case errors.Is(err, input.ErrMissingUnit):
		return http.StatusBadRequest, ErrorViewModel{
			Title:   "Falta la unidad",
			Message: "Debes indicar una unidad para realizar la conversión.",
		}

	case errors.Is(err, input.ErrInvalidValue):
		return http.StatusBadRequest, ErrorViewModel{
			Title:   "Valor no válido",
			Message: "El valor numérico introducido no tiene un formato válido.",
		}

	case errors.Is(err, app.ErrUnavailable):
		return http.StatusBadRequest, ErrorViewModel{
			Title:   "Unidad no reconocida",
			Message: "La unidad introducida no está disponible en UConversor.",
		}

	case errors.Is(err, conversion.ErrOutOfRange):
		return http.StatusBadRequest, ErrorViewModel{
			Title:   "Valor fuera de rango",
			Message: "El valor introducido no es válido para esa unidad.",
		}

	case errors.Is(err, conversion.ErrMissingSource):
		return http.StatusInternalServerError, ErrorViewModel{
			Title:   "No fue posible completar la conversión",
			Message: "Inténtalo nuevamente más tarde.",
		}

	default:
		return http.StatusInternalServerError, ErrorViewModel{
			Title:   "No fue posible completar la conversión",
			Message: "Inténtalo nuevamente más tarde.",
		}
	}
}
