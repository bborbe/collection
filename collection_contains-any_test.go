// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = DescribeTable("ContainsAny",
	func(a []string, b []string, expectedResult bool) {
		Expect(collection.ContainsAny(a, b)).To(Equal(expectedResult))
	},
	Entry("empty b", []string{"a", "b", "c"}, []string{}, false),
	Entry("empty a and b", []string{}, []string{}, false),
	Entry("empty a", []string{}, []string{"a"}, false),
	Entry("all match", []string{"a", "b", "c"}, []string{"a", "b", "c"}, true),
	Entry("one match", []string{"a", "b", "c"}, []string{"a"}, true),
	Entry("one match at end", []string{"a", "b", "c"}, []string{"c"}, true),
	Entry("no match", []string{"a", "b"}, []string{"c", "d"}, false),
	Entry("partial match", []string{"a", "b", "c"}, []string{"c", "d", "e"}, true),
	Entry("duplicate elements in b", []string{"a"}, []string{"a", "a"}, true),
	Entry("a subset of b", []string{"a", "b"}, []string{"a", "b", "c"}, true),
	Entry("b subset of a", []string{"a", "b", "c", "d"}, []string{"b", "c"}, true),
)
