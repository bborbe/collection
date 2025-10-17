// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import (
	"fmt"
	"strings"
	"sync"
)

// Set represents a thread-safe collection of unique elements.
// This implementation uses a map for O(1) average-case lookups, additions, and deletions.
type Set[T comparable] interface {
	// Add inserts elements into the set. Duplicate elements are automatically ignored.
	// Multiple elements can be added in a single call with only one mutex lock.
	Add(elements ...T)
	// Remove deletes an element from the set.
	Remove(element T)
	// Contains reports whether an element is present in the set.
	Contains(element T) bool
	// Slice returns all elements as a slice in arbitrary order.
	Slice() []T
	// Length returns the number of elements in the set.
	Length() int
	// String returns a human-readable string representation of the set.
	String() string
}

// NewSet creates a new thread-safe set for comparable types.
// It accepts optional initial elements to populate the set.
// Duplicate elements are automatically handled.
//
// Performance: This implementation uses a map-based approach with O(1) average-case
// operations for Add, Remove, and Contains. Initialization is O(n) for n elements.
//
// Example:
//
//	set := collection.NewSet(1, 2, 3)
//	empty := collection.NewSet[int]()
func NewSet[T comparable](elements ...T) Set[T] {
	s := &set[T]{
		data: make(map[T]struct{}),
	}
	s.Add(elements...)
	return s
}

type set[T comparable] struct {
	mux  sync.Mutex
	data map[T]struct{}
}

func (s *set[T]) Add(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()
	for _, element := range elements {
		s.data[element] = struct{}{}
	}
}

func (s *set[T]) Remove(element T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.data, element)
}

func (s *set[T]) Contains(element T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, found := s.data[element]
	return found
}

func (s *set[T]) Slice() []T {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]T, 0, len(s.data))
	for k := range s.data {
		result = append(result, k)
	}
	return result
}

func (s *set[T]) Length() int {
	s.mux.Lock()
	defer s.mux.Unlock()

	return len(s.data)
}

// String returns a human-readable string representation of the set.
// Format: "Set[element1, element2, ...]" for non-empty sets, "Set[]" for empty sets.
// Note: Element order is non-deterministic due to map iteration and may vary between calls.
func (s *set[T]) String() string {
	s.mux.Lock()
	defer s.mux.Unlock()

	if len(s.data) == 0 {
		return "Set[]"
	}

	var b strings.Builder
	b.WriteString("Set[")
	first := true
	for k := range s.data {
		if !first {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%v", k)
		first = false
	}
	b.WriteString("]")
	return b.String()
}
