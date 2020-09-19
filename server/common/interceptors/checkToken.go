package interceptors

import (
	"log"
	"primitivofr/owly/server/user/user_server/keycloak"
)

func checkTokenInterceptor(token string) (bool, error) {
	log.Println(token)
	adminGuy, err := keycloak.InitAdmin()
	if err != nil {
		log.Println(err)
	}
	return adminGuy.VerifyToken(token)
}
