#!/bin/bash

docker build -t init_server_image . && \ 
docker build -t auth_server ./auth && \
docker build -t chatroom_server ./chatroom && \
docker build -t user_server ./user && \
docker build -t message_server ./message