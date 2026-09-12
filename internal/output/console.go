// File: console.go
// Responsibility: Present structured conversion results and application messages on the CLI with terminal styling.
// Receives: model.ConversionResult or user-facing error text.
// Produces: Ordered terminal output through an injected io.Writer.
// Previous logical stage: output/formatter.go or app-level errors.
// Next logical stage: End of current CLI flow.
// Important restrictions: Performs no parsing, scientific validation, or conversion; ANSI styling belongs only to this CLI presenter.
package output

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unit-converter/internal/model"
)

const (
	ansiReset   = "\x1b[0m"
	ansiBold    = "\x1b[1m"
	ansiRed     = "\x1b[31m"
	ansiGreen   = "\x1b[32m"
	ansiYellow  = "\x1b[33m"
	ansiBlue    = "\x1b[34m"
	ansiMagenta = "\x1b[35m"
	ansiCyan    = "\x1b[36m"
	ansiGray    = "\x1b[90m"
)

type Console struct {
	W     io.Writer
	Color bool
}

// NewConsole creates the CLI presenter. Set the NO_COLOR environment variable
// to disable ANSI sequences when output is redirected or plain text is desired.
func NewConsole(w io.Writer) Console {
	return Console{
		W:     w,
		Color: os.Getenv("NO_COLOR") == "" && strings.ToLower(os.Getenv("TERM")) != "dumb",
	}
}

func (c Console) Println(text string) {
	fmt.Fprintln(c.W, text)
}

func (c Console) PrintPrompt(text string) {
	fmt.Fprint(c.W, c.paint(ansiCyan+ansiBold, text))
}

func (c Console) PrintError(err error) {
	fmt.Fprintln(c.W, c.paint(ansiRed+ansiBold, "ERROR:"), c.paint(ansiRed, err.Error()))
}
func (c *Console) PrintWelcome() {
	fmt.Fprintln(c.W)

	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "╔══════════════════════════════════════════════════════════════╗"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "║                                                              ║"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "║   _   _  ____                                               ║"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "║  | | | |/ ___|___  _ ____   _____ _ __ ___  ___  _ __      ║"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "║  | | | | |   / _ \\| '_ \\ \\ / / _ \\ '__/ __|/ _ \\| '__|     ║"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "║  | |_| | |__| (_) | | | \\ V /  __/ |  \\__ \\ (_) | |        ║"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "║   \\___/ \\____\\___/|_| |_|\\_/ \\___|_|  |___/\\___/|_|        ║"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "║                                                              ║"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "║              CONVERSOR AUTOMÁTICO DE UNIDADES                ║"))
	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "╠══════════════════════════════════════════════════════════════╣"))

	fmt.Fprintln(c.W, "║ UConversor permite introducir un valor junto con su unidad.  ║")
	fmt.Fprintln(c.W, "║ El programa detecta automáticamente la magnitud y genera     ║")
	fmt.Fprintln(c.W, "║ todas las conversiones compatibles disponibles.              ║")
	fmt.Fprintln(c.W, "║                                                              ║")

	fmt.Fprintln(
		c.W,
		"║ Ejemplos: "+c.paint(ansiYellow, "24mph   5km   2200Ω   32°F   1ly")+"                  ║",
	)

	fmt.Fprintln(c.W, c.paint(ansiCyan+ansiBold, "╚══════════════════════════════════════════════════════════════╝"))
	fmt.Fprintln(c.W)
}

func (c Console) PrintResult(result model.ConversionResult) {
	fmt.Fprintln(c.W)
	fmt.Fprintln(c.W, c.paint(ansiGreen+ansiBold, "✓ CONVERSIÓN COMPLETADA"))
	fmt.Fprintf(c.W, "%s %s\n", c.paint(ansiMagenta+ansiBold, "Familia :"), result.Family)
	fmt.Fprintf(c.W, "%s %s\n", c.paint(ansiBlue+ansiBold, "Magnitud:"), result.Magnitude)
	fmt.Fprintf(c.W, "%s %.4f %s\n\n", c.paint(ansiYellow+ansiBold, "Entrada :"), result.Input.Value, result.Input.Symbol)

	c.printTable(result)
}

func (c Console) printTable(result model.ConversionResult) {
	border := fmt.Sprintf(
		"+-%s-+-%s-+-%s-+",
		strings.Repeat("-", unitColumnWidth),
		strings.Repeat("-", symbolColumnWidth),
		strings.Repeat("-", valueColumnWidth),
	)

	fmt.Fprintln(c.W, c.paint(ansiGray, border))
	fmt.Fprintf(
		c.W,
		"| %-*s | %-*s | %*s |\n",
		unitColumnWidth, c.paint(ansiCyan+ansiBold, "UNIDAD"),
		symbolColumnWidth, c.paint(ansiCyan+ansiBold, "SÍMBOLO"),
		valueColumnWidth, c.paint(ansiCyan+ansiBold, "VALOR"),
	)
	fmt.Fprintln(c.W, c.paint(ansiGray, border))

	for _, v := range result.Values {
		// Pad before applying ANSI styles so escape sequences do not affect alignment.
		name := fmt.Sprintf("%-*s", unitColumnWidth, v.Name)
		symbol := fmt.Sprintf("%-*s", symbolColumnWidth, v.Symbol)
		value := fmt.Sprintf("%*s", valueColumnWidth, formatValue(v.Value))
		fmt.Fprintf(
			c.W,
			"| %s | %s | %s |\n",
			name,
			c.paint(ansiYellow, symbol),
			c.paint(ansiGreen, value),
		)
	}

	fmt.Fprintln(c.W, c.paint(ansiGray, border))
}

func (c Console) paint(code, text string) string {
	if !c.Color {
		return text
	}
	return code + text + ansiReset
}
