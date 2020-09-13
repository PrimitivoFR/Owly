package main

import (
	"log"
	"primitivofr/owly/server/chatroom/chatroom_server"
	"primitivofr/owly/server/message/message_server"
	"primitivofr/owly/server/user/user_server"
)

func main() {
	log.Println("Hi from Owly Backend")

	// Need to think about the server-side architecture in production
	// Should we keep it like this ?
	// Or build a docker container for each gRPC server (more flexibility and control) (respects the micro-services pattern principle)

	go chatroom_server.StartServer()
	go message_server.StartServer()
	user_server.StartServer()
}
