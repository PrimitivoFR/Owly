package main

import (
	"log"
	authserver "primitivofr/owly/server-auth/auth_server"
)

func main() {
	log.Println("AuthServer started...")
	authserver.StartServer()
}
