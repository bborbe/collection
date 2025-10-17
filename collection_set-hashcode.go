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
//
// Performance: This implementation uses a map-based approach with O(1) average-case
// operations for Add, Remove, and Contains. This provides better performance than
// SetEqual for large sets or performance-critical code.
type SetHashCode[T HasHashCode] interface {
	// Add inserts elements into the set, using their hash codes for uniqueness.
	// Duplicate elements (same hash code) are automatically ignored.
	// Multiple elements can be added in a single call with only one mutex lock.
	Add(elements ...T)
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
// It accepts optional initial elements to populate the set.
// Duplicate elements (same hash code) are automatically handled.
//
// Performance: This implementation uses a map-based approach with O(1) average-case
// operations. Initialization is O(n) for n elements, making it suitable for large sets.
//
// Example:
//
//	type User struct { ID int; Name string }
//	func (u User) HashCode() string { return fmt.Sprintf("user-%d", u.ID) }
//	set := collection.NewSetHashCode(User{1, "Alice"}, User{2, "Bob"})
func NewSetHashCode[T HasHashCode](elements ...T) SetHashCode[T] {
	s := &setHashCode[T]{
		data: make(map[string]T),
	}
	s.Add(elements...)
	return s
}

type setHashCode[T HasHashCode] struct {
	mux  sync.Mutex
	data map[string]T
}

func (s *setHashCode[T]) Add(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		s.data[element.HashCode()] = element
	}
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
