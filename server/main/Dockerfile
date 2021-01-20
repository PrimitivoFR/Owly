# ---- Base ----
FROM golang

ENV GO111MODULE=on
WORKDIR /server
COPY ./go.mod .
RUN go mod download

COPY . . 

EXPOSE 50051
EXPOSE 50052
EXPOSE 50053
EXPOSE 50054

CMD [ "go", "run", "main.go" ]