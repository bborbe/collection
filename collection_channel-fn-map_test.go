// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import (
	"context"
	"fmt"
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/collection"
)

var _ = Describe("ChannelFnMap", func() {
	It("handles empty channel", func() {
		ctx := context.Background()
		var mapCalled bool

		err := collection.ChannelFnMap(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				return nil // Send nothing
			},
			func(ctx context.Context, s string) error {
				mapCalled = true
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(mapCalled).To(BeFalse())
	})

	It("processes single element", func() {
		ctx := context.Background()
		var processed []string

		err := collection.ChannelFnMap(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				ch <- "test"
				return nil
			},
			func(ctx context.Context, s string) error {
				processed = append(processed, s)
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(processed).To(Equal([]string{"test"}))
	})

	It("processes multiple elements", func() {
		ctx := context.Background()
		var processed []string
		var mu sync.Mutex

		err := collection.ChannelFnMap(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				ch <- "first"
				ch <- "second"
				ch <- "third"
				return nil
			},
			func(ctx context.Context, s string) error {
				mu.Lock()
				processed = append(processed, s)
				mu.Unlock()
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(processed).To(HaveLen(3))
		Expect(processed).To(ContainElements("first", "second", "third"))
	})

	It("handles error from producer function", func() {
		ctx := context.Background()

		err := collection.ChannelFnMap(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				return fmt.Errorf("producer error")
			},
			func(ctx context.Context, s string) error {
				return nil
			},
		)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("map channel failed"))
		Expect(err.Error()).To(ContainSubstring("producer error"))
	})

	It("handles error from map function", func() {
		ctx := context.Background()

		err := collection.ChannelFnMap(
			ctx,
			func(ctx context.Context, ch chan<- string) error {
				ch <- "test"
				return nil
			},
			func(ctx context.Context, s string) error {
				return fmt.Errorf("map error for %s", s)
			},
		)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("map channel failed"))
		Expect(err.Error()).To(ContainSubstring("map failed"))
		Expect(err.Error()).To(ContainSubstring("map error for test"))
	})

	It("handles context cancellation", func() {
		ctx, cancel := context.WithCancel(context.Background())

		done := make(chan struct{})
		go func() {
			defer close(done)
			err := collection.ChannelFnMap(
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
				func(ctx context.Context, i int) error {
					return nil
				},
			)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("context canceled"))
		}()

		cancel() // Cancel immediately
		Eventually(done).Should(BeClosed())
	})

	It("processes custom types", func() {
		ctx := context.Background()
		var processedNames []string
		var mu sync.Mutex

		users := []User{
			{Firstname: "Alice", Age: 25},
			{Firstname: "Bob", Age: 30},
			{Firstname: "Charlie", Age: 35},
		}

		err := collection.ChannelFnMap(
			ctx,
			func(ctx context.Context, ch chan<- User) error {
				for _, user := range users {
					ch <- user
				}
				return nil
			},
			func(ctx context.Context, user User) error {
				mu.Lock()
				processedNames = append(processedNames, user.Firstname)
				mu.Unlock()
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(processedNames).To(HaveLen(3))
		Expect(processedNames).To(ContainElements("Alice", "Bob", "Charlie"))
	})

	It("handles concurrent processing", func() {
		ctx := context.Background()
		var counter int
		var mu sync.Mutex

		err := collection.ChannelFnMap(
			ctx,
			func(ctx context.Context, ch chan<- int) error {
				for i := 0; i < 100; i++ {
					ch <- i
				}
				return nil
			},
			func(ctx context.Context, i int) error {
				mu.Lock()
				counter++
				mu.Unlock()
				return nil
			},
		)

		Expect(err).To(BeNil())
		Expect(counter).To(Equal(100))
	})
})
