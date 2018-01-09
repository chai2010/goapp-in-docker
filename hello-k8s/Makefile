default:
	docker build -t hello-server .

hello:
	kubectl run hello-server --image=hello-server --port=8080
	kubectl expose deployment hello-server --type=NodePort

update-hello:
	kubectl set image deployment/hello-server hello-server=hello-server

start:
	minikube start
	eval $(minikube docker-env)
	minikube status

start-with-proxy:
	minikube start --docker-env ALL_PROXY=socks5://localhost:2080
	eval $(minikube docker-env)
	minikube status

start-with-mirror:
	minikube start --registry-mirror="https://registry.docker-cn.com"
	eval $(minikube docker-env)
	minikube status

stop:
	minikube stop
	eval $(minikube docker-env -u)

delete:
	minikube delete

pods:
	kubectl get pods

services:
	kubectl get services
