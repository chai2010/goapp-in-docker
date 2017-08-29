// Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://stackoverflow.com/questions/32309030/go-1-5-cross-compile-using-cgo-on-os-x-to-linux-and-windows

package main

/*
#cgo linux LDFLAGS: -lm -lstdc++

#include <math.h>

static int answerToLife() {
	return 42;
}
*/
import "C"
import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	flagEnableHttp = flag.Bool("http", false, "enable http server")
)

func main() {
	flag.Parse()

	fmt.Println("hello v2!!!")
	fmt.Println("answerToLife:", C.answerToLife())
	fmt.Println("sqrt(9527):", float64(C.sqrt(9527)))

	if *flagEnableHttp {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Hello, 世界")
			fmt.Fprintf(w, "Hello, 世界")
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
