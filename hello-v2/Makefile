# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

TARG=hello_v2

SOURCES=$(wildcard ./*.go)

default: docker
	docker run -it $(TARG)

docker: $(TARG).linux.exe Dockerfile
	docker build -t $(TARG) .

$(TARG).linux.exe: $(SOURCES) Makefile
	go fmt ./...
	CGO_ENABLED=1 GOOS=linux go build -o $(TARG).linux.exe $(SOURCES)

fmt:
	go fmt ./...

run:
	go run $(SOURCES)

test:
	go test ./...

clean:
	-rm $(TARG)
