#!/bin/bash

docker build -t init_server_image . && \ 
docker build -t owly-server-auth ./auth && \
docker build -t owly-server-chatoom ./chatroom && \
docker build -t owly-server-user ./user && \
docker build -t owly-server-message ./message