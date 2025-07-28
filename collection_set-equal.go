// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import (
	"sync"
)

// HasEqual represents types that can compare themselves for equality with another value.
type HasEqual[V any] interface {
	Equal(value V) bool
}

// SetEqual represents a thread-safe set for types that implement HasEqual.
// Elements are uniquely identified by their Equal method.
type SetEqual[T HasEqual[T]] interface {
	// Add inserts an element into the set, using its Equal method for uniqueness.
	Add(element T)
	// Remove deletes an element from the set using its Equal method for matching.
	Remove(element T)
	// Contains reports whether an element matching the given value is present in the set.
	Contains(element T) bool
	// Slice returns all elements as a slice in arbitrary order.
	Slice() []T
	// Length returns the number of elements in the set.
	Length() int
}

// NewSetEqual creates a new thread-safe set for types that implement HasEqual.
func NewSetEqual[T HasEqual[T]]() SetEqual[T] {
	return &setEqual[T]{
		data: make([]T, 0),
	}
}

type setEqual[T HasEqual[T]] struct {
	mux  sync.Mutex
	data []T
}

func (s *setEqual[T]) Add(element T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.contains(element) {
		return
	}
	s.data = append(s.data, element)
}

func (s *setEqual[T]) Remove(element T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]T, 0, len(s.data))
	for _, e := range s.data {
		if e.Equal(element) {
			continue
		}
		result = append(result, e)
	}
	s.data = result
}

func (s *setEqual[T]) Contains(element T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.contains(element)
}

func (s *setEqual[T]) contains(element T) bool {
	for _, e := range s.data {
		if e.Equal(element) {
			return true
		}
	}
	return false
}

func (s *setEqual[T]) Slice() []T {
	s.mux.Lock()
	defer s.mux.Unlock()
	return Copy(s.data)
}

func (s *setEqual[T]) Length() int {
	s.mux.Lock()
	defer s.mux.Unlock()

	return len(s.data)
}
