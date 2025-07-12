// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Set", func() {
	var set collection.Set[User]
	BeforeEach(func() {
		set = collection.NewSet[User]()
	})
	Context("Splice", func() {
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
	Context("Remove", func() {
		It("removes existing element", func() {
			user := User{Firstname: "Alice", Age: 25}
			set.Add(user)
			Expect(set.Length()).To(Equal(1))

			set.Remove(user)
			Expect(set.Length()).To(Equal(0))
			Expect(set.Contains(user)).To(BeFalse())
		})

		It("does nothing when removing non-existent element", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			set.Add(user1)
			Expect(set.Length()).To(Equal(1))

			set.Remove(user2) // Remove element that was never added
			Expect(set.Length()).To(Equal(1))
			Expect(set.Contains(user1)).To(BeTrue())
		})

		It("removes element from set with multiple elements", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}

			set.Add(user1)
			set.Add(user2)
			set.Add(user3)
			Expect(set.Length()).To(Equal(3))

			set.Remove(user2)
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains(user1)).To(BeTrue())
			Expect(set.Contains(user2)).To(BeFalse())
			Expect(set.Contains(user3)).To(BeTrue())
		})
	})
	Context("Contains", func() {
		It("returns false for empty set", func() {
			user := User{Firstname: "Alice", Age: 25}
			Expect(set.Contains(user)).To(BeFalse())
		})

		It("returns true for existing element", func() {
			user := User{Firstname: "Alice", Age: 25}
			set.Add(user)
			Expect(set.Contains(user)).To(BeTrue())
		})

		It("returns false for non-existent element", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			set.Add(user1)
			Expect(set.Contains(user2)).To(BeFalse())
		})

		It("works with multiple elements", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}

			set.Add(user1)
			set.Add(user3)

			Expect(set.Contains(user1)).To(BeTrue())
			Expect(set.Contains(user2)).To(BeFalse())
			Expect(set.Contains(user3)).To(BeTrue())
		})
	})
})
