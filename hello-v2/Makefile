# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/mattn/go-sqlite3/issues/106#issuecomment-226280613
# http://blog.hashbangbash.com/2014/04/linking-golang-statically/

TARG=hello_v2

SOURCES=$(wildcard ./*.go)

GO_LDFLAGS=--ldflags '-w -s -extldflags "-static"'
DOCKER_VOLUME=-v $(shell go env GOPATH):/go -w /go/src/$(shell go list)

default: docker
	docker run -it $(TARG)

docker: $(TARG).linux.exe Dockerfile
	docker build -t $(TARG) .

$(TARG).linux.exe: $(SOURCES) Makefile
	go fmt ./...
	docker run --rm $(DOCKER_VOLUME) golang:latest go build $(GO_LDFLAGS) -o $(TARG).linux.exe $(SOURCES)

build:
	docker run --rm $(DOCKER_VOLUME) golang:latest go build $(GO_LDFLAGS) -o $(TARG).linux.exe $(SOURCES)

fmt:
	go fmt ./...

run:
	go run $(SOURCES)

test:
	go test ./...

clean:
	-rm $(TARG).linux.exe
	-rm $(TARG)
