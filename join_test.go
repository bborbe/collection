// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = DescribeTable("Join",
	func(a []string, b []string, expectedResult []string) {
		Expect(collection.Join(a, b)).To(Equal(expectedResult))
	},
	Entry("empty", []string{}, []string{}, []string{}),
	Entry("with elements", []string{"a", "b"}, []string{"c", "d"}, []string{"a", "b", "c", "d"}),
)
