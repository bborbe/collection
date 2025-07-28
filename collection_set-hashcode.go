// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import (
	"sync"
)

// HasHashCode represents types that can provide a string hash code for themselves.
type HasHashCode interface {
	HashCode() string
}

// SetHashCode represents a thread-safe set for types that implement HasHashCode.
// Elements are uniquely identified by their hash code.
type SetHashCode[T HasHashCode] interface {
	// Add inserts an element into the set, using its hash code for uniqueness.
	Add(element T)
	// Remove deletes an element from the set by its hash code.
	Remove(element T)
	// Contains reports whether an element with the given hash code is present in the set.
	Contains(element T) bool
	// Slice returns all elements as a slice in arbitrary order.
	Slice() []T
	// Length returns the number of elements in the set.
	Length() int
}

// NewSetHashCode creates a new thread-safe set for types that implement HasHashCode.
func NewSetHashCode[T HasHashCode]() SetHashCode[T] {
	return &setHashCode[T]{
		data: make(map[string]T),
	}
}

type setHashCode[T HasHashCode] struct {
	mux  sync.Mutex
	data map[string]T
}

func (s *setHashCode[T]) Add(element T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.data[element.HashCode()] = element
}

func (s *setHashCode[T]) Remove(element T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.data, element.HashCode())
}

func (s *setHashCode[T]) Contains(element T) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, found := s.data[element.HashCode()]
	return found
}

func (s *setHashCode[T]) Slice() []T {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]T, 0, len(s.data))
	for _, v := range s.data {
		result = append(result, v)
	}
	return result
}

func (s *setHashCode[T]) Length() int {
	s.mux.Lock()
	defer s.mux.Unlock()

	return len(s.data)
}
