ARG SERVER_COMMON

FROM ${SERVER_COMMON} as builder

ENV GO111MODULE=on
WORKDIR /server/message
COPY ./go.mod .
RUN go mod download

COPY . . 

RUN  go build -o ./build/message_server .

FROM golang as release 

COPY --from=builder /server/message/build .

RUN chmod 700 ./message_server
EXPOSE 50053
CMD [ "./message_server" ]