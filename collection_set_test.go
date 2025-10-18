// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	"encoding"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Set", func() {
	var set collection.Set[User]
	BeforeEach(func() {
		set = collection.NewSet[User]()
	})
	Context("NewSet with variadic constructor", func() {
		It("creates empty set with no arguments", func() {
			emptySet := collection.NewSet[int]()
			Expect(emptySet.Length()).To(Equal(0))
		})

		It("creates set with single element", func() {
			singleSet := collection.NewSet(42)
			Expect(singleSet.Length()).To(Equal(1))
			Expect(singleSet.Contains(42)).To(BeTrue())
		})

		It("creates set with multiple elements", func() {
			multiSet := collection.NewSet(1, 2, 3, 4, 5)
			Expect(multiSet.Length()).To(Equal(5))
			Expect(multiSet.Contains(1)).To(BeTrue())
			Expect(multiSet.Contains(3)).To(BeTrue())
			Expect(multiSet.Contains(5)).To(BeTrue())
		})

		It("handles duplicate elements correctly", func() {
			dupSet := collection.NewSet(1, 2, 2, 3, 3, 3)
			Expect(dupSet.Length()).To(Equal(3))
			Expect(dupSet.Contains(1)).To(BeTrue())
			Expect(dupSet.Contains(2)).To(BeTrue())
			Expect(dupSet.Contains(3)).To(BeTrue())
		})

		It("works with string type", func() {
			stringSet := collection.NewSet("apple", "banana", "cherry")
			Expect(stringSet.Length()).To(Equal(3))
			Expect(stringSet.Contains("apple")).To(BeTrue())
			Expect(stringSet.Contains("banana")).To(BeTrue())
			Expect(stringSet.Contains("cherry")).To(BeTrue())
		})

		It("works with struct type", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			userSet := collection.NewSet(user1, user2)
			Expect(userSet.Length()).To(Equal(2))
			Expect(userSet.Contains(user1)).To(BeTrue())
			Expect(userSet.Contains(user2)).To(BeTrue())
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
	Context("ContainsAll", func() {
		It("returns true for empty list", func() {
			user := User{Firstname: "Alice", Age: 25}
			set.Add(user)
			Expect(set.ContainsAll()).To(BeTrue())
		})

		It("returns true when all elements exist", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2, user3)

			Expect(set.ContainsAll(user1, user2)).To(BeTrue())
			Expect(set.ContainsAll(user1, user2, user3)).To(BeTrue())
		})

		It("returns false when at least one element is missing", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2)

			Expect(set.ContainsAll(user1, user3)).To(BeFalse())
			Expect(set.ContainsAll(user1, user2, user3)).To(BeFalse())
		})

		It("returns false when all elements are missing", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1)

			Expect(set.ContainsAll(user2, user3)).To(BeFalse())
		})
	})
	Context("ContainsAny", func() {
		It("returns false for empty list", func() {
			user := User{Firstname: "Alice", Age: 25}
			set.Add(user)
			Expect(set.ContainsAny()).To(BeFalse())
		})

		It("returns true when at least one element exists", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2)

			Expect(set.ContainsAny(user1)).To(BeTrue())
			Expect(set.ContainsAny(user1, user3)).To(BeTrue())
			Expect(set.ContainsAny(user3, user1)).To(BeTrue())
		})

		It("returns false when no elements exist", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1)

			Expect(set.ContainsAny(user2, user3)).To(BeFalse())
		})

		It("returns true when all elements exist", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			set.Add(user1, user2)

			Expect(set.ContainsAny(user1, user2)).To(BeTrue())
		})
	})
	Context("String", func() {
		It("returns empty set string", func() {
			Expect(set.String()).To(Equal("Set[]"))
		})

		It("returns string with single element", func() {
			set.Add(User{Firstname: "Alice", Age: 25})
			result := set.String()
			Expect(result).To(HavePrefix("Set["))
			Expect(result).To(HaveSuffix("]"))
			Expect(result).To(ContainSubstring("Alice"))
			Expect(result).To(ContainSubstring("25"))
		})

		It("returns string with multiple elements", func() {
			set.Add(User{Firstname: "Alice", Age: 25})
			set.Add(User{Firstname: "Bob", Age: 30})
			result := set.String()
			Expect(result).To(HavePrefix("Set["))
			Expect(result).To(HaveSuffix("]"))
			Expect(result).To(ContainSubstring("Alice"))
			Expect(result).To(ContainSubstring("Bob"))
			Expect(result).To(ContainSubstring(", "))
		})

		It("works with simple types", func() {
			intSet := collection.NewSet(1, 2, 3)
			result := intSet.String()
			Expect(result).To(HavePrefix("Set["))
			Expect(result).To(HaveSuffix("]"))
			Expect(result).To(ContainSubstring("1"))
			Expect(result).To(ContainSubstring("2"))
			Expect(result).To(ContainSubstring("3"))
		})
	})
	Context("UnmarshalText", func() {
		It("parses single value", func() {
			set := collection.NewSet[string]()
			unmarshaler := set.(encoding.TextUnmarshaler)
			err := unmarshaler.UnmarshalText([]byte("value1"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(1))
			Expect(set.Contains("value1")).To(BeTrue())
		})

		It("parses multiple comma-separated values", func() {
			set := collection.NewSet[string]()
			unmarshaler := set.(encoding.TextUnmarshaler)
			err := unmarshaler.UnmarshalText([]byte("value1,value2,value3"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
			Expect(set.Contains("value3")).To(BeTrue())
		})

		It("handles values with whitespace", func() {
			set := collection.NewSet[string]()
			unmarshaler := set.(encoding.TextUnmarshaler)
			err := unmarshaler.UnmarshalText([]byte("value1, value2 , value3"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
			Expect(set.Contains("value3")).To(BeTrue())
		})

		It("handles empty string", func() {
			set := collection.NewSet[string]()
			unmarshaler := set.(encoding.TextUnmarshaler)
			err := unmarshaler.UnmarshalText([]byte(""))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(0))
		})

		It("handles trailing comma", func() {
			set := collection.NewSet[string]()
			unmarshaler := set.(encoding.TextUnmarshaler)
			err := unmarshaler.UnmarshalText([]byte("value1,value2,"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
		})

		It("handles leading comma", func() {
			set := collection.NewSet[string]()
			unmarshaler := set.(encoding.TextUnmarshaler)
			err := unmarshaler.UnmarshalText([]byte(",value1,value2"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
		})

		It("handles duplicate values", func() {
			set := collection.NewSet[string]()
			unmarshaler := set.(encoding.TextUnmarshaler)
			err := unmarshaler.UnmarshalText([]byte("value1,value2,value1"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
		})

		It("replaces existing set contents", func() {
			set := collection.NewSet("old1", "old2")
			unmarshaler := set.(encoding.TextUnmarshaler)
			err := unmarshaler.UnmarshalText([]byte("new1,new2"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains("new1")).To(BeTrue())
			Expect(set.Contains("new2")).To(BeTrue())
			Expect(set.Contains("old1")).To(BeFalse())
			Expect(set.Contains("old2")).To(BeFalse())
		})
	})
	Context("ParseSetFromStrings", func() {
		It("creates set from string slice", func() {
			set := collection.ParseSetFromStrings[string]([]string{"a", "b", "c"})
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
		})

		It("handles empty slice", func() {
			set := collection.ParseSetFromStrings[string]([]string{})
			Expect(set.Length()).To(Equal(0))
		})

		It("handles duplicates", func() {
			set := collection.ParseSetFromStrings[string]([]string{"a", "b", "a", "c", "b"})
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
		})
	})
	Context("ParseSetFromString", func() {
		It("parses comma-separated values", func() {
			set := collection.ParseSetFromString[string]("a,b,c")
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
		})

		It("trims whitespace", func() {
			set := collection.ParseSetFromString[string](" a , b , c ")
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
		})

		It("handles empty string", func() {
			set := collection.ParseSetFromString[string]("")
			Expect(set.Length()).To(Equal(0))
		})
	})
})
