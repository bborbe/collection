// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Map", func() {
	It("returns nil for empty slice", func() {
		var counter int
		err := collection.Map([]string{}, func(value string) error {
			counter++
			return nil
		})
		Expect(err).To(BeNil())
		Expect(counter).To(Equal(0))
	})

	It("executes function for each element successfully", func() {
		var results []string
		err := collection.Map([]string{"a", "b", "c"}, func(value string) error {
			results = append(results, value+"_processed")
			return nil
		})
		Expect(err).To(BeNil())
		Expect(results).To(Equal([]string{"a_processed", "b_processed", "c_processed"}))
	})

	It("returns error on first failure", func() {
		var counter int
		err := collection.Map([]string{"a", "b", "c"}, func(value string) error {
			counter++
			if value == "b" {
				return fmt.Errorf("error processing %s", value)
			}
			return nil
		})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("error processing b"))
		Expect(counter).To(Equal(2)) // Should stop after "b"
	})

	It("processes all elements when no errors occur", func() {
		var sum int
		err := collection.Map([]int{1, 2, 3, 4, 5}, func(value int) error {
			sum += value
			return nil
		})
		Expect(err).To(BeNil())
		Expect(sum).To(Equal(15))
	})

	It("works with custom types", func() {
		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 17},
			{Firstname: "Charlie", Age: 30},
		}
		var names []string
		err := collection.Map(users, func(user User) error {
			names = append(names, user.Firstname)
			return nil
		})
		Expect(err).To(BeNil())
		Expect(names).To(Equal([]string{"Alice", "Bob", "Charlie"}))
	})

	It("validates all elements with errors", func() {
		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "", Age: 17}, // Invalid - empty name
			{Firstname: "Charlie", Age: 30},
		}
		err := collection.Map(users, func(user User) error {
			if user.Firstname == "" {
				return fmt.Errorf("user has empty firstname")
			}
			return nil
		})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("user has empty firstname"))
	})
})
