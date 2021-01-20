package main

import (
	"log"
	authserver "github.com/primitivofr/owly/server/auth/auth_server"
)

func main() {
	log.Println("AuthServer started...")
	authserver.StartServer()
}
