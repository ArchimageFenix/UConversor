// File: main.go
// Responsibility: Provide the initial CLI boundary and delegate work to internal/app.
// Receives: One user expression from command arguments or standard input.
// Produces: Styled terminal result or error message on the console.
// Previous logical stage: User/operating system.
// Next logical stage: internal/app/app.go.
// Important restrictions: main must remain minimal and contain no parsing, formulas, conversion, or scientific validation logic.
package main

import (
	"bufio"
	"os"
	"strings"
	"unit-converter/internal/app"
	"unit-converter/internal/catalog"
	"unit-converter/internal/output"
	"unit-converter/internal/output/help"
	"unit-converter/internal/units"
)

func main() {
	registry := units.NewRegistry(catalog.All())
	application := app.New(registry)
	console := output.NewConsole(os.Stdout)
	console.PrintWelcome()

	raw := strings.TrimSpace(strings.Join(os.Args[1:], " "))
	if raw == "" {
		console.PrintPrompt("Ingrese una medida (ej. 24mph, 5km, 2200Ω): ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			raw = scanner.Text()
		}
	}
	if strings.EqualFold(strings.TrimSpace(raw), "h") {
		viewer := help.NewViewer(os.Stdin, os.Stdout)
		viewer.Show()
		return
	}
	result, err := application.Convert(raw)
	if err != nil {
		console.PrintError(err)
		return
	}

	console.PrintResult(result)
}
