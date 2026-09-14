// File: main.go
//
// Responsibility:
//   - Start the UConversor web frontend.
//   - Compose the application, templates, handlers and HTTP server.
//   - Resolve the listening port from the deployment environment.
//   - Load the Web security configuration.
//   - Transfer execution control to the web infrastructure.
//
// Receives:
//   - PORT environment variable when supplied by the deployment platform.
//   - config/security.yaml when available.
//   - No conversion data directly.
//
// Produces:
//   - Running UConversor HTTP service.
//
// Previous logical stage:
//   - Operating system or deployment platform starts the web executable.
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
	"os"

	"unit-converter/internal/app"
	"unit-converter/internal/catalog"
	"unit-converter/internal/units"

	uweb "unit-converter/internal/web"
)

const (
	defaultWebPort     = "8080"
	securityConfigPath = "config/security.yaml"
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

	configResult, err := uweb.LoadSecurityConfig(
		securityConfigPath,
	)
	if err != nil {
		log.Fatal(err)
	}

	if configResult.UsedDefaults {
		log.Printf(
			"ADVERTENCIA: %s no encontrado; usando configuración de seguridad predeterminada",
			securityConfigPath,
		)
	}

	handler := uweb.NewHandler(
		templates,
		application,
		configResult.Config,
	)
	address := resolveAddress()

	server, err := uweb.NewServer(
		address,
		handler,
		configResult.Config,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"UConversor Web escuchando en %s",
		address,
	)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// resolveAddress determines the HTTP listening address.
//
// Receives:
//   - PORT environment variable.
//
// Produces:
//   - Address in the form ":<port>".
//
// Previous logical stage:
//   - Deployment environment or local operating system.
//
// Next logical stage:
//   - NewServer() in internal/web/server.go.
//
// Important restrictions:
//   - PORT is controlled by the execution environment.
//   - If PORT is absent, local development must continue using port 8080.
//   - Must not contain HTTP or scientific processing logic.
func resolveAddress() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultWebPort
	}

	return ":" + port
}
