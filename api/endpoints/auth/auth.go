package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adalessa/qatar/pkg/endpoint"
)

type AuthEndpoint struct {
	endpoint endpoint.Endpoint
}

const LoginURI = "api/v1/user/login"

func NewEndpoint(domain string) *AuthEndpoint {
	return &AuthEndpoint{
		endpoint: endpoint.New(domain),
	}
}

func (a *AuthEndpoint) Login(
	email string,
	password string,
) (*LoginResponse, error) {

	requestBody := CredentialRequest{
		Email:    email,
		Password: password,
	}

	resp, err := a.endpoint.Request(http.MethodPost, LoginURI, requestBody)

	if err != nil {
		return nil, fmt.Errorf("%w error getting login", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, fmt.Errorf("%w error reading the response login", err)
	}

	loginResponse := LoginResponse{}

	return &loginResponse, json.Unmarshal(body, &loginResponse)
}
