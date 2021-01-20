all: minikube-start k-secrets minikube-mount 

minikube-start:
	minikube stop
	minikube start --cpus 4 --memory 8192

minikube-mount:
	minikube mount ${PWD}:/owly

k-secrets:
	kubectl create secret generic keycloak-creds --from-env-file=keycloak/.keycloak-env

port-fwd-all:
	kubectl port-forward service/user-server-srv 50051:50051 &
	kubectl port-forward service/chatroom-server-srv 50056:50052 &
	kubectl port-forward service/message-server-srv 50053:50053 &
	kubectl port-forward service/auth-server-srv 50054:50054

test-build:
	docker-compose -f docker-compose.yml -f docker-compose.devandtest.yml -f docker-compose.test.yml build

test:
	docker-compose -f docker-compose.yml -f docker-compose.devandtest.yml -f docker-compose.test.yml up --exit-code-from server

clean-docker-cache:
	docker volume prune -f && docker system prune -f && docker container prune -f

dep-proto-grpc:
	export GO111MODULE=auto
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	go get -u google.golang.org/grpc

gen-proto-server:
	protoc protofiles/user.proto --go_out=plugins=grpc:./server/
	protoc protofiles/user.proto --go_out=plugins=grpc:./server/common/pb/

	protoc protofiles/chatroom.proto --go_out=plugins=grpc:./server/
	protoc protofiles/chatroom.proto --go_out=plugins=grpc:./server/common/pb/

	protoc protofiles/message.proto --go_out=plugins=grpc:./server/
	protoc protofiles/message.proto --go_out=plugins=grpc:./server/common/pb/

	protoc protofiles/auth.proto --go_out=plugins=grpc:./server/
	protoc protofiles/auth.proto --go_out=plugins=grpc:./server/common/pb/
	