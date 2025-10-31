// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = DescribeTable(
	"Intersect",
	func(a []string, b []string, expected []string) {
		result := collection.Intersect(a, b)
		Expect(result).To(Equal(expected))
	},
	Entry("both empty", []string{}, []string{}, []string{}),
	Entry("first empty", []string{}, []string{"a", "b"}, []string{}),
	Entry("second empty", []string{"a", "b"}, []string{}, []string{}),
	Entry("no intersection", []string{"a", "b"}, []string{"c", "d"}, []string{}),
	Entry(
		"complete intersection",
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"},
	),
	Entry(
		"partial intersection",
		[]string{"a", "b", "c"},
		[]string{"b", "c", "d"},
		[]string{"b", "c"},
	),
	Entry("single common element", []string{"a", "b", "c"}, []string{"b", "d", "e"}, []string{"b"}),
	Entry(
		"preserves order from first slice",
		[]string{"c", "a", "b"},
		[]string{"a", "b", "c"},
		[]string{"c", "a", "b"},
	),
	Entry(
		"handles duplicates in first slice",
		[]string{"a", "b", "a", "c"},
		[]string{"a", "c"},
		[]string{"a", "c"},
	),
	Entry(
		"handles duplicates in second slice",
		[]string{"a", "b"},
		[]string{"a", "a", "b", "b"},
		[]string{"a", "b"},
	),
	Entry(
		"handles duplicates in both slices",
		[]string{"a", "a", "b", "b"},
		[]string{"a", "a", "c"},
		[]string{"a"},
	),
)

var _ = DescribeTable("Intersect with integers",
	func(a []int, b []int, expected []int) {
		result := collection.Intersect(a, b)
		Expect(result).To(Equal(expected))
	},
	Entry("both empty", []int{}, []int{}, []int{}),
	Entry("no intersection", []int{1, 2, 3}, []int{4, 5, 6}, []int{}),
	Entry("complete intersection", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}),
	Entry("partial intersection", []int{1, 2, 3, 4}, []int{3, 4, 5, 6}, []int{3, 4}),
	Entry("single element", []int{42}, []int{42, 100}, []int{42}),
	Entry("preserves order", []int{5, 3, 1}, []int{1, 2, 3, 4, 5}, []int{5, 3, 1}),
)

var _ = Describe("Intersect with custom types", func() {
	It("finds common users", func() {
		a := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Charlie", Age: 35},
		}
		b := []User{
			{Firstname: "Bob", Age: 30},
			{Firstname: "David", Age: 40},
			{Firstname: "Alice", Age: 25},
		}
		result := collection.Intersect(a, b)

		expected := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
		}
		Expect(result).To(Equal(expected))
	})

	It("returns empty when no common users", func() {
		a := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
		}
		b := []User{
			{Firstname: "Charlie", Age: 35},
			{Firstname: "David", Age: 40},
		}
		result := collection.Intersect(a, b)

		Expect(result).To(BeEmpty())
	})

	It("preserves order from first slice", func() {
		a := []User{
			{Firstname: "Charlie", Age: 35},
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
		}
		b := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Charlie", Age: 35},
		}
		result := collection.Intersect(a, b)

		expected := []User{
			{Firstname: "Charlie", Age: 35},
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
		}
		Expect(result).To(Equal(expected))
	})

	It("handles duplicate users", func() {
		a := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Alice", Age: 25}, // Duplicate
		}
		b := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Alice", Age: 25}, // Duplicate
		}
		result := collection.Intersect(a, b)

		expected := []User{
			{Firstname: "Alice", Age: 25},
		}
		Expect(result).To(Equal(expected))
	})
})
