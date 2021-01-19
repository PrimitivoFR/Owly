module primitivofr/owly/server-message

go 1.15

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20201229214741-2366c2514674
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.2
	github.com/ktr0731/evans v0.9.1 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rgamba/evtwebsocket v0.0.0-20181029234908-48b8cd9f8616
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	primitivofr/owly/server-common v0.0.0
)

replace primitivofr/owly/server-common => ../common
