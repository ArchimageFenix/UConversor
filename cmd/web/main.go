// File: main.go
//
// Responsibility:
//   - Start the UConversor web frontend.
//   - Compose the application, templates, handlers and HTTP server.
//   - Transfer execution control to the web infrastructure.
//
// Receives:
//   - No conversion data directly.
//
// Produces:
//   - Running UConversor HTTP service.
//
// Previous logical stage:
//   - Operating system starts the web executable.
//
// Next logical stage:
//   - internal/web/server.go.
//
// Important restrictions:
//   - Must remain minimal.
//   - Must not contain HTTP handlers.
//   - Must not contain parsing logic.
//   - Must not contain scientific conversion logic.
//   - Must only compose and connect application dependencies.
package main

import (
	"log"

	"unit-converter/internal/app"
	"unit-converter/internal/catalog"
	"unit-converter/internal/units"

	uweb "unit-converter/internal/web"
)

func main() {
	registry := units.NewRegistry(
		catalog.All(),
	)

	application := app.New(
		registry,
	)

	templates, err := uweb.LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	handler := uweb.NewHandler(
		templates,
		application,
	)

	server, err := uweb.NewServer(
		":8080",
		handler,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"UConversor Web disponible en http://localhost:8080",
	)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
