// File: normalizer.go
// Responsibility: Normalize harmless input variations before parsing.
// Receives: Raw user text.
// Produces: Trimmed input with decimal comma normalized when unambiguous.
// Previous logical stage: CLI input.
// Next logical stage: input/parser.go.
// Important restrictions: Must not alter case-sensitive scientific symbols or invent units.
package input

import "strings"

func Normalize(raw string) string {
	s := strings.TrimSpace(raw)
	// Spanish decimal comma is accepted when no decimal point is already present.
	if strings.Count(s, ",") == 1 && !strings.Contains(s, ".") {
		s = strings.Replace(s, ",", ".", 1)
	}
	return s
}
