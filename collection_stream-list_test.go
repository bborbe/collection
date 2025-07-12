// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("StreamList", func() {
	It("streams empty list", func() {
		ctx := context.Background()
		ch := make(chan string, 10)

		err := collection.StreamList(ctx, []string{}, ch)
		Expect(err).To(BeNil())

		close(ch)
		var result []string
		for item := range ch {
			result = append(result, item)
		}
		Expect(result).To(BeEmpty())
	})

	It("streams single element", func() {
		ctx := context.Background()
		ch := make(chan string, 10)

		err := collection.StreamList(ctx, []string{"test"}, ch)
		Expect(err).To(BeNil())

		close(ch)
		var result []string
		for item := range ch {
			result = append(result, item)
		}
		Expect(result).To(Equal([]string{"test"}))
	})

	It("streams multiple elements", func() {
		ctx := context.Background()
		ch := make(chan string, 10)

		input := []string{"a", "b", "c", "d"}
		err := collection.StreamList(ctx, input, ch)
		Expect(err).To(BeNil())

		close(ch)
		var result []string
		for item := range ch {
			result = append(result, item)
		}
		Expect(result).To(Equal(input))
	})

	It("streams integers", func() {
		ctx := context.Background()
		ch := make(chan int, 10)

		input := []int{1, 2, 3, 4, 5}
		err := collection.StreamList(ctx, input, ch)
		Expect(err).To(BeNil())

		close(ch)
		var result []int
		for item := range ch {
			result = append(result, item)
		}
		Expect(result).To(Equal(input))
	})

	It("handles context cancellation", func() {
		ctx, cancel := context.WithCancel(context.Background())
		ch := make(chan int, 1) // Small buffer to force blocking

		// Start streaming in goroutine
		done := make(chan error, 1)
		go func() {
			done <- collection.StreamList(ctx, []int{1, 2, 3, 4, 5}, ch)
		}()

		// Cancel context after short delay
		time.Sleep(10 * time.Millisecond)
		cancel()

		// Should get context cancelled error
		select {
		case err := <-done:
			Expect(err).To(Equal(context.Canceled))
		case <-time.After(time.Second):
			Fail("StreamList did not respond to context cancellation")
		}
	})

	It("streams custom types", func() {
		ctx := context.Background()
		ch := make(chan User, 10)

		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
		}

		err := collection.StreamList(ctx, users, ch)
		Expect(err).To(BeNil())

		close(ch)
		var result []User
		for user := range ch {
			result = append(result, user)
		}
		Expect(result).To(Equal(users))
	})

	It("works with unbuffered channel", func() {
		ctx := context.Background()
		ch := make(chan string) // Unbuffered

		// Read from channel in goroutine
		result := make([]string, 0)
		done := make(chan bool)
		go func() {
			defer func() { done <- true }()
			for item := range ch {
				result = append(result, item)
			}
		}()

		// Stream data
		err := collection.StreamList(ctx, []string{"a", "b"}, ch)
		Expect(err).To(BeNil())
		close(ch)

		<-done
		Expect(result).To(Equal([]string{"a", "b"}))
	})
})
