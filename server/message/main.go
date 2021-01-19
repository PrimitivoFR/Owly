package main

import (
	"log"
	"primitivofr/owly/server/message/message_server"
)

func main() {
	log.Println("MessageServer started...")
	message_server.StartServer()

}
