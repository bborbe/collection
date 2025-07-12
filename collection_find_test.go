// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Find", func() {
	Context("with strings", func() {
		It("returns NotFoundError for empty slice", func() {
			result, err := collection.Find([]string{}, func(value string) bool {
				return value == "a"
			})
			Expect(err).To(Equal(collection.NotFoundError))
			Expect(result).To(BeNil())
		})

		It("finds first matching element", func() {
			result, err := collection.Find([]string{"a", "b", "c"}, func(value string) bool {
				return value == "b"
			})
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(*result).To(Equal("b"))
		})

		It("finds first element when multiple match", func() {
			result, err := collection.Find([]string{"a", "b", "b", "c"}, func(value string) bool {
				return value == "b"
			})
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(*result).To(Equal("b"))
		})

		It("returns NotFoundError when no element matches", func() {
			result, err := collection.Find([]string{"a", "b", "c"}, func(value string) bool {
				return value == "d"
			})
			Expect(err).To(Equal(collection.NotFoundError))
			Expect(result).To(BeNil())
		})
	})

	Context("with integers", func() {
		It("finds number greater than threshold", func() {
			result, err := collection.Find([]int{1, 5, 10, 15}, func(value int) bool {
				return value > 8
			})
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(*result).To(Equal(10))
		})

		It("finds even number", func() {
			result, err := collection.Find([]int{1, 3, 4, 5, 6}, func(value int) bool {
				return value%2 == 0
			})
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(*result).To(Equal(4))
		})
	})

	Context("with custom types", func() {
		It("finds user by name", func() {
			users := []User{
				{Firstname: "Alice", Age: 25},
				{Firstname: "Bob", Age: 17},
				{Firstname: "Charlie", Age: 30},
			}
			result, err := collection.Find(users, func(user User) bool {
				return user.Firstname == "Bob"
			})
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.Firstname).To(Equal("Bob"))
			Expect(result.Age).To(Equal(17))
		})

		It("finds user by age criteria", func() {
			users := []User{
				{Firstname: "Alice", Age: 25},
				{Firstname: "Bob", Age: 17},
				{Firstname: "Charlie", Age: 30},
			}
			result, err := collection.Find(users, func(user User) bool {
				return user.Age >= 30
			})
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result.Firstname).To(Equal("Charlie"))
		})

		It("returns NotFoundError when user not found", func() {
			users := []User{
				{Firstname: "Alice", Age: 25},
				{Firstname: "Bob", Age: 17},
			}
			result, err := collection.Find(users, func(user User) bool {
				return user.Firstname == "Dave"
			})
			Expect(err).To(Equal(collection.NotFoundError))
			Expect(result).To(BeNil())
		})
	})
})
