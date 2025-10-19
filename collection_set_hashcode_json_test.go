// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("SetHashCode JSON", func() {
	Context("MarshalJSON", func() {
		It("marshals empty set to empty JSON array", func() {
			set := collection.NewSetHashCode[User]()
			data, err := json.Marshal(set)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal("[]"))
		})

		It("marshals set with hashable types", func() {
			set := collection.NewSetHashCode(
				User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
				User{Firstname: "Bob", Lastname: "Brown", Age: 25},
			)
			data, err := json.Marshal(set)
			Expect(err).NotTo(HaveOccurred())

			var result []User
			err = json.Unmarshal(data, &result)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(HaveLen(2))
		})
	})

	Context("UnmarshalJSON", func() {
		It("unmarshals empty JSON array to empty set", func() {
			set := collection.NewSetHashCode[User]()
			err := json.Unmarshal([]byte("[]"), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(0))
		})

		It("unmarshals JSON array with hashable types", func() {
			set := collection.NewSetHashCode[User]()
			jsonData := `[{"Firstname":"Alice","Lastname":"Anderson","Age":30},{"Firstname":"Bob","Lastname":"Brown","Age":25}]`
			err := json.Unmarshal([]byte(jsonData), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(
				set.Contains(User{Firstname: "Alice", Lastname: "Anderson", Age: 30}),
			).To(BeTrue())
			Expect(set.Contains(User{Firstname: "Bob", Lastname: "Brown", Age: 25})).To(BeTrue())
		})

		It("handles duplicate elements using HashCode method", func() {
			set := collection.NewSetHashCode[User]()
			// Exact duplicates - same hash code, last one wins
			jsonData := `[{"Firstname":"Alice","Lastname":"Anderson","Age":30},{"Firstname":"Alice","Lastname":"Anderson","Age":30}]`
			err := json.Unmarshal([]byte(jsonData), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(1))
			Expect(
				set.Contains(User{Firstname: "Alice", Lastname: "Anderson", Age: 30}),
			).To(BeTrue())
		})

		It("replaces existing set contents", func() {
			set := collection.NewSetHashCode(
				User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
				User{Firstname: "Bob", Lastname: "Brown", Age: 25},
			)
			jsonData := `[{"Firstname":"Charlie","Lastname":"Clark","Age":35}]`
			err := json.Unmarshal([]byte(jsonData), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(1))
			Expect(
				set.Contains(User{Firstname: "Alice", Lastname: "Anderson", Age: 30}),
			).To(BeFalse())
			Expect(
				set.Contains(User{Firstname: "Charlie", Lastname: "Clark", Age: 35}),
			).To(BeTrue())
		})
	})

	Context("Round-trip marshal/unmarshal", func() {
		It("preserves set contents", func() {
			original := collection.NewSetHashCode(
				User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
				User{Firstname: "Bob", Lastname: "Brown", Age: 25},
				User{Firstname: "Charlie", Lastname: "Clark", Age: 35},
			)
			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			restored := collection.NewSetHashCode[User]()
			err = json.Unmarshal(data, restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Length()).To(Equal(original.Length()))
			for _, val := range original.Slice() {
				Expect(restored.Contains(val)).To(BeTrue())
			}
		})
	})

	Context("Integration with structs containing sets", func() {
		It("marshals and unmarshals struct with SetHashCode field", func() {
			type Container struct {
				Name   string                       `json:"name"`
				Values collection.SetHashCode[User] `json:"values"`
			}

			original := Container{
				Name: "test",
				Values: collection.NewSetHashCode(
					User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
					User{Firstname: "Bob", Lastname: "Brown", Age: 25},
				),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored Container
			restored.Values = collection.NewSetHashCode[User]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Name).To(Equal("test"))
			Expect(restored.Values.Length()).To(Equal(2))
			Expect(
				restored.Values.Contains(User{Firstname: "Alice", Lastname: "Anderson", Age: 30}),
			).To(BeTrue())
			Expect(
				restored.Values.Contains(User{Firstname: "Bob", Lastname: "Brown", Age: 25}),
			).To(BeTrue())
		})

		It("handles empty SetHashCode field", func() {
			type Application struct {
				Name  string                       `json:"name"`
				Users collection.SetHashCode[User] `json:"users"`
			}

			original := Application{
				Name:  "empty",
				Users: collection.NewSetHashCode[User](),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal(`{"name":"empty","users":[]}`))

			var restored Application
			restored.Users = collection.NewSetHashCode[User]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Name).To(Equal("empty"))
			Expect(restored.Users.Length()).To(Equal(0))
		})

		It("handles multiple SetHashCode fields", func() {
			type MultiApp struct {
				Admins collection.SetHashCode[User] `json:"admins"`
				Users  collection.SetHashCode[User] `json:"users"`
			}

			original := MultiApp{
				Admins: collection.NewSetHashCode(
					User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
				),
				Users: collection.NewSetHashCode(
					User{Firstname: "Bob", Lastname: "Brown", Age: 25},
					User{Firstname: "Charlie", Lastname: "Clark", Age: 35},
				),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored MultiApp
			restored.Admins = collection.NewSetHashCode[User]()
			restored.Users = collection.NewSetHashCode[User]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Admins.Length()).To(Equal(1))
			Expect(restored.Users.Length()).To(Equal(2))
		})
	})
})
