// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = DescribeTable("ContainsAll",
	func(a []string, b []string, expectedResult bool) {
		Expect(collection.ContainsAll(a, b)).To(Equal(expectedResult))
	},
	Entry("empty", []string{}, []string{}, true),
	Entry("multiple", []string{"a", "b", "c"}, []string{"a", "b", "c"}, true),
	Entry("same length", []string{"a"}, []string{"b"}, false),
	Entry("different length", []string{"a"}, []string{"b", "c"}, false),
	Entry("different length", []string{"a"}, []string{"a", "a"}, true),
	Entry("different length", []string{"a", "a"}, []string{"a"}, true),
	Entry("b contains extra element", []string{"a", "b"}, []string{"a", "b", "c"}, false),
	Entry("a contains all elements of b", []string{"a", "b", "c", "d"}, []string{"a", "b"}, true),
	Entry("a subset of b", []string{"a", "b"}, []string{"a", "b", "c"}, false),
)
