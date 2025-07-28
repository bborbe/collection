// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

// Map applies the given function to each element in the slice.
// If any function call returns an error, Map stops and returns that error.
func Map[T any](list []T, fn func(value T) error) error {
	for _, element := range list {
		if err := fn(element); err != nil {
			return err
		}
	}
	return nil
}
