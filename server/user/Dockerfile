ARG SERVER_COMMON

FROM ${SERVER_COMMON} as builder

ENV GO111MODULE=on
WORKDIR /server/user
COPY ./go.mod .
RUN go mod download

COPY . . 

RUN  go build -o ./build/user_server .

FROM golang as release 

COPY --from=builder /server/user/build .

RUN chmod 700 ./user_server
EXPOSE 50051
CMD [ "./user_server" ]