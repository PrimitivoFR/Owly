package interceptors

import (
	"log"
	common_keycloak "primitivofr/owly/server/common/keycloak"
)

func checkTokenInterceptor(token string) (bool, error) {
	log.Println(token)
	adminGuy, err := common_keycloak.InitAdmin()
	if err != nil {
		log.Println(err)
	}
	return adminGuy.VerifyToken(token)
}
