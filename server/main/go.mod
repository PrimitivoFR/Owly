module github.com/primitivofr/owly/server/main

go 1.15

require (
	github.com/Nerzal/gocloak/v7 v7.11.0
	github.com/boltdb/bolt v1.3.1
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20201229214741-2366c2514674
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2
	github.com/jhump/protoreflect v1.7.0
	github.com/joho/godotenv v1.3.0
	github.com/ktr0731/evans v0.9.1
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/pkg/errors v0.9.1
	github.com/primitivofr/owly/server/auth v0.0.0
	github.com/primitivofr/owly/server/chatroom v0.0.0
	github.com/primitivofr/owly/server/common v0.0.0
	github.com/primitivofr/owly/server/message v0.0.0
	github.com/primitivofr/owly/server/user v0.0.0
	github.com/rgamba/evtwebsocket v0.0.0-20181029234908-48b8cd9f8616
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
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
