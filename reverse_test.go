// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Reverse", func() {
	var input []string
	var result []string
	BeforeEach(func() {
		input = []string{"a", "b", "c"}
	})
	JustBeforeEach(func() {
		result = collection.Reverse(input)
	})
	It("returns correct result", func() {
		Expect(result).To(Equal([]string{"c", "b", "a"}))
	})
})
