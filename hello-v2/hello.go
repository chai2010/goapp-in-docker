// Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://stackoverflow.com/questions/32309030/go-1-5-cross-compile-using-cgo-on-os-x-to-linux-and-windows

package main

/*
static int answerToLife() {
	return 42;
}
*/
import "C"
import (
	"fmt"
)

func main() {
	fmt.Println("hello v2!!!", C.answerToLife())
}
