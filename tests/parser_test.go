// File: parser_test.go
// Responsibility: Verify normalization and parsing independently of conversion.
// Receives: Representative user input strings.
// Produces: Automated pass/fail evidence for parser behavior.
// Previous logical stage: input package implementation.
// Next logical stage: conversion integration tests.
// Important restrictions: Does not test unit recognition or formulas.
package tests

import (
	"testing"
	"unit-converter/internal/input"
)

func TestParseCompactAndSpaced(t *testing.T) {
	for _, raw := range []string{"24mph", "24 mph", "2,5km"} {
		p, err := input.Parse(input.Normalize(raw))
		if err != nil {
			t.Fatalf("%q: %v", raw, err)
		}
		if p.Unit == "" {
			t.Fatalf("%q: unit empty", raw)
		}
	}
}
