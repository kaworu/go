// errorcheck

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func x() {
}

func main() {
	until {  // ERROR "missing condition"
	}

	until x(); {  // ERROR "missing condition"
	}
}
