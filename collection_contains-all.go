// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// ContainsAll returns true if all elements from slice b are present in slice a.
// It checks whether a is a superset of b (a ⊇ b).
// Empty slice b always returns true. Duplicate elements in b are treated as single elements.
func ContainsAll[T comparable](a []T, b []T) bool {
	mapA := make(map[T]bool)
	for _, aa := range a {
		mapA[aa] = true
	}
	for _, bb := range b {
		if mapA[bb] == false {
			return false
		}
	}
	return true
}
