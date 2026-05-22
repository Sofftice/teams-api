package teamsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AuthService struct {
	c *Client
}

type RefreshAccessTokenResponse struct {
	TokenType             string `json:"token_type"`
	Scope                 string `json:"scope"`
	ExpiresIn             int    `json:"expires_in"`
	ExtInspiresIn         int    `json:"ext_expires_int"`
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	IdToken               string `json:"id_token"`
	ClientInfo            string `json:"client_info"`
}

func (a *AuthService) RefreshAccessToken(ctx context.Context) (RefreshAccessTokenResponse, error) {
	var body RefreshAccessTokenResponse

	data := url.Values{}
	data.Set("client_id", a.c.config.ClientId)
	data.Set("scope", "https://api.spaces.skype.com/.default openid profile offline_access")
	data.Set("grant_type", "refresh_token")
	data.Set("client_info", "1")
	data.Set("refresh_token", a.c.RefreshToken)
	data.Set("claims", "{\"access_token\":{\"xms_cc\":{\"values\":[\"CP1\"]}}}")

	req, err := http.NewRequestWithContext(ctx, "POST", "https://login.microsoftonline.com/fe0ca031-7f0a-4b82-8eea-e46b799fdc86/oauth2/v2.0/token?client-request-id=Core-69617e44-2977-4af4-9277-2f11eec04c78", strings.NewReader(data.Encode()))
	if err != nil {
		return body, fmt.Errorf("unable to create request: %s", err.Error())
	}

	resp, err := a.c.doRequest(ctx, req)
	if err != nil {
		return body, fmt.Errorf("request failed: %s", err.Error())
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return body, fmt.Errorf("failed to decode response body: %s", err.Error())
	}

	a.c.AccessToken = body.AccessToken
	a.c.RefreshToken = body.RefreshToken

	return body, nil
}
