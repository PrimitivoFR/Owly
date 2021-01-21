package interceptors

import (
	"log"

	common_keycloak "github.com/primitivofr/owly/server/common/keycloak"
)

func checkTokenInterceptor(token string) (bool, error) {

	adminGuy, err := common_keycloak.InitAdmin()
	if err != nil {
		log.Println(err)
		return false, err
	}
	return adminGuy.VerifyToken(token)
}
