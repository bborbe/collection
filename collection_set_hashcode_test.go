// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("SetHashCode", func() {
	var set collection.SetHashCode[User]
	BeforeEach(func() {
		set = collection.NewSetHashCode[User]()
	})
	It("returns no error", func() {
		ptr := collection.Ptr("admin")
		set.Add(User{
			Firstname: "Ben",
			Lastname:  "Bo",
			Age:       23,
			Groups:    ptr,
		})
		set.Add(User{
			Firstname: "Ben",
			Lastname:  "Bo",
			Age:       23,
			Groups:    ptr,
		})
		Expect(set.Slice()).To(HaveLen(1))
	})
	It("returns no error", func() {
		set.Add(User{
			Firstname: "Ben",
			Lastname:  "Bo",
			Age:       23,
		})
		set.Add(User{
			Firstname: "Ben",
			Lastname:  "Bo",
			Age:       24,
		})
		Expect(set.Slice()).To(HaveLen(2))
	})
	Context("Length", func() {
		var length int
		JustBeforeEach(func() {
			length = set.Length()
		})
		Context("empty", func() {
			It("has correct lenght", func() {
				Expect(length).To(Equal(0))
			})
		})
		Context("with elements", func() {
			BeforeEach(func() {
				set.Add(User{
					Firstname: "Ben",
					Lastname:  "Bo",
					Age:       23,
				})
				set.Add(User{
					Firstname: "Ben",
					Lastname:  "Bo",
					Age:       24,
				})
			})
			It("has correct lenght", func() {
				Expect(length).To(Equal(2))
			})
		})
	})
})
