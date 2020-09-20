package common_jwt

import (
	"primitivofr/owly/server/user/user_server/keycloak"
)

func DecodeJWT(token string) {

}
func ExtractUUIDfromJWT(token string) (string, error) {
	adminGuy, err := keycloak.InitAdmin()
	if err != nil {
		return "", err
	}
	return adminGuy.ExtractUUIDfromToken(token), nil
}
