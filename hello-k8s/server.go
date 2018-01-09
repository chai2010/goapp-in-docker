// Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	flagEnableHttp = flag.Bool("http", true, "enable http server")
)

func main() {
	flag.Parse()

	fmt.Println("hello v2!!!")

	if *flagEnableHttp {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Hello, 世界")
			fmt.Fprintf(w, "Hello, 世界")
		})
		log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
	}
}
