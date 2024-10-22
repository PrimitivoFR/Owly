#!/bin/bash

# Golang
protoc protofiles/user.proto --go_out=plugins=grpc:./server/
protoc protofiles/user.proto --go_out=plugins=grpc:./server/common/pb/

protoc protofiles/chatroom.proto --go_out=plugins=grpc:./server/
protoc protofiles/chatroom.proto --go_out=plugins=grpc:./server/common/pb/

protoc protofiles/message.proto --go_out=plugins=grpc:./server/
protoc protofiles/message.proto --go_out=plugins=grpc:./server/common/pb/

protoc protofiles/auth.proto --go_out=plugins=grpc:./server/
protoc protofiles/auth.proto --go_out=plugins=grpc:./server/common/pb/