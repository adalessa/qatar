package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adalessa/qatar/api/endpoints/auth"
)

func TestMain(t *testing.T) {
	testEmail := "alpha@test.com"
	testPassword := "123456"
	expectedPath := "/api/v1/user/login"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type: application/json header, got: %s", r.Header.Get("Content-Type"))
		}

		defer r.Body.Close()

		content := new(auth.CredentialRequest)

		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Could not read the request body got: %s", err)
		}

		err = json.Unmarshal(resp, content)
		if err != nil {
			t.Errorf("Could not decode json got: %s", err)
		}

		if content.Email != testEmail {
			t.Errorf("Wrong email receive expectin %s but got %s", testEmail, content.Email)
		}

		if content.Password != testPassword {
			t.Errorf("Wrong password receive expectin %s but got %s", testPassword, content.Password)
		}

		w.WriteHeader(http.StatusOK)
		response := auth.LoginResponse{
			Status: "success",
			Credential: auth.Credential{
				Token: "asdf123456789",
			},
		}
		responseBody, err := json.Marshal(response)
		if err != nil {
			t.Errorf("Could endecode json got: %s", err)
		}

		w.Write(responseBody)
	}))
	defer server.Close()

	authEndpoint := auth.NewEndpoint(server.URL)

	tokenResp, err := authEndpoint.Login(testEmail, testPassword)
	if err != nil {
		t.Error(err)
	}
	if tokenResp.Status != "success" {
		t.Error("Token is not successfull")
	}
}
