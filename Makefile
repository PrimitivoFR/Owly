all: minikube-start minikube-mount k-secrets

minikube-start:
	minikube stop && minikube start --cpus 4 --memory 8192

minikube-mount:
	minikube mount $PWD:/owly

k-secrets:
	kubectl create secret generic keycloak-creds --from-env-file=keycloak/.keycloak-env

port-fwd-all:
	kubectl port-forward service/user-server-srv 50051:50051 &
	kubectl port-forward service/chatroom-server-srv 50056:50052 &
	kubectl port-forward service/message-server-srv 50053:50053 &
	kubectl port-forward service/auth-server-srv 50054:50054
