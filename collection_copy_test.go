// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Copy", func() {
	It("returns empty slice for empty input", func() {
		result := collection.Copy([]string{})
		Expect(result).To(BeEmpty())
		Expect(result).ToNot(BeNil())
	})

	It("creates a copy of string slice", func() {
		original := []string{"a", "b", "c"}
		result := collection.Copy(original)

		Expect(result).To(Equal(original))
		// Verify they are different slices by modifying one

		// Modify original to verify independence
		original[0] = "modified"
		Expect(result[0]).To(Equal("a")) // Copy should be unchanged
	})

	It("creates a copy of integer slice", func() {
		original := []int{1, 2, 3, 4, 5}
		result := collection.Copy(original)

		Expect(result).To(Equal(original))
		// Verify they are different slices by modifying one

		// Modify original to verify independence
		original[0] = 999
		Expect(result[0]).To(Equal(1)) // Copy should be unchanged
	})

	It("creates a copy of custom type slice", func() {
		original := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
		}
		result := collection.Copy(original)

		Expect(result).To(Equal(original))
		// Verify they are different slices by modifying one

		// Modify original to verify independence
		original[0].Firstname = "Modified"
		Expect(result[0].Firstname).To(Equal("Alice")) // Copy should be unchanged
	})

	It("handles single element slice", func() {
		original := []string{"single"}
		result := collection.Copy(original)

		Expect(result).To(Equal(original))
		Expect(result).To(HaveLen(1))
		Expect(result[0]).To(Equal("single"))
	})

	It("preserves slice capacity behavior", func() {
		original := []int{1, 2, 3}
		result := collection.Copy(original)

		// Both should be able to grow independently
		original = append(original, 4)
		result = append(result, 5)

		Expect(original).To(Equal([]int{1, 2, 3, 4}))
		Expect(result).To(Equal([]int{1, 2, 3, 5}))
	})
})
