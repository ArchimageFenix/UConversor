// File: registry.go
// Responsibility: Index supported units and perform early recognition/rejection.
// Receives: Unit definitions from catalog packages and normalized unit tokens from input.
// Produces: A recognized Unit or an unavailable result.
// Previous logical stage: input/normalizer.go and catalog files.
// Next logical stage: conversion/validator.go.
// Important restrictions: Unknown units stop here and must never reach scientific validation or conversion.
package units

import (
	"strings"
)

type Registry struct {
	byToken map[string]Unit
	all     []Unit
}

func NewRegistry(defs []Unit) *Registry {
	r := &Registry{byToken: make(map[string]Unit), all: append([]Unit(nil), defs...)}
	for _, u := range defs {
		r.byToken[u.Symbol] = u
		for _, alias := range u.Aliases {
			r.byToken[alias] = u
		}
	}
	return r
}

func (r *Registry) Find(token string) (Unit, bool) {
	if u, ok := r.byToken[token]; ok {
		return u, true
	}
	// Conservative fallback only for aliases that are explicitly lowercase ASCII-like.
	lower := strings.ToLower(token)
	if u, ok := r.byToken[lower]; ok {
		return u, true
	}
	return Unit{}, false
}

func (r *Registry) UnitsForMagnitude(m Magnitude) []Unit {
	out := make([]Unit, 0)
	for _, u := range r.all {
		if u.Magnitude == m {
			out = append(out, u)
		}
	}
	return out
}

// AllUnits returns a copy of every unit currently registered.
//
// The returned slice is intentionally detached from the registry's
// internal storage so callers cannot modify the registry state.
func (r *Registry) AllUnits() []Unit {
	return append([]Unit(nil), r.all...)
}
