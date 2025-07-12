// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = DescribeTable("Contains",
	func(list []string, value string, expectedResult bool) {
		Expect(collection.Contains(list, value)).To(Equal(expectedResult))
	},
	Entry("empty list", []string{}, "a", false),
	Entry("single element - found", []string{"a"}, "a", true),
	Entry("single element - not found", []string{"a"}, "b", false),
	Entry("multiple elements - found at start", []string{"a", "b", "c"}, "a", true),
	Entry("multiple elements - found at middle", []string{"a", "b", "c"}, "b", true),
	Entry("multiple elements - found at end", []string{"a", "b", "c"}, "c", true),
	Entry("multiple elements - not found", []string{"a", "b", "c"}, "d", false),
	Entry("duplicate elements - found", []string{"a", "a", "b"}, "a", true),
)

var _ = DescribeTable("Contains with integers",
	func(list []int, value int, expectedResult bool) {
		Expect(collection.Contains(list, value)).To(Equal(expectedResult))
	},
	Entry("empty list", []int{}, 1, false),
	Entry("single element - found", []int{42}, 42, true),
	Entry("single element - not found", []int{42}, 24, false),
	Entry("multiple elements - found", []int{1, 2, 3}, 2, true),
	Entry("multiple elements - not found", []int{1, 2, 3}, 4, false),
)
