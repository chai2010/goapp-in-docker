# hello-v2

```
localhost:hello-v2 chai$ make
go fmt ./...
CGO_ENABLED=1 GOOS=linux go build -o hello_v2.linux.exe ./hello.go
# runtime/cgo
ld: unknown option: --build-id=none
clang: error: linker command failed with exit code 1 (use -v to see invocation)
make: *** [hello_v2.linux.exe] Error 2
localhost:hello-v2 chai$ 
```

在 macOS 上 cgo 交叉编译错误!
