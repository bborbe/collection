// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Filter", func() {
	Context("with strings", func() {
		It("returns empty slice for empty input", func() {
			result := collection.Filter([]string{}, func(value string) bool {
				return len(value) > 0
			})
			Expect(result).To(BeEmpty())
		})

		It("returns all elements when all match", func() {
			result := collection.Filter([]string{"a", "b", "c"}, func(value string) bool {
				return len(value) == 1
			})
			Expect(result).To(Equal([]string{"a", "b", "c"}))
		})

		It("returns no elements when none match", func() {
			result := collection.Filter([]string{"a", "b", "c"}, func(value string) bool {
				return len(value) > 1
			})
			Expect(result).To(BeEmpty())
		})

		It("returns subset when some match", func() {
			result := collection.Filter([]string{"a", "bb", "c", "dd"}, func(value string) bool {
				return len(value) == 2
			})
			Expect(result).To(Equal([]string{"bb", "dd"}))
		})
	})

	Context("with integers", func() {
		It("filters even numbers", func() {
			result := collection.Filter([]int{1, 2, 3, 4, 5, 6}, func(value int) bool {
				return value%2 == 0
			})
			Expect(result).To(Equal([]int{2, 4, 6}))
		})

		It("filters numbers greater than threshold", func() {
			result := collection.Filter([]int{1, 5, 10, 15, 20}, func(value int) bool {
				return value > 10
			})
			Expect(result).To(Equal([]int{15, 20}))
		})
	})

	Context("with custom types", func() {
		It("filters users by age", func() {
			users := []User{
				{Firstname: "Alice", Age: 25},
				{Firstname: "Bob", Age: 17},
				{Firstname: "Charlie", Age: 30},
			}
			result := collection.Filter(users, func(user User) bool {
				return user.Age >= 18
			})
			Expect(result).To(HaveLen(2))
			Expect(result[0].Firstname).To(Equal("Alice"))
			Expect(result[1].Firstname).To(Equal("Charlie"))
		})
	})
})
