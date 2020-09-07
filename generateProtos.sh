#!/bin/bash

# Golang
protoc protofiles/user.proto --go_out=plugins=grpc:./server/user/userpb