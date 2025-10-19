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

var _ = Describe("Set JSON", func() {
	Context("MarshalJSON", func() {
		It("marshals empty set to empty JSON array", func() {
			set := collection.NewSet[int]()
			data, err := json.Marshal(set)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal("[]"))
		})

		It("marshals set with primitive types (int)", func() {
			set := collection.NewSet(1, 2, 3)
			data, err := json.Marshal(set)
			Expect(err).NotTo(HaveOccurred())

			var result []int
			err = json.Unmarshal(data, &result)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(ConsistOf(1, 2, 3))
		})

		It("marshals set with string types", func() {
			set := collection.NewSet("apple", "banana", "cherry")
			data, err := json.Marshal(set)
			Expect(err).NotTo(HaveOccurred())

			var result []string
			err = json.Unmarshal(data, &result)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(ConsistOf("apple", "banana", "cherry"))
		})

		It("marshals set with complex types (struct)", func() {
			type Person struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}
			set := collection.NewSet(
				Person{Name: "Alice", Age: 30},
				Person{Name: "Bob", Age: 25},
			)
			data, err := json.Marshal(set)
			Expect(err).NotTo(HaveOccurred())

			var result []Person
			err = json.Unmarshal(data, &result)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(ConsistOf(
				Person{Name: "Alice", Age: 30},
				Person{Name: "Bob", Age: 25},
			))
		})
	})

	Context("UnmarshalJSON", func() {
		It("unmarshals empty JSON array to empty set", func() {
			set := collection.NewSet[int]()
			err := json.Unmarshal([]byte("[]"), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(0))
		})

		It("unmarshals JSON array with primitive types", func() {
			set := collection.NewSet[int]()
			err := json.Unmarshal([]byte("[1,2,3]"), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains(1)).To(BeTrue())
			Expect(set.Contains(2)).To(BeTrue())
			Expect(set.Contains(3)).To(BeTrue())
		})

		It("unmarshals JSON array with strings", func() {
			set := collection.NewSet[string]()
			err := json.Unmarshal([]byte(`["apple","banana","cherry"]`), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains("apple")).To(BeTrue())
			Expect(set.Contains("banana")).To(BeTrue())
			Expect(set.Contains("cherry")).To(BeTrue())
		})

		It("unmarshals JSON array with complex types", func() {
			type Person struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}
			set := collection.NewSet[Person]()
			jsonData := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`
			err := json.Unmarshal([]byte(jsonData), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(2))
			Expect(set.Contains(Person{Name: "Alice", Age: 30})).To(BeTrue())
			Expect(set.Contains(Person{Name: "Bob", Age: 25})).To(BeTrue())
		})

		It("handles duplicate elements in JSON array", func() {
			set := collection.NewSet[int]()
			err := json.Unmarshal([]byte("[1,2,3,2,1]"), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
		})

		It("replaces existing set contents", func() {
			set := collection.NewSet(1, 2, 3)
			err := json.Unmarshal([]byte("[4,5,6]"), set)
			Expect(err).NotTo(HaveOccurred())
			Expect(set.Length()).To(Equal(3))
			Expect(set.Contains(1)).To(BeFalse())
			Expect(set.Contains(4)).To(BeTrue())
			Expect(set.Contains(5)).To(BeTrue())
			Expect(set.Contains(6)).To(BeTrue())
		})
	})

	Context("Round-trip marshal/unmarshal", func() {
		It("preserves primitive types", func() {
			original := collection.NewSet(42, 7, 13)
			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			restored := collection.NewSet[int]()
			err = json.Unmarshal(data, restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Length()).To(Equal(original.Length()))
			for _, val := range original.Slice() {
				Expect(restored.Contains(val)).To(BeTrue())
			}
		})

		It("preserves complex structs", func() {
			type Config struct {
				Key   string `json:"key"`
				Value int    `json:"value"`
			}
			original := collection.NewSet(
				Config{Key: "timeout", Value: 30},
				Config{Key: "retries", Value: 3},
			)
			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			restored := collection.NewSet[Config]()
			err = json.Unmarshal(data, restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Length()).To(Equal(original.Length()))
			for _, val := range original.Slice() {
				Expect(restored.Contains(val)).To(BeTrue())
			}
		})
	})

	Context("Integration with structs containing sets", func() {
		It("marshals and unmarshals struct with Set field", func() {
			type Container struct {
				Name   string              `json:"name"`
				Values collection.Set[int] `json:"values"`
			}

			original := Container{
				Name:   "test",
				Values: collection.NewSet(1, 2, 3),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored Container
			restored.Values = collection.NewSet[int]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Name).To(Equal("test"))
			Expect(restored.Values.Length()).To(Equal(3))
			Expect(restored.Values.Contains(1)).To(BeTrue())
			Expect(restored.Values.Contains(2)).To(BeTrue())
			Expect(restored.Values.Contains(3)).To(BeTrue())
		})

		It("handles empty Set field in struct", func() {
			type Container struct {
				Name   string              `json:"name"`
				Values collection.Set[int] `json:"values"`
			}

			original := Container{
				Name:   "empty",
				Values: collection.NewSet[int](),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal(`{"name":"empty","values":[]}`))

			var restored Container
			restored.Values = collection.NewSet[int]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Name).To(Equal("empty"))
			Expect(restored.Values.Length()).To(Equal(0))
		})

		It("handles multiple Set fields in struct", func() {
			type MultiContainer struct {
				Integers collection.Set[int]    `json:"integers"`
				Strings  collection.Set[string] `json:"strings"`
				Active   bool                   `json:"active"`
			}

			original := MultiContainer{
				Integers: collection.NewSet(1, 2, 3),
				Strings:  collection.NewSet("a", "b", "c"),
				Active:   true,
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored MultiContainer
			restored.Integers = collection.NewSet[int]()
			restored.Strings = collection.NewSet[string]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Active).To(BeTrue())
			Expect(restored.Integers.Length()).To(Equal(3))
			Expect(restored.Strings.Length()).To(Equal(3))
			Expect(restored.Integers.Contains(1)).To(BeTrue())
			Expect(restored.Strings.Contains("a")).To(BeTrue())
		})

		It("handles nested structs with Set fields", func() {
			type Inner struct {
				Tags collection.Set[string] `json:"tags"`
			}
			type Outer struct {
				Name  string `json:"name"`
				Inner Inner  `json:"inner"`
			}

			original := Outer{
				Name: "parent",
				Inner: Inner{
					Tags: collection.NewSet("tag1", "tag2", "tag3"),
				},
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored Outer
			restored.Inner.Tags = collection.NewSet[string]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Name).To(Equal("parent"))
			Expect(restored.Inner.Tags.Length()).To(Equal(3))
			Expect(restored.Inner.Tags.Contains("tag1")).To(BeTrue())
		})

		It("handles Set of complex types in struct", func() {
			type Config struct {
				Key   string `json:"key"`
				Value int    `json:"value"`
			}
			type Application struct {
				Name    string                 `json:"name"`
				Configs collection.Set[Config] `json:"configs"`
			}

			original := Application{
				Name: "myapp",
				Configs: collection.NewSet(
					Config{Key: "timeout", Value: 30},
					Config{Key: "retries", Value: 3},
				),
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored Application
			restored.Configs = collection.NewSet[Config]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.Name).To(Equal("myapp"))
			Expect(restored.Configs.Length()).To(Equal(2))
			Expect(restored.Configs.Contains(Config{Key: "timeout", Value: 30})).To(BeTrue())
			Expect(restored.Configs.Contains(Config{Key: "retries", Value: 3})).To(BeTrue())
		})

		It("preserves other fields when Set field is present", func() {
			type ComplexStruct struct {
				ID       int                    `json:"id"`
				Name     string                 `json:"name"`
				Active   bool                   `json:"active"`
				Tags     collection.Set[string] `json:"tags"`
				Metadata map[string]string      `json:"metadata"`
			}

			original := ComplexStruct{
				ID:     123,
				Name:   "test",
				Active: true,
				Tags:   collection.NewSet("tag1", "tag2"),
				Metadata: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			}

			data, err := json.Marshal(original)
			Expect(err).NotTo(HaveOccurred())

			var restored ComplexStruct
			restored.Tags = collection.NewSet[string]()
			err = json.Unmarshal(data, &restored)
			Expect(err).NotTo(HaveOccurred())

			Expect(restored.ID).To(Equal(123))
			Expect(restored.Name).To(Equal("test"))
			Expect(restored.Active).To(BeTrue())
			Expect(restored.Tags.Length()).To(Equal(2))
			Expect(restored.Metadata).To(HaveLen(2))
			Expect(restored.Metadata["key1"]).To(Equal("value1"))
		})
	})
})
