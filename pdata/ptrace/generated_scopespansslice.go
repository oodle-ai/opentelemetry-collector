// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"sort"

	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlptrace "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/trace/v1"
)

// ScopeSpansSlice logically represents a slice of ScopeSpans.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewScopeSpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeSpansSlice struct {
	orig  *[]*otlptrace.ScopeSpans
	state *internal.State
}

func newScopeSpansSlice(orig *[]*otlptrace.ScopeSpans, state *internal.State) ScopeSpansSlice {
	return ScopeSpansSlice{orig: orig, state: state}
}

// NewScopeSpansSlice creates a ScopeSpansSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewScopeSpansSlice() ScopeSpansSlice {
	orig := []*otlptrace.ScopeSpans(nil)
	state := internal.StateMutable
	return newScopeSpansSlice(&orig, &state)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewScopeSpansSlice()".
func (es ScopeSpansSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es ScopeSpansSlice) At(i int) ScopeSpans {
	return newScopeSpans((*es.orig)[i], es.state)
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ScopeSpansSlice can be initialized:
//
//	es := NewScopeSpansSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es ScopeSpansSlice) EnsureCapacity(newCap int) {
	es.state.AssertMutable()
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlptrace.ScopeSpans, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty ScopeSpans.
// It returns the newly added ScopeSpans.
func (es ScopeSpansSlice) AppendEmpty() ScopeSpans {
	es.state.AssertMutable()
	*es.orig = append(*es.orig, &otlptrace.ScopeSpans{})
	return es.At(es.Len() - 1)
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ScopeSpansSlice) MoveAndAppendTo(dest ScopeSpansSlice) {
	es.state.AssertMutable()
	dest.state.AssertMutable()
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ScopeSpansSlice) RemoveIf(f func(ScopeSpans) bool) {
	es.state.AssertMutable()
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		newLen++
	}
	*es.orig = (*es.orig)[:newLen]
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es ScopeSpansSlice) CopyTo(dest ScopeSpansSlice) {
	dest.state.AssertMutable()
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newScopeSpans((*es.orig)[i], es.state).CopyTo(newScopeSpans((*dest.orig)[i], dest.state))
		}
		return
	}
	origs := make([]otlptrace.ScopeSpans, srcLen)
	wrappers := make([]*otlptrace.ScopeSpans, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newScopeSpans((*es.orig)[i], es.state).CopyTo(newScopeSpans(wrappers[i], dest.state))
	}
	*dest.orig = wrappers
}

// Sort sorts the ScopeSpans elements within ScopeSpansSlice given the
// provided less function so that two instances of ScopeSpansSlice
// can be compared.
func (es ScopeSpansSlice) Sort(less func(a, b ScopeSpans) bool) {
	es.state.AssertMutable()
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
}
