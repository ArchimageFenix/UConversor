// File: source.go
// Responsibility: Represent the scientific or normative provenance of a conversion definition.
// Receives: Source metadata declared by unit catalogs.
// Produces: Source values attached to unit definitions and formulas.
// Previous logical stage: catalog definitions.
// Next logical stage: conversion/validator.go.
// Important restrictions: Sources are metadata only; this file performs no conversion or I/O.
package model

type SourceKind string

const (
	SourceOfficial SourceKind = "official"
	SourceExact    SourceKind = "exact"
	SourceDerived  SourceKind = "derived"
	SourceCustom   SourceKind = "custom"
)

type Source struct {
	Kind      SourceKind
	Authority string
	Reference string
	URL       string
	Notes     string
}
