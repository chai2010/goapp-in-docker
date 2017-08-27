# hello-v1

```
localhost:hello-v1 chai$ make
CGO_ENABLED=0 GOOS=linux go build -o hello_v1 .
docker build -t hello_v1 .
Sending build context to Docker daemon  1.049MB
Step 1/3 : FROM scratch
 ---> 
Step 2/3 : ADD hello_v1 /
 ---> Using cache
 ---> 8d38b64b3c9d
Step 3/3 : CMD /hello_v1
 ---> Using cache
 ---> 65f64ef57736
Successfully built 65f64ef57736
Successfully tagged hello_v1:latest
docker run -it hello_v1
hello docker v1!!!
localhost:hello-v1 chai$
```
