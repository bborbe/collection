// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("UnPtr", func() {
	It("dereferences string pointer", func() {
		value := "test"
		ptr := &value
		result := collection.UnPtr(ptr)

		Expect(result).To(Equal("test"))
	})

	It("dereferences integer pointer", func() {
		value := 42
		ptr := &value
		result := collection.UnPtr(ptr)

		Expect(result).To(Equal(42))
	})

	It("dereferences boolean pointer", func() {
		value := true
		ptr := &value
		result := collection.UnPtr(ptr)

		Expect(result).To(BeTrue())
	})

	It("dereferences custom type pointer", func() {
		user := User{Firstname: "Alice", Age: 25}
		ptr := &user
		result := collection.UnPtr(ptr)

		Expect(result.Firstname).To(Equal("Alice"))
		Expect(result.Age).To(Equal(25))
	})

	It("returns zero value for nil string pointer", func() {
		var ptr *string
		result := collection.UnPtr(ptr)

		Expect(result).To(Equal(""))
	})

	It("returns zero value for nil integer pointer", func() {
		var ptr *int
		result := collection.UnPtr(ptr)

		Expect(result).To(Equal(0))
	})

	It("returns zero value for nil boolean pointer", func() {
		var ptr *bool
		result := collection.UnPtr(ptr)

		Expect(result).To(BeFalse())
	})

	It("returns zero value for nil custom type pointer", func() {
		var ptr *User
		result := collection.UnPtr(ptr)

		Expect(result.Firstname).To(Equal(""))
		Expect(result.Age).To(Equal(0))
	})

	It("works with Ptr function", func() {
		original := "test"
		ptr := collection.Ptr(original)
		result := collection.UnPtr(ptr)

		Expect(result).To(Equal("test"))
	})
})
