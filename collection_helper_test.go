// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection_test

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Age       int
	Groups    *string
}

func (u User) Equal(user User) bool {
	return u == user
}

func (u User) HashCode() string {
	return fmt.Sprintf("%#v", u)
}
