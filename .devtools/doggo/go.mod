module doggo

go 1.15

require (
	github.com/primitivofr/owly/server/common v0.0.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.35.0
)

replace github.com/primitivofr/owly/server/common => ../../server/common
