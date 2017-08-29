# hello-v2

```
localhost:hello-v2 chai$ make
go fmt ./...
docker run --rm -v /Users/chai/go:/go -w /go/src/github.com/chai2010/goapp-in-docker/hello-v2 golang:latest go build --ldflags '-extldflags "-static"' -o hello_v2.linux.exe ./hello.go
# command-line-arguments
/tmp/go-link-324585367/000002.o: In function `_cgo_b0c710f30cfd_C2func_getaddrinfo':
/tmp/go-build/net/_obj/cgo-gcc-prolog:46: warning: Using 'getaddrinfo' in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
docker build -t hello_v2 .
Sending build context to Docker daemon  7.605MB
Step 1/3 : FROM scratch
 --->
Step 2/3 : ADD hello_v2.linux.exe /
 ---> 09c2a9d614e6
Removing intermediate container abdea7399e29
Step 3/3 : CMD /hello_v2.linux.exe
 ---> Running in deb5c697586b
 ---> cad602d141ef
Removing intermediate container deb5c697586b
Successfully built cad602d141ef
Successfully tagged hello_v2:latest
docker run -it hello_v2
hello v2!!!
answerToLife: 42
sqrt(9527): 97.60635225229964
localhost:hello-v2 chai$
```
