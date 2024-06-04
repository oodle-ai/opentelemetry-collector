// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pcommon

import (
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
)

// Float64Slice represents a []float64 slice.
// The instance of Float64Slice can be assigned to multiple objects since it's immutable.
//
// Must use NewFloat64Slice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Float64Slice internal.Float64Slice

func (ms Float64Slice) getOrig() *[]float64 {
	return internal.GetOrigFloat64Slice(internal.Float64Slice(ms))
}

func (ms Float64Slice) getState() *internal.State {
	return internal.GetFloat64SliceState(internal.Float64Slice(ms))
}

// NewFloat64Slice creates a new empty Float64Slice.
func NewFloat64Slice() Float64Slice {
	orig := []float64(nil)
	state := internal.StateMutable
	return Float64Slice(internal.NewFloat64Slice(&orig, &state))
}

// AsRaw returns a copy of the []float64 slice.
func (ms Float64Slice) AsRaw() []float64 {
	return copyFloat64Slice(nil, *ms.getOrig())
}

// FromRaw copies raw []float64 into the slice Float64Slice.
func (ms Float64Slice) FromRaw(val []float64) {
	ms.getState().AssertMutable()
	*ms.getOrig() = copyFloat64Slice(*ms.getOrig(), val)
}

// Len returns length of the []float64 slice value.
// Equivalent of len(float64Slice).
func (ms Float64Slice) Len() int {
	return len(*ms.getOrig())
}

// At returns an item from particular index.
// Equivalent of float64Slice[i].
func (ms Float64Slice) At(i int) float64 {
	return (*ms.getOrig())[i]
}

// SetAt sets float64 item at particular index.
// Equivalent of float64Slice[i] = val
func (ms Float64Slice) SetAt(i int, val float64) {
	ms.getState().AssertMutable()
	(*ms.getOrig())[i] = val
}

// EnsureCapacity ensures Float64Slice has at least the specified capacity.
//  1. If the newCap <= cap, then is no change in capacity.
//  2. If the newCap > cap, then the slice capacity will be expanded to the provided value which will be equivalent of:
//     buf := make([]float64, len(float64Slice), newCap)
//     copy(buf, float64Slice)
//     float64Slice = buf
func (ms Float64Slice) EnsureCapacity(newCap int) {
	ms.getState().AssertMutable()
	oldCap := cap(*ms.getOrig())
	if newCap <= oldCap {
		return
	}

	newOrig := make([]float64, len(*ms.getOrig()), newCap)
	copy(newOrig, *ms.getOrig())
	*ms.getOrig() = newOrig
}

// Append appends extra elements to Float64Slice.
// Equivalent of float64Slice = append(float64Slice, elms...)
func (ms Float64Slice) Append(elms ...float64) {
	ms.getState().AssertMutable()
	*ms.getOrig() = append(*ms.getOrig(), elms...)
}

// MoveTo moves all elements from the current slice overriding the destination and
// resetting the current instance to its zero value.
func (ms Float64Slice) MoveTo(dest Float64Slice) {
	ms.getState().AssertMutable()
	dest.getState().AssertMutable()
	*dest.getOrig() = *ms.getOrig()
	*ms.getOrig() = nil
}

// CopyTo copies all elements from the current slice overriding the destination.
func (ms Float64Slice) CopyTo(dest Float64Slice) {
	dest.getState().AssertMutable()
	*dest.getOrig() = copyFloat64Slice(*dest.getOrig(), *ms.getOrig())
}

func copyFloat64Slice(dst, src []float64) []float64 {
	dst = dst[:0]
	return append(dst, src...)
}
