// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/yupeng9/go.lib/gosh"
)

func main() {
	gosh.InitChildMain()
	fmt.Println("Hello, world!")
}
