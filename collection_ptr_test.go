// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Ptr", func() {
	It("creates pointer to string", func() {
		value := "test"
		ptr := collection.Ptr(value)

		Expect(ptr).ToNot(BeNil())
		Expect(*ptr).To(Equal("test"))
		// Verify independent memory - modify original, ptr should be unchanged
		value = "changed"
		Expect(*ptr).To(Equal("test"))
	})

	It("creates pointer to integer", func() {
		value := 42
		ptr := collection.Ptr(value)

		Expect(ptr).ToNot(BeNil())
		Expect(*ptr).To(Equal(42))
		// Verify independent memory - modify original, ptr should be unchanged
		value = 100
		Expect(*ptr).To(Equal(42))
	})

	It("creates pointer to boolean", func() {
		value := true
		ptr := collection.Ptr(value)

		Expect(ptr).ToNot(BeNil())
		Expect(*ptr).To(BeTrue())
	})

	It("creates pointer to custom type", func() {
		user := User{Firstname: "Alice", Age: 25}
		ptr := collection.Ptr(user)

		Expect(ptr).ToNot(BeNil())
		Expect(ptr.Firstname).To(Equal("Alice"))
		Expect(ptr.Age).To(Equal(25))
		// Verify independent memory - modify original, ptr should be unchanged
		user.Firstname = "Changed"
		Expect(ptr.Firstname).To(Equal("Alice"))
	})

	It("creates pointer to zero value", func() {
		var value int
		ptr := collection.Ptr(value)

		Expect(ptr).ToNot(BeNil())
		Expect(*ptr).To(Equal(0))
	})

	It("allows modification through pointer", func() {
		value := "original"
		ptr := collection.Ptr(value)
		*ptr = "modified"

		Expect(*ptr).To(Equal("modified"))
		Expect(value).To(Equal("original")) // Original unchanged
	})
})
