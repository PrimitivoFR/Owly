# ---- Base ----
FROM golang

ENV GO111MODULE=on
WORKDIR /server
COPY ./go.mod .
RUN go mod download

COPY . . 
