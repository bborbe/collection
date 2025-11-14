// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	"context"
	"fmt"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Map", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("returns empty slice for empty input", func() {
		result, err := collection.Map(
			ctx,
			[]string{},
			func(ctx context.Context, value string) (int, error) {
				return len(value), nil
			},
		)
		Expect(err).To(BeNil())
		Expect(result).To(BeEmpty())
	})

	It("transforms slice of strings to slice of ints successfully", func() {
		result, err := collection.Map(
			ctx,
			[]string{"a", "bb", "ccc"},
			func(ctx context.Context, value string) (int, error) {
				return len(value), nil
			},
		)
		Expect(err).To(BeNil())
		Expect(result).To(Equal([]int{1, 2, 3}))
	})

	It("transforms slice of ints to slice of strings successfully", func() {
		result, err := collection.Map(
			ctx,
			[]int{1, 2, 3},
			func(ctx context.Context, value int) (string, error) {
				return strconv.Itoa(value), nil
			},
		)
		Expect(err).To(BeNil())
		Expect(result).To(Equal([]string{"1", "2", "3"}))
	})

	It("returns error on first failure and partial results", func() {
		result, err := collection.Map(
			ctx,
			[]string{"1", "2", "invalid", "4"},
			func(ctx context.Context, value string) (int, error) {
				num, parseErr := strconv.Atoi(value)
				if parseErr != nil {
					return 0, fmt.Errorf("failed to parse %s: %w", value, parseErr)
				}
				return num, nil
			},
		)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("failed to parse invalid"))
		Expect(result).To(Equal([]int{1, 2})) // Partial results before error
	})

	It("transforms custom types to primitive types", func() {
		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Charlie", Age: 35},
		}
		result, err := collection.Map(
			ctx,
			users,
			func(ctx context.Context, user User) (string, error) {
				return user.Firstname, nil
			},
		)
		Expect(err).To(BeNil())
		Expect(result).To(Equal([]string{"Alice", "Bob", "Charlie"}))
	})

	It("transforms primitive types to custom types", func() {
		names := []string{"Alice", "Bob", "Charlie"}
		result, err := collection.Map(
			ctx,
			names,
			func(ctx context.Context, name string) (User, error) {
				return User{Firstname: name, Age: 0}, nil
			},
		)
		Expect(err).To(BeNil())
		Expect(result).To(HaveLen(3))
		Expect(result[0].Firstname).To(Equal("Alice"))
		Expect(result[1].Firstname).To(Equal("Bob"))
		Expect(result[2].Firstname).To(Equal("Charlie"))
	})

	It("validates and transforms with errors", func() {
		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "", Age: 17}, // Invalid - empty name
			{Firstname: "Charlie", Age: 30},
		}
		result, err := collection.Map(
			ctx,
			users,
			func(ctx context.Context, user User) (string, error) {
				if user.Firstname == "" {
					return "", fmt.Errorf("user has empty firstname")
				}
				return user.Firstname, nil
			},
		)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("user has empty firstname"))
		Expect(result).To(Equal([]string{"Alice"})) // Only first element before error
	})

	It("transforms with complex calculations", func() {
		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Charlie", Age: 35},
		}
		result, err := collection.Map(
			ctx,
			users,
			func(ctx context.Context, user User) (string, error) {
				return fmt.Sprintf("%s is %d years old", user.Firstname, user.Age), nil
			},
		)
		Expect(err).To(BeNil())
		Expect(result).To(Equal([]string{
			"Alice is 25 years old",
			"Bob is 30 years old",
			"Charlie is 35 years old",
		}))
	})
})
