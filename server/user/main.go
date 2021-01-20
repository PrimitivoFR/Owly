package main

import (
	"log"
	"github.com/primitivofr/owly/server/user/user_server"
)

func main() {
	log.Println("UserServer started...")
	user_server.StartServer()
}
