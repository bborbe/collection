// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// ContainsAny returns true if at least one element from slice b is present in slice a.
// It checks whether a and b have any intersection (a ∩ b ≠ ∅).
// Empty slice b always returns false (no elements to check).
func ContainsAny[T comparable](a []T, b []T) bool {
	if len(b) == 0 {
		return false
	}
	mapA := make(map[T]bool)
	for _, aa := range a {
		mapA[aa] = true
	}
	for _, bb := range b {
		if mapA[bb] {
			return true
		}
	}
	return false
}
