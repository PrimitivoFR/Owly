ARG SERVER_COMMON

FROM ${SERVER_COMMON} as builder

ENV GO111MODULE=on
WORKDIR /server/auth
COPY ./go.mod .
RUN go mod download

COPY . . 

RUN  go build -o ./build/auth_server .

FROM golang as release 

COPY --from=builder /server/auth/build .

RUN chmod 700 ./auth_server
EXPOSE 50054
CMD [ "./auth_server" ]