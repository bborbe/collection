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

var _ = Describe("SetEqual JSON", func() {
	Context("MarshalJSON", func() {
		It("marshals empty set to empty JSON array", func() {
			set := collection.NewSetEqual[User]()
			data, err := json.Marshal(set)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal("[]"))
		})

		It("marshals set with custom Equal types", func() {
			set := collection.NewSetEqual(
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

		It("preserves insertion order", func() {
			set := collection.NewSetEqual(
				User{Firstname: "Charlie", Lastname: "Clark", Age: 35},
				User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
				User{Firstname: "Bob", Lastname: "Brown", Age: 25},
			)
			data, err := json.Marshal(set)
			Expect(err).NotTo(HaveOccurred())

			var result []User
			err = json.Unmarshal(data, &result)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(HaveLen(3))
			Expect(result[0].Firstname).To(Equal("Charlie"))
			Expect(result[1].Firstname).To(Equal("Alice"))
			Expect(result[2].Firstname).To(Equal("Bob"))
		})
	})

	Context("UnmarshalJSON", func() {
		It("unmarshals empty JSON array to empty set", func() {
			set := collection.NewSetEqual[User]()
			err := json.Unmarshal([]byte("[]"), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(0))
		})

		It("unmarshals JSON array with custom types", func() {
			set := collection.NewSetEqual[User]()
			jsonData := `[{"Firstname":"Alice","Lastname":"Anderson","Age":30},{"Firstname":"Bob","Lastname":"Brown","Age":25}]`
			err := json.Unmarshal([]byte(jsonData), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(
				set.Contains(User{Firstname: "Alice", Lastname: "Anderson", Age: 30}),
			).To(BeTrue())
			Expect(set.Contains(User{Firstname: "Bob", Lastname: "Brown", Age: 25})).To(BeTrue())
		})

		It("handles duplicate elements using Equal method", func() {
			set := collection.NewSetEqual[User]()
			// Exact duplicates - Equal returns true
			jsonData := `[{"Firstname":"Alice","Lastname":"Anderson","Age":30},{"Firstname":"Alice","Lastname":"Anderson","Age":30}]`
			err := json.Unmarshal([]byte(jsonData), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(1))
		})

		It("replaces existing set contents", func() {
			set := collection.NewSetEqual(
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
			original := collection.NewSetEqual(
				User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
				User{Firstname: "Bob", Lastname: "Brown", Age: 25},
				User{Firstname: "Charlie", Lastname: "Clark", Age: 35},
			)
			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			restored := collection.NewSetEqual[User]()
			err = json.Unmarshal(data, restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Length()).To(Equal(original.Length()))
			for _, val := range original.Slice() {
				Expect(restored.Contains(val)).To(BeTrue())
			}
		})
	})

	Context("Integration with structs containing SetEqual", func() {
		It("marshals and unmarshals struct with SetEqual field", func() {
			type Application struct {
				Name  string                    `json:"name"`
				Users collection.SetEqual[User] `json:"users"`
			}

			original := Application{
				Name: "myapp",
				Users: collection.NewSetEqual(
					User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
					User{Firstname: "Bob", Lastname: "Brown", Age: 25},
				),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored Application
			restored.Users = collection.NewSetEqual[User]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Name).To(Equal("myapp"))
			Expect(restored.Users.Length()).To(Equal(2))
			Expect(
				restored.Users.Contains(User{Firstname: "Alice", Lastname: "Anderson", Age: 30}),
			).To(BeTrue())
		})

		It("handles empty SetEqual field", func() {
			type Application struct {
				Name  string                    `json:"name"`
				Users collection.SetEqual[User] `json:"users"`
			}

			original := Application{
				Name:  "empty",
				Users: collection.NewSetEqual[User](),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal(`{"name":"empty","users":[]}`))

			var restored Application
			restored.Users = collection.NewSetEqual[User]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Name).To(Equal("empty"))
			Expect(restored.Users.Length()).To(Equal(0))
		})

		It("handles multiple SetEqual fields", func() {
			type MultiApp struct {
				Admins collection.SetEqual[User] `json:"admins"`
				Users  collection.SetEqual[User] `json:"users"`
			}

			original := MultiApp{
				Admins: collection.NewSetEqual(
					User{Firstname: "Alice", Lastname: "Anderson", Age: 30},
				),
				Users: collection.NewSetEqual(
					User{Firstname: "Bob", Lastname: "Brown", Age: 25},
					User{Firstname: "Charlie", Lastname: "Clark", Age: 35},
				),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored MultiApp
			restored.Admins = collection.NewSetEqual[User]()
			restored.Users = collection.NewSetEqual[User]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Admins.Length()).To(Equal(1))
			Expect(restored.Users.Length()).To(Equal(2))
		})
	})
})
