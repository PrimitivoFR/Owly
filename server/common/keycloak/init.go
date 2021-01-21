package common_keycloak

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Nerzal/gocloak/v7"
	"github.com/pkg/errors"
)

var keycloak_admin_user string
var keycloak_admin_pass string
var keycloak_hostname string

func loadEnvVars() {
	log.Println("Loading env vars")
	keycloak_admin_user = os.Getenv("KEYCLOAK_USER")
	keycloak_admin_pass = os.Getenv("KEYCLOAK_PASSWORD")
	keycloak_hostname = os.Getenv("KEYCLOAK_HOSTNAME")

}

func checkConnection() {
	for {
		res, err := http.Get("http://" + keycloak_hostname + ":8080/auth/")
		if err == nil && res != nil {
			if res.StatusCode == 200 {
				log.Println("Keycloak instance up, ready and reachable !")
				break
			}
		}
		time.Sleep(2 * time.Second)

		log.Println("Waiting for keycloak instance @ keycloak-srv...")
		log.Println(err)
	}
}

func InitAdmin() (*AdminGuy, error) {
	loadEnvVars()
	checkConnection()
	client := gocloak.NewClient("http://" + keycloak_hostname + ":8080/")
	if client == nil {
		log.Fatal()
		return nil, errors.Errorf("Couldn't connect to Keycloak instance")
	}

	adminGuy, err := NewAdmin(keycloak_admin_user, keycloak_admin_pass, client)

	return adminGuy, err

}
