// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

type CustomString string

var _ = DescribeTable("Compare",
	func(a, b string, expected int) {
		result := collection.Compare(a, b)
		Expect(result).To(Equal(expected))
	},
	Entry("equal strings", "abc", "abc", 0),
	Entry("first string less than second", "abc", "def", -1),
	Entry("first string greater than second", "def", "abc", 1),
	Entry("empty strings", "", "", 0),
	Entry("empty vs non-empty", "", "abc", -1),
	Entry("non-empty vs empty", "abc", "", 1),
	Entry("case sensitive - uppercase vs lowercase", "ABC", "abc", -1),
	Entry("case sensitive - lowercase vs uppercase", "abc", "ABC", 1),
	Entry("numeric strings", "123", "456", -1),
	Entry("different lengths - shorter first", "ab", "abc", -1),
	Entry("different lengths - longer first", "abc", "ab", 1),
)

var _ = Describe("Compare with custom string types", func() {
	It("compares custom string type", func() {
		a := CustomString("apple")
		b := CustomString("banana")
		result := collection.Compare(a, b)

		Expect(result).To(Equal(-1)) // "apple" < "banana"
	})

	It("compares equal custom strings", func() {
		a := CustomString("test")
		b := CustomString("test")
		result := collection.Compare(a, b)

		Expect(result).To(Equal(0))
	})

	It("works with string and custom string type", func() {
		a := "hello"
		b := CustomString("world")
		result := collection.Compare(a, string(b))

		Expect(result).To(Equal(-1)) // "hello" < "world"
	})
})
