package main

import (
	"log"
	"primitivofr/owly/server/chatroom/chatroom_server"
	"primitivofr/owly/server/user/user_server"
)

func main() {
	log.Println("Hi from Owly Backend")
	go chatroom_server.StartServer()
	user_server.StartServer()
}
