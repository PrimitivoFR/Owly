module doggo

go 1.15

require (
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.31.1
	primitivofr/owly/server v0.0.0-00010101000000-000000000000
)

replace primitivofr/owly/server => ../../server
