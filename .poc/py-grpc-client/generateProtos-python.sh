#!/bin/bash

# Python protofile generation
cd ../../protofiles
python -m grpc_tools.protoc -I. --python_out=../.poc/py-grpc-client --grpc_python_out=../.poc/py-grpc-client user.proto
python -m grpc_tools.protoc -I. --python_out=../.poc/py-grpc-client --grpc_python_out=../.poc/py-grpc-client chatroom.proto
python -m grpc_tools.protoc -I. --python_out=../.poc/py-grpc-client --grpc_python_out=../.poc/py-grpc-client message.proto
python -m grpc_tools.protoc -I. --python_out=../.poc/py-grpc-client --grpc_python_out=../.poc/py-grpc-client auth.proto