package help

// File: viewer.go
// Responsibility: Present and navigate the UConversor help manual in the terminal.
// Receives: Help content from content.go and terminal input/output streams.
// Produces: An interactive, read-only paginated help view.
// Previous logical stage: internal/output/help/content.go.
// Next logical stage: Return control to the CLI.
// Important restrictions:
//   - Must not perform unit conversions.
//   - Must not modify help content.
//   - Must not contain scientific formulas or unit definitions.
//   - Must remain independent from the conversion engine.

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const pageSize = 20

// Viewer represents the interactive terminal help viewer.
type Viewer struct {
	In  io.Reader
	Out io.Writer
}

const (
	helpBackground = "\033[48;5;17m"  // Azul oscuro
	helpForeground = "\033[38;5;252m" // Blanco suave
	helpAccent     = "\033[38;5;117m" // Azul claro
	helpReset      = "\033[0m"
)

// NewViewer creates a help viewer using the provided input and output streams.
func NewViewer(in io.Reader, out io.Writer) *Viewer {
	return &Viewer{
		In:  in,
		Out: out,
	}
}

// Show opens the help manual and keeps control until the user exits.
func (v *Viewer) Show() {
	lines := Lines()

	if len(lines) == 0 {
		return
	}

	page := 0
	scanner := bufio.NewScanner(v.In)

	for {
		totalPages := (len(lines) + pageSize - 1) / pageSize

		if page >= totalPages {
			page = totalPages - 1
		}

		v.printPage(lines, page, totalPages)

		if !scanner.Scan() {
			return
		}

		command := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch command {
		case "", "n":
			if page < totalPages-1 {
				page++
			}

		case "p":
			if page > 0 {
				page--
			}

		case "q":
			return
		}
	}
}

// printPage prints only the portion of the manual belonging to the current page.
func (v *Viewer) printPage(lines []string, page, totalPages int) {
	start := page * pageSize
	end := start + pageSize

	if end > len(lines) {
		end = len(lines)
	}

	fmt.Fprintln(v.Out)
	fmt.Fprintln(
		v.Out,
		v.paintAccent("════════════════ UConversor - Ayuda ════════════════"),
	)

	for _, line := range lines[start:end] {
		fmt.Fprintln(v.Out, v.paint(line))
	}

	fmt.Fprintln(
		v.Out,
		v.paintAccent(
			fmt.Sprintf(
				"──────── Página %d/%d ──────── [Enter/N] Siguiente  [P] Anterior  [Q] Salir",
				page+1,
				totalPages,
			),
		),
	)

	fmt.Fprint(v.Out, v.paint("> "))
}

func (v *Viewer) paint(text string) string {
	return helpBackground + helpForeground + text + helpReset
}

func (v *Viewer) paintAccent(text string) string {
	return helpBackground + helpAccent + text + helpReset
}
