module primitivofr/owly/server-user

go 1.15

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.2
	github.com/ktr0731/evans v0.9.1 // indirect
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	primitivofr/owly/server-common v0.0.0
)

replace primitivofr/owly/server-common => ../common
