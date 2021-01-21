module github.com/primitivofr/owly/server/auth

go 1.15

require (
	github.com/Nerzal/gocloak/v7 v7.11.0
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/ktr0731/evans v0.9.1 // indirect
	github.com/primitivofr/owly/server/common v0.0.0
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.4.4 // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/primitivofr/owly/server/common => ../common
