default: run

docker: go_app
	docker build -t go_app .

go_app: hello.go Makefile
	CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -installsuffix netgo -o go_app .

run: docker
	docker run -it go_app

clean:
	-rm go_app
