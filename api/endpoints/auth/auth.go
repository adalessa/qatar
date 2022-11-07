package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthEndpoint struct {
	domain string
}

const LoginURI = "api/v1/user/login"

func NewEndpoint(domain string) *AuthEndpoint {
	return &AuthEndpoint{domain: domain}
}

func (a *AuthEndpoint) Login(
	email string,
	password string,
) (*LoginResponse, error) {
	path := fmt.Sprintf("%s/%s", a.domain, LoginURI)
	content := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    email,
		Password: password,
	}

	json_data, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, fmt.Errorf("%w error creating the request to login", err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("%w error getting login", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w error reading the response login", err)
	}

	loginResponse := LoginResponse{}

	return &loginResponse, json.Unmarshal(body, &loginResponse)
}
