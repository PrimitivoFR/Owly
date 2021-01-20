module github.com/primitivofr/owly/server/chatroom

go 1.15

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2
	github.com/ktr0731/evans v0.9.1 // indirect
	github.com/primitivofr/owly/server/auth v0.0.0
	github.com/primitivofr/owly/server/common v0.0.0
	github.com/primitivofr/owly/server/message v0.0.0
	github.com/primitivofr/owly/server/user v0.0.0
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.4.4
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/primitivofr/owly/server/auth => ../auth
	github.com/primitivofr/owly/server/chatroom => ../chatroom
	github.com/primitivofr/owly/server/common => ../common
	github.com/primitivofr/owly/server/message => ../message
	github.com/primitivofr/owly/server/user => ../user

)
