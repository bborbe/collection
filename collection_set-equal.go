// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import (
	"fmt"
	"strings"
	"sync"
)

// HasEqual represents types that can compare themselves for equality with another value.
type HasEqual[V any] interface {
	Equal(value V) bool
}

// SetEqual represents a thread-safe set for types that implement HasEqual.
// Elements are uniquely identified by their Equal method.
//
// Performance: This implementation uses a slice-based approach. Operations have
// O(n) complexity where n is the number of elements. For large sets or performance-critical
// code, consider using SetHashCode which provides O(1) average-case operations.
type SetEqual[T HasEqual[T]] interface {
	// Add inserts elements into the set, using the Equal method for uniqueness checking.
	// Duplicate elements are automatically ignored.
	// Multiple elements can be added in a single call with only one mutex lock.
	Add(elements ...T)
	// Remove deletes an element from the set using its Equal method for matching.
	Remove(element T)
	// Contains reports whether an element matching the given value is present in the set.
	Contains(element T) bool
	// Slice returns all elements as a slice in arbitrary order.
	Slice() []T
	// Length returns the number of elements in the set.
	Length() int
	// String returns a human-readable string representation of the set.
	String() string
}

// NewSetEqual creates a new thread-safe set for types that implement HasEqual.
// It accepts optional initial elements to populate the set.
// Duplicate elements are automatically handled using the Equal method.
//
// Performance: This implementation uses a slice-based approach with O(n) operations.
// Initialization with k elements has O(k²) complexity due to uniqueness checks.
// For better performance with large sets, use NewSetHashCode instead.
//
// Example:
//
//	type User struct { ID int; Name string }
//	func (u User) Equal(other User) bool { return u.ID == other.ID }
//	set := collection.NewSetEqual(User{1, "Alice"}, User{2, "Bob"})
func NewSetEqual[T HasEqual[T]](elements ...T) SetEqual[T] {
	s := &setEqual[T]{
		data: make([]T, 0),
	}
	s.Add(elements...)
	return s
}

type setEqual[T HasEqual[T]] struct {
	mux  sync.Mutex
	data []T
}

func (s *setEqual[T]) Add(elements ...T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, element := range elements {
		if s.contains(element) {
			continue
		}
		s.data = append(s.data, element)
	}
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

// String returns a human-readable string representation of the set.
// Format: "SetEqual[element1, element2, ...]" for non-empty sets, "SetEqual[]" for empty sets.
// Elements appear in the order they were added to the set.
func (s *setEqual[T]) String() string {
	s.mux.Lock()
	defer s.mux.Unlock()

	if len(s.data) == 0 {
		return "SetEqual[]"
	}

	var b strings.Builder
	b.WriteString("SetEqual[")
	for i, e := range s.data {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%v", e)
	}
	b.WriteString("]")
	return b.String()
}
