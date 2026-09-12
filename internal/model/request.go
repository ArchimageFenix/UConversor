// File: request.go
// Responsibility: Carry a parsed conversion request through the application core.
// Receives: Numeric value and recognized unit symbol.
// Produces: A typed request consumed by validation and conversion.
// Previous logical stage: input/parser.go and units/registry.go.
// Next logical stage: conversion/validator.go.
// Important restrictions: Contains no parsing, formatting, or conversion logic.
package model

type ConversionRequest struct {
	Value      float64
	InputUnit  string
	UnitSymbol string
	Magnitude  string
	Family     string
}
