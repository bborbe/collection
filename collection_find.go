// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import stderrors "errors"

// NotFoundError is returned when an element is not found in a collection.
var NotFoundError = stderrors.New("not found")

// Find returns a pointer to the first element in the slice that satisfies the predicate function.
// If no element is found, it returns nil and NotFoundError.
func Find[T any](list []T, match func(value T) bool) (*T, error) {
	for _, e := range list {
		if match(e) {
			return &e, nil
		}
	}
	return nil, NotFoundError
}
