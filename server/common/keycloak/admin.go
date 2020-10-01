package common_keycloak

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v7"
	"github.com/dgrijalva/jwt-go/v4"
)

type AdminGuy struct {
	Token  string
	Client gocloak.GoCloak
}

func NewAdmin(username string, password string, client gocloak.GoCloak) (*AdminGuy, error) {

	token, err := client.LoginAdmin(context.Background(), username, password, "master")
	if err != nil {
		return nil, err
	}

	return &AdminGuy{Token: token.AccessToken, Client: client}, nil
}

func (adminGuy *AdminGuy) CreateUser(user gocloak.User) (string, error) {

	ID, err := adminGuy.Client.CreateUser(context.Background(), adminGuy.Token, "OWLY", user)
	return ID, err

}

func (adminGuy *AdminGuy) SetUserPassword(userId string, password string, tempo bool) error {
	return adminGuy.Client.SetPassword(context.Background(), adminGuy.Token, userId, "OWLY", password, tempo)
}

func (adminGuy *AdminGuy) SearchUserByUsername(username string) ([]*gocloak.User, error) {
	searchParams := gocloak.GetUsersParams{
		Search: &username,
	}
	return adminGuy.Client.GetUsers(context.Background(), adminGuy.Token, "OWLY", searchParams)
}

func (adminGuy *AdminGuy) GetClientSecret(clientName string) (string, error) {

	var cli_id string

	res1, err := adminGuy.GetAllClients()
	if err != nil {
		return "", err
	}

	for _, cli := range res1 {
		if clientID := cli.ClientID; *clientID == clientName {
			cli_id = *cli.ID
		}
	}

	res, err := adminGuy.Client.GetClientSecret(context.Background(), adminGuy.Token, "OWLY", cli_id)
	if err != nil {
		log.Printf("Error @ GetClientSecret: %v", err)
	}

	return *res.Value, err
}

func (adminGuy *AdminGuy) GetAllClients() ([]*gocloak.Client, error) {
	bill := true
	return adminGuy.Client.GetClients(context.Background(), adminGuy.Token, "OWLY", gocloak.GetClientsParams{ViewableOnly: &bill})
}
func (adminGuy *AdminGuy) GetClientId(clientName string) (string, error) {
	var cli_id string

	res1, err := adminGuy.GetAllClients()
	if err != nil {
		return "", err
	}

	for _, cli := range res1 {
		if clientID := cli.ClientID; *clientID == clientName {
			cli_id = *cli.ID
		}
	}
	return cli_id, nil
}

func (adminGuy *AdminGuy) VerifyToken(token string) (bool, error) {

	cli_secret, err := adminGuy.GetClientSecret("owlycli")
	if err != nil {
		return false, err
	}

	res, err := adminGuy.Client.RetrospectToken(context.Background(), token, "owlycli", cli_secret, "OWLY")
	if err != nil {
		log.Printf("Fails at adminGuy.VerifyToken %v", err)
		return false, err
	}
	return *res.Active, nil
}

func (adminGuy *AdminGuy) ExtractUUIDfromToken(token string) string {
	decoded, _, _ := adminGuy.Client.DecodeAccessToken(context.Background(), token, "OWLY", "")
	claims := *decoded.Claims.(*jwt.MapClaims)
	return claims["sub"].(string)
}
