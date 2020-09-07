package main

import (
	"log"
	"primitivofr/owly/server/user/user_server"
)

func main() {
	log.Println("Hi from Owly Backend")
	user_server.StartServer()
}
