// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	"context"
	"encoding"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("Set", func() {
	var set collection.Set[User]
	var ctx context.Context
	BeforeEach(func() {
		set = collection.NewSet[User]()
		ctx = context.Background()
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
			It("has correct length", func() {
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
			It("has correct length", func() {
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
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte("value1"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(1))
			Expect(set.Contains("value1")).To(BeTrue())
		})

		It("parses multiple comma-separated values", func() {
			set := collection.NewSet[string]()
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte("value1,value2,value3"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
			Expect(set.Contains("value3")).To(BeTrue())
		})

		It("handles values with whitespace", func() {
			set := collection.NewSet[string]()
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte("value1, value2 , value3"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
			Expect(set.Contains("value3")).To(BeTrue())
		})

		It("handles empty string", func() {
			set := collection.NewSet[string]()
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte(""))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(0))
		})

		It("handles trailing comma", func() {
			set := collection.NewSet[string]()
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte("value1,value2,"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
		})

		It("handles leading comma", func() {
			set := collection.NewSet[string]()
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte(",value1,value2"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
		})

		It("handles duplicate values", func() {
			set := collection.NewSet[string]()
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte("value1,value2,value1"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains("value1")).To(BeTrue())
			Expect(set.Contains("value2")).To(BeTrue())
		})

		It("replaces existing set contents", func() {
			set := collection.NewSet("old1", "old2")
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
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
	Context("UnmarshalText with custom string types", func() {
		It("unmarshals into Set[CustomStringType]", func() {
			set := collection.NewSet[CustomStringType]()
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte("alpha,beta,gamma"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains(CustomStringType("alpha"))).To(BeTrue())
			Expect(set.Contains(CustomStringType("beta"))).To(BeTrue())
			Expect(set.Contains(CustomStringType("gamma"))).To(BeTrue())
		})

		It("marshals from Set[CustomStringType]", func() {
			set := collection.NewSet[CustomStringType](
				CustomStringType("alpha"),
				CustomStringType("beta"),
				CustomStringType("gamma"),
			)
			marshaler, ok := set.(encoding.TextMarshaler)
			Expect(ok).To(BeTrue())
			data, err := marshaler.MarshalText()
			Expect(err).NotTo(HaveOccurred())

			// Parse the result back to verify it contains all elements
			result := string(data)
			Expect(result).To(ContainSubstring("alpha"))
			Expect(result).To(ContainSubstring("beta"))
			Expect(result).To(ContainSubstring("gamma"))
		})

		It("round-trips marshal/unmarshal with custom string type", func() {
			original := collection.NewSet[CustomStringType](
				CustomStringType("one"),
				CustomStringType("two"),
				CustomStringType("three"),
			)

			// Marshal
			marshaler, ok := original.(encoding.TextMarshaler)
			Expect(ok).To(BeTrue())
			data, err := marshaler.MarshalText()
			Expect(err).NotTo(HaveOccurred())

			// Unmarshal
			reconstructed := collection.NewSet[CustomStringType]()
			unmarshaler, ok := reconstructed.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err = unmarshaler.UnmarshalText(data)
			Expect(err).NotTo(HaveOccurred())

			// Verify
			Expect(reconstructed.Length()).To(Equal(original.Length()))
			Expect(reconstructed.Contains(CustomStringType("one"))).To(BeTrue())
			Expect(reconstructed.Contains(CustomStringType("two"))).To(BeTrue())
			Expect(reconstructed.Contains(CustomStringType("three"))).To(BeTrue())
		})
	})

	Context("Each", func() {
		It("returns nil for empty set", func() {
			err := set.Each(ctx, func(ctx context.Context, u User) error {
				return nil
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("calls fn for each element", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2, user3)

			visited := make(map[string]bool)
			err := set.Each(ctx, func(ctx context.Context, u User) error {
				visited[u.Firstname] = true
				return nil
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(visited).To(HaveLen(3))
			Expect(visited["Alice"]).To(BeTrue())
			Expect(visited["Bob"]).To(BeTrue())
			Expect(visited["Charlie"]).To(BeTrue())
		})

		It("stops on first error", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2, user3)

			count := 0
			testErr := errors.New("test error")
			err := set.Each(ctx, func(ctx context.Context, u User) error {
				count++
				if u.Firstname == "Bob" {
					return testErr
				}
				return nil
			})

			Expect(err).To(Equal(testErr))
			Expect(count).To(BeNumerically(">=", 1))
			Expect(count).To(BeNumerically("<=", 3))
		})

		It("works with simple types", func() {
			intSet := collection.NewSet(1, 2, 3, 4, 5)
			sum := 0
			err := intSet.Each(ctx, func(ctx context.Context, n int) error {
				sum += n
				return nil
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(sum).To(Equal(15))
		})

		It("allows mutation of external state", func() {
			intSet := collection.NewSet(1, 2, 3)
			results := make([]int, 0)
			err := intSet.Each(ctx, func(ctx context.Context, n int) error {
				results = append(results, n*2)
				return nil
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(results).To(HaveLen(3))
			Expect(results).To(ContainElements(2, 4, 6))
		})
	})

	Context("Clone", func() {
		It("creates independent copy of empty set", func() {
			clone := set.Clone()
			Expect(clone.Length()).To(Equal(0))
		})

		It("creates independent copy with all elements", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			set.Add(user1, user2)

			clone := set.Clone()
			Expect(clone.Length()).To(Equal(2))
			Expect(clone.Contains(user1)).To(BeTrue())
			Expect(clone.Contains(user2)).To(BeTrue())
		})

		It("modifications to clone don't affect original", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2)

			clone := set.Clone()
			clone.Add(user3)
			clone.Remove(user1)

			Expect(clone.Length()).To(Equal(2))
			Expect(clone.Contains(user2)).To(BeTrue())
			Expect(clone.Contains(user3)).To(BeTrue())

			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains(user1)).To(BeTrue())
			Expect(set.Contains(user2)).To(BeTrue())
			Expect(set.Contains(user3)).To(BeFalse())
		})
	})

	Context("Without", func() {
		It("returns empty set when excluding all elements", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			set.Add(user1, user2)

			result := set.Without(user1, user2)
			Expect(result.Length()).To(Equal(0))
		})

		It("returns full set when excluding no elements", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			set.Add(user1, user2)

			result := set.Without()
			Expect(result.Length()).To(Equal(2))
			Expect(result.Contains(user1)).To(BeTrue())
			Expect(result.Contains(user2)).To(BeTrue())
		})

		It("returns set without specified elements", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2, user3)

			result := set.Without(user2)
			Expect(result.Length()).To(Equal(2))
			Expect(result.Contains(user1)).To(BeTrue())
			Expect(result.Contains(user2)).To(BeFalse())
			Expect(result.Contains(user3)).To(BeTrue())
		})

		It("doesn't modify original set", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2, user3)

			result := set.Without(user2, user3)

			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains(user1)).To(BeTrue())
			Expect(set.Contains(user2)).To(BeTrue())
			Expect(set.Contains(user3)).To(BeTrue())

			Expect(result.Length()).To(Equal(1))
			Expect(result.Contains(user1)).To(BeTrue())
		})

		It("handles excluding non-existent elements", func() {
			user1 := User{Firstname: "Alice", Age: 25}
			user2 := User{Firstname: "Bob", Age: 30}
			user3 := User{Firstname: "Charlie", Age: 35}
			set.Add(user1, user2)

			result := set.Without(user3)
			Expect(result.Length()).To(Equal(2))
			Expect(result.Contains(user1)).To(BeTrue())
			Expect(result.Contains(user2)).To(BeTrue())
		})

		It("works with simple types", func() {
			intSet := collection.NewSet(1, 2, 3, 4, 5)
			result := intSet.Without(2, 4)
			Expect(result.Length()).To(Equal(3))
			Expect(result.Contains(1)).To(BeTrue())
			Expect(result.Contains(3)).To(BeTrue())
			Expect(result.Contains(5)).To(BeTrue())
		})
	})

	Context("Type aliases for Set", func() {
		It("unmarshals into type alias CustomStringTypeSet", func() {
			var set CustomStringTypeSet = collection.NewSet[CustomStringType]()
			unmarshaler, ok := set.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err := unmarshaler.UnmarshalText([]byte("foo,bar,baz"))
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains(CustomStringType("foo"))).To(BeTrue())
			Expect(set.Contains(CustomStringType("bar"))).To(BeTrue())
			Expect(set.Contains(CustomStringType("baz"))).To(BeTrue())
		})

		It("marshals from type alias CustomStringTypeSet", func() {
			var set CustomStringTypeSet = collection.NewSet[CustomStringType](
				CustomStringType("foo"),
				CustomStringType("bar"),
			)
			marshaler, ok := set.(encoding.TextMarshaler)
			Expect(ok).To(BeTrue())
			data, err := marshaler.MarshalText()
			Expect(err).NotTo(HaveOccurred())

			result := string(data)
			Expect(result).To(ContainSubstring("foo"))
			Expect(result).To(ContainSubstring("bar"))
		})

		It("round-trips marshal/unmarshal with type alias", func() {
			var original CustomStringTypeSet = collection.NewSet[CustomStringType](
				CustomStringType("apple"),
				CustomStringType("banana"),
				CustomStringType("cherry"),
			)

			// Marshal
			marshaler, ok := original.(encoding.TextMarshaler)
			Expect(ok).To(BeTrue())
			data, err := marshaler.MarshalText()
			Expect(err).NotTo(HaveOccurred())

			// Unmarshal
			var reconstructed CustomStringTypeSet = collection.NewSet[CustomStringType]()
			unmarshaler, ok := reconstructed.(encoding.TextUnmarshaler)
			Expect(ok).To(BeTrue())
			err = unmarshaler.UnmarshalText(data)
			Expect(err).NotTo(HaveOccurred())

			// Verify
			Expect(reconstructed.Length()).To(Equal(original.Length()))
			Expect(reconstructed.Contains(CustomStringType("apple"))).To(BeTrue())
			Expect(reconstructed.Contains(CustomStringType("banana"))).To(BeTrue())
			Expect(reconstructed.Contains(CustomStringType("cherry"))).To(BeTrue())
		})
	})
})
