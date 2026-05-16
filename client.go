package teamsapi

import (
	//"fmt"
	"net/http"
)

type Client struct {
	Access_token  string
	Refresh_token string
	Id_token      string
	Client_info   string
	HTTPClient    *http.Client
	config        Config

	//Messages *MessagesService
	Auth *AuthService
}

// Generate a new client
func NewClient(accessToken, refreshToken, idToken, clientInfo string) *Client {

	c := &Client{
		Access_token:  accessToken,
		Refresh_token: refreshToken,
		Id_token:      idToken,
		Client_info:   clientInfo,
		HTTPClient:    &http.Client{},
		config:        DefaultConfig(),
		Auth:          &AuthService{},
	}

	c.Auth.c = c
	return c
}
