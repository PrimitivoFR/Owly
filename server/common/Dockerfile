FROM golang as builder

ENV GO111MODULE=on
WORKDIR /server/common
COPY ./go.mod .
RUN go mod download

COPY . .