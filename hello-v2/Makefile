# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/mattn/go-sqlite3/issues/106#issuecomment-226280613
# docker run --rm -v $GOPATH:/go -w /go/src/path/to/app golang:latest go build -v

# http://blog.hashbangbash.com/2014/04/linking-golang-statically/

# docker run --rm -v $(shell go env GOPATH):/go -w /go/src/$(shell go list) golang:latest go build --ldflags '-extldflags "-static"' -o $(TARG).linux.exe $(SOURCES)

TARG=hello_v2

SOURCES=$(wildcard ./*.go)

default: docker
	docker run -it $(TARG)

docker: $(TARG).linux.exe Dockerfile
	docker build -t $(TARG) .

$(TARG).linux.exe: $(SOURCES) Makefile
	go fmt ./...
	docker run --rm -v $(shell go env GOPATH):/go -w /go/src/$(shell go list) golang:latest go build --ldflags '-extldflags "-static"' -o $(TARG).linux.exe $(SOURCES)

build:
	docker run --rm -v $(shell go env GOPATH):/go -w /go/src/$(shell go list) golang:latest go build -o $(TARG).linux.exe $(SOURCES)

fmt:
	go fmt ./...

run:
	go run $(SOURCES)

test:
	go test ./...

clean:
	-rm $(TARG).linux.exe
	-rm $(TARG)
