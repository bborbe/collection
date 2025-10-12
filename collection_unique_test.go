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
	"Unique",
	func(input []string, expected []string) {
		result := collection.Unique(input)
		Expect(result).To(Equal(expected))
	},
	Entry("empty slice", []string{}, []string{}),
	Entry("single element", []string{"a"}, []string{"a"}),
	Entry("no duplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}),
	Entry("all duplicates", []string{"a", "a", "a"}, []string{"a"}),
	Entry("mixed duplicates", []string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}),
	Entry(
		"preserves first occurrence order",
		[]string{"c", "a", "b", "a", "c"},
		[]string{"c", "a", "b"},
	),
)

var _ = DescribeTable("Unique with integers",
	func(input []int, expected []int) {
		result := collection.Unique(input)
		Expect(result).To(Equal(expected))
	},
	Entry("empty slice", []int{}, []int{}),
	Entry("single element", []int{42}, []int{42}),
	Entry("no duplicates", []int{1, 2, 3}, []int{1, 2, 3}),
	Entry("consecutive duplicates", []int{1, 1, 2, 2, 3, 3}, []int{1, 2, 3}),
	Entry("non-consecutive duplicates", []int{1, 2, 1, 3, 2}, []int{1, 2, 3}),
	Entry("all same", []int{5, 5, 5, 5}, []int{5}),
)

var _ = Describe("Unique with custom types", func() {
	It("removes duplicate users", func() {
		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Alice", Age: 25}, // Duplicate
			{Firstname: "Charlie", Age: 35},
			{Firstname: "Bob", Age: 30}, // Duplicate
		}
		result := collection.Unique(users)

		expected := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Charlie", Age: 35},
		}
		Expect(result).To(Equal(expected))
	})

	It("preserves order of first occurrence", func() {
		users := []User{
			{Firstname: "Charlie", Age: 35},
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Alice", Age: 25}, // Duplicate
		}
		result := collection.Unique(users)

		expected := []User{
			{Firstname: "Charlie", Age: 35},
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
		}
		Expect(result).To(Equal(expected))
	})
})
