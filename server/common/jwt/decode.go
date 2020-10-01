package common_jwt

import (
	common_keycloak "primitivofr/owly/server/common/keycloak"
)

func ExtractUUIDfromJWT(token string) (string, error) {
	adminGuy, err := common_keycloak.InitAdmin()
	if err != nil {
		return "", err
	}
	return adminGuy.ExtractUUIDfromToken(token), nil
}
