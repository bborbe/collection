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

var _ = Describe("ChannelFnCount", func() {
	It("counts zero elements from empty channel", func() {
		ctx := context.Background()
		count, err := collection.ChannelFnCount(ctx, func(ctx context.Context, ch chan<- string) error {
			return nil // Send nothing
		})

		Expect(err).To(BeNil())
		Expect(count).To(Equal(0))
	})

	It("counts single element", func() {
		ctx := context.Background()
		count, err := collection.ChannelFnCount(ctx, func(ctx context.Context, ch chan<- string) error {
			ch <- "test"
			return nil
		})

		Expect(err).To(BeNil())
		Expect(count).To(Equal(1))
	})

	It("counts multiple elements", func() {
		ctx := context.Background()
		count, err := collection.ChannelFnCount(ctx, func(ctx context.Context, ch chan<- string) error {
			for i := 0; i < 5; i++ {
				ch <- fmt.Sprintf("item-%d", i)
			}
			return nil
		})

		Expect(err).To(BeNil())
		Expect(count).To(Equal(5))
	})

	It("counts integers", func() {
		ctx := context.Background()
		count, err := collection.ChannelFnCount(ctx, func(ctx context.Context, ch chan<- int) error {
			for i := 1; i <= 10; i++ {
				ch <- i
			}
			return nil
		})

		Expect(err).To(BeNil())
		Expect(count).To(Equal(10))
	})

	It("handles error from producer function", func() {
		ctx := context.Background()
		count, err := collection.ChannelFnCount(ctx, func(ctx context.Context, ch chan<- string) error {
			ch <- "before error"
			return fmt.Errorf("producer error")
		})

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("count channel failed"))
		Expect(err.Error()).To(ContainSubstring("producer error"))
		Expect(count).To(Equal(-1))
	})

	It("handles context cancellation", func() {
		ctx, cancel := context.WithCancel(context.Background())

		done := make(chan struct{})
		go func() {
			defer close(done)
			count, err := collection.ChannelFnCount(ctx, func(ctx context.Context, ch chan<- int) error {
				for i := 0; i < 1000; i++ {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case ch <- i:
					}
				}
				return nil
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("context canceled"))
			Expect(count).To(Equal(-1))
		}()

		cancel() // Cancel immediately
		Eventually(done).Should(BeClosed())
	})

	It("counts custom types", func() {
		ctx := context.Background()
		count, err := collection.ChannelFnCount(ctx, func(ctx context.Context, ch chan<- User) error {
			ch <- User{Firstname: "Alice", Age: 25}
			ch <- User{Firstname: "Bob", Age: 30}
			ch <- User{Firstname: "Charlie", Age: 35}
			return nil
		})

		Expect(err).To(BeNil())
		Expect(count).To(Equal(3))
	})
})
