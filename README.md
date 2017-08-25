# 运行Go应用的最小Docker环境

```
chai-mba:goapp-in-docker chai$ make
docker build -t go_app .
Sending build context to Docker daemon  1.105MB
Step 1/3 : FROM scratch
 ---> 
Step 2/3 : ADD go_app /
 ---> Using cache
 ---> a926d6467117
Step 3/3 : CMD /go_app
 ---> Using cache
 ---> 33dff05e5dce
Successfully built 33dff05e5dce
Successfully tagged go_app:latest
docker run -it go_app
hello docker!!!
chai-mba:goapp-in-docker chai$ 
```
