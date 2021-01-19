module primitivofr/owly/server-chatroom

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	go.mongodb.org/mongo-driver v1.4.4
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	primitivofr/owly/server-common v0.0.0
)

replace primitivofr/owly/server-common => ../common