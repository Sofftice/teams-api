package teamsapi

import (
	//"fmt"
	"net/http"
)

type Client struct {
	AccessToken  string
	RefreshToken string
	IdToken      string
	ClientInfo   string
	HTTPClient   *http.Client
	config       Config

	//Messages *MessagesService
	Auth *AuthService
}

// Generate a new client
func NewClient(accessToken, refreshToken, idToken, clientInfo string) *Client {
	c := &Client{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IdToken:      idToken,
		ClientInfo:   clientInfo,
		HTTPClient:   &http.Client{},
		config:       DefaultConfig(),
		Auth:         &AuthService{},
	}

	c.Auth.c = c
	return c
}
