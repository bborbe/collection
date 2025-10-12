// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("ChannelFnList", func() {
	It("returns empty list from empty channel", func() {
		ctx := context.Background()
		result, err := collection.ChannelFnList(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				return nil // Send nothing
			},
		)

		Expect(err).To(BeNil())
		Expect(result).To(BeEmpty())
	})

	It("collects single element", func() {
		ctx := context.Background()
		result, err := collection.ChannelFnList(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				ch <- "test"
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(result).To(Equal([]string{"test"}))
	})

	It("collects multiple elements in order", func() {
		ctx := context.Background()
		result, err := collection.ChannelFnList(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				ch <- "first"
				ch <- "second"
				ch <- "third"
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(result).To(Equal([]string{"first", "second", "third"}))
	})

	It("collects integers", func() {
		ctx := context.Background()
		result, err := collection.ChannelFnList(
			ctx,
			func(ctx context.Context, ch chan<- int) error {
				for i := 1; i <= 5; i++ {
					ch <- i
				}
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(result).To(Equal([]int{1, 2, 3, 4, 5}))
	})

	It("handles error from producer function", func() {
		ctx := context.Background()
		result, err := collection.ChannelFnList(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				ch <- "before error"
				return fmt.Errorf("producer error")
			},
		)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("convert channel to list failed"))
		Expect(err.Error()).To(ContainSubstring("producer error"))
		Expect(result).To(BeNil())
	})

	It("handles context cancellation", func() {
		ctx, cancel := context.WithCancel(context.Background())

		done := make(chan struct{})
		go func() {
			defer close(done)
			result, err := collection.ChannelFnList(
				ctx,
				func(ctx context.Context, ch chan<- int) error {
					for i := 0; i < 1000; i++ {
						select {
						case <-ctx.Done():
							return ctx.Err()
						case ch <- i:
						}
					}
					return nil
				},
			)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("context canceled"))
			Expect(result).To(BeNil())
		}()

		cancel() // Cancel immediately
		Eventually(done).Should(BeClosed())
	})

	It("collects custom types", func() {
		ctx := context.Background()
		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Charlie", Age: 35},
		}

		result, err := collection.ChannelFnList(
			ctx,
			func(ctx context.Context, ch chan<- User) error {
				for _, user := range users {
					ch <- user
				}
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(result).To(Equal(users))
	})

	It("handles large number of elements", func() {
		ctx := context.Background()
		expected := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			expected[i] = i
		}

		result, err := collection.ChannelFnList(
			ctx,
			func(ctx context.Context, ch chan<- int) error {
				for i := 0; i < 1000; i++ {
					ch <- i
				}
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(result).To(Equal(expected))
	})
})
