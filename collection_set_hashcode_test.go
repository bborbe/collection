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
	Context("NewSetHashCode with variadic constructor", func() {
		It("creates empty set with no arguments", func() {
			emptySet := collection.NewSetHashCode[User]()
			Expect(emptySet.Length()).To(Equal(0))
		})

		It("creates set with single element", func() {
			user := User{Firstname: "Alice", Age: 25}
			singleSet := collection.NewSetHashCode(user)
			Expect(singleSet.Length()).To(Equal(1))
			Expect(singleSet.Contains(user)).To(BeTrue())
		})

		It("creates set with multiple elements", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			multiSet := collection.NewSetHashCode(user1, user2, user3)
			Expect(multiSet.Length()).To(Equal(3))
			Expect(multiSet.Contains(user1)).To(BeTrue())
			Expect(multiSet.Contains(user2)).To(BeTrue())
			Expect(multiSet.Contains(user3)).To(BeTrue())
		})

		It("handles duplicate elements correctly", func() {
			user := User{Firstname: "Alice", Lastname: "Smith", Age: 25}
			dupSet := collection.NewSetHashCode(user, user, user)
			Expect(dupSet.Length()).To(Equal(1))
			Expect(dupSet.Contains(user)).To(BeTrue())
		})

		It("treats different users as separate elements", func() {
			user1 := User{Firstname: "Alice", Lastname: "Smith", Age: 25}
			user2 := User{Firstname: "Alice", Lastname: "Jones", Age: 25}
			hashSet := collection.NewSetHashCode(user1, user2)
			Expect(hashSet.Length()).To(Equal(2))
		})
	})
	Context("Add with variadic parameters", func() {
		It("adds multiple elements in single call", func() {
			set.Add(
				User{Firstname: "Alice", Age: 25},
				User{Firstname: "Bob", Age: 30},
				User{Firstname: "Charlie", Age: 35},
			)
			Expect(set.Length()).To(Equal(3))
		})

		It("handles duplicates in variadic add", func() {
			user := User{Firstname: "Alice", Age: 25}
			set.Add(user, user, user)
			Expect(set.Length()).To(Equal(1))
		})

		It("adds single element", func() {
			user := User{Firstname: "Alice", Age: 25}
			set.Add(user)
			Expect(set.Length()).To(Equal(1))
			Expect(set.Contains(user)).To(BeTrue())
		})
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
	})
})
