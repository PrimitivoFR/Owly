package keycloak

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v7"
)

type AdminGuy struct {
	Token  string
	Client gocloak.GoCloak
}

func NewAdmin(username string, password string, client gocloak.GoCloak) *AdminGuy {

	token, err := client.LoginAdmin(context.Background(), username, password, "master")
	if err != nil {
		log.Fatalf("Something wrong with the credentials or url: %v", err)
	}

	return &AdminGuy{Token: token.AccessToken, Client: client}
}

func (adminGuy *AdminGuy) CreateUser(user gocloak.User) (string, error) {

	ID, err := adminGuy.Client.CreateUser(context.Background(), adminGuy.Token, "OWLY", user)
	if err != nil {
		panic("Oh no!, failed to create user :(")
	}
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
