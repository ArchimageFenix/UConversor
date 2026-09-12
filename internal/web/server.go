// File: server.go
//
// Responsibility:
//   - Configure and run the HTTP server.
//   - Register application routes.
//   - Serve embedded static resources.
//
// Receives:
//   - Server address.
//   - Configured Handler.
//   - Embedded static resources.
//
// Produces:
//   - Configured HTTP server ready to accept requests.
//
// Previous logical stage:
//   - cmd/web/main.go.
//
// Next logical stage:
//   - handlers.go for application requests.
//   - Embedded static resources for /static/ requests.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not contain scientific logic.
//   - Must not recognize units.
//   - Must not prepare presentation view models.
//   - Must not classify application errors.
//   - Request processing belongs to handlers.go.
package web

import (
	"fmt"
	"io/fs"
	"net/http"
)

// Server contains the HTTP infrastructure used by
// the UConversor web interface.
//
// It is responsible only for HTTP server configuration,
// route registration and server execution.
type Server struct {
	httpServer *http.Server
}

// NewServer creates the UConversor HTTP server.
//
// Receives:
//   - Address where the server will listen.
//   - Handler containing the web application request logic.
//
// Produces:
//   - Configured Server.
//
// Previous logical stage:
//   - cmd/web/main.go creates the application's dependencies.
//
// Next logical stage:
//   - registerRoutes().
//
// Important restrictions:
//   - Does not create app.App.
//   - Does not load HTML templates.
//   - Does not execute conversions.
//   - Does not contain presentation logic.
func NewServer(
	address string,
	handler *Handler,
) (*Server, error) {
	staticFS, err := fs.Sub(
		assets,
		"static",
	)
	if err != nil {
		return nil, fmt.Errorf(
			"cargar recursos estáticos: %w",
			err,
		)
	}

	server := &Server{}

	mux := http.NewServeMux()

	server.registerRoutes(
		mux,
		staticFS,
		handler,
	)

	server.httpServer = &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return server, nil
}

// registerRoutes defines the HTTP routes exposed by UConversor.
//
// Routes:
//
//	GET  /
//	     Displays the initial web interface.
//
//	GET  /examples
//	     Displays the units currently supported by UConversor.
//
//	POST /convert
//	     Processes a conversion request.
//
//	/static/
//	     Serves embedded CSS and other static resources.
//
// Previous logical stage:
//   - NewServer().
//
// Next logical stage:
//   - handlers.go for "/", "/examples" and "/convert".
//   - Embedded filesystem for "/static/".
//
// Important restrictions:
//   - Routes only requests.
//   - Must not execute application logic.
//   - Must not perform conversion validation.
//   - Must not render HTML directly.
func (s *Server) registerRoutes(
	mux *http.ServeMux,
	staticFS fs.FS,
	handler *Handler,
) {
	mux.HandleFunc(
		"/",
		handler.handleHome,
	)

	mux.HandleFunc(
		"/examples",
		handler.handleExamples,
	)

	mux.HandleFunc(
		"/convert",
		handler.handleConvert,
	)

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.FS(staticFS),
			),
		),
	)
}

// ListenAndServe starts the configured HTTP server.
//
// Produces:
//   - An error when the HTTP server stops or cannot start.
//
// Previous logical stage:
//   - cmd/web/main.go.
//
// Next logical stage:
//   - HTTP request lifecycle.
//
// Important restrictions:
//   - Does not interpret server errors.
//   - Startup/shutdown policy belongs to the composition root.
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
