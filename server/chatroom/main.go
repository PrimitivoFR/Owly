package main

import (
	"log"
	"github.com/primitivofr/owly/server/chatroom/chatroom_server"
)

func main() {
	log.Println("ChatroomServer started...")
	chatroom_server.StartServer()
}
