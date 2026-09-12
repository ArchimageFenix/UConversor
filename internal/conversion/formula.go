// File: formula.go
// Responsibility: Document the generic conversion strategy used by Unit transforms.
// Receives: Unit transformation metadata.
// Produces: Mathematical conversion through a magnitude reference unit.
// Previous logical stage: catalog definitions.
// Next logical stage: conversion/engine.go.
// Important restrictions: Current generic model supports linear and affine transforms; specialized formulas may be added later.
package conversion

// The generic transformation is intentionally represented in units.Unit:
// reference = value*scale + offset
// value     = (reference-offset)/scale
// This supports ordinary scale conversions and affine temperature conversions.
