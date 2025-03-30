// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Exclude", func() {
	var list []string
	var result []string
	var exclude []string
	JustBeforeEach(func() {
		result = collection.Exclude(list, exclude...)
	})
	Context("empty", func() {
		BeforeEach(func() {
			list = []string{}
			exclude = []string{}
		})
		It("returns empty result", func() {
			Expect(result).To(HaveLen(0))
			Expect(result).To(Equal([]string{}))
		})
	})
	Context("no exclude", func() {
		BeforeEach(func() {
			list = []string{"a", "b", "c"}
			exclude = []string{}
		})
		It("returns empty result", func() {
			Expect(result).To(HaveLen(3))
			Expect(result).To(Equal([]string{"a", "b", "c"}))
		})
	})
	Context("not found exclude", func() {
		BeforeEach(func() {
			list = []string{"a", "b", "c"}
			exclude = []string{"d", "e", "f"}
		})
		It("returns empty result", func() {
			Expect(result).To(HaveLen(3))
			Expect(result).To(Equal([]string{"a", "b", "c"}))
		})
	})
	Context("with exclude", func() {
		BeforeEach(func() {
			list = []string{"a", "b", "c"}
			exclude = []string{"a"}
		})
		It("returns empty result", func() {
			Expect(result).To(HaveLen(2))
			Expect(result).To(Equal([]string{"b", "c"}))
		})
	})
})
