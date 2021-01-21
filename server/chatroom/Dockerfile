ARG SERVER_COMMON

FROM ${SERVER_COMMON} as builder

ENV GO111MODULE=on
WORKDIR /server/chatroom
COPY ./go.mod .
RUN go mod download

COPY . . 

RUN  go build -o ./build/chatroom_server .

FROM golang as release 

COPY --from=builder /server/chatroom/build .

RUN chmod 700 ./chatroom_server
EXPOSE 50052
CMD [ "./chatroom_server" ]