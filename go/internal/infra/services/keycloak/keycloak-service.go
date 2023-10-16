package keycloakuser

import (
	"github.com/Nerzal/gocloak"
)

type keycloak struct {
	gocloak      gocloak.GoCloak
	clientId     string
	clientSecret string
	realm        string
}

func NewKeycloak(clientId, clientSecret, realm string) *keycloak {
	return &keycloak{
		gocloak:      gocloak.NewClient("http://localhost:8080"),
		clientId:     clientId,
		clientSecret: clientSecret,
		realm:        realm,
	}
}

func (k *keycloak) getToken() string {
	token, err := k.gocloak.LoginAdmin(k.clientId, k.clientSecret, k.realm)
	if err != nil {
		panic(err)
	}
	return *token.AccessToken
}

func (k *keycloak) Create(username, email, password string) (*gocloak.JWT, error) {
	return k.gocloak.CreateUser(k.getToken(), k.realm, gocloak.User{
		Username: &username,
		Email:    &email,
		Enabled:  &[]bool{true}[0],
	}, true)
}
