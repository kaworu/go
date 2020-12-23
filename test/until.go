// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test until loops.

package main

func assertequal(is, shouldbe int, msg string) {
	if is != shouldbe {
		print("assertion fail ", msg, ", ", is, " should be ", shouldbe, "\n")
		panic(1)
	}
}

func main() {
	var i, sum int

	i = 0
	until false {
		i = i + 1
		if i > 5 {
			break
		}
	}
	assertequal(i, 6, "break")

	sum = 0
	until i := 0; i > 10 {
		sum = sum + i
		i++
	}
	assertequal(sum, 55, "all two")

	sum = 0
	until sum >= 100 {
		sum = sum + 9
	}
	assertequal(sum, 99+9, "only one")

	sum = 0
	until i := 0; i > 10 {
		if i % 2 == 0 {
			i++
			continue
		}
		sum = sum + i
		i++
	}
	assertequal(sum, 1+3+5+7+9, "continue")
}
