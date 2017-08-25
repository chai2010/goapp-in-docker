default: run

docker: go_app
	docker build -t go_app .

go_app:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go_app .

run: docker
	docker run -it go_app

clean:
	-rm go_app
