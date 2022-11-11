package endpoint

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type Endpoint struct {
	domain string
	token  string
}

func New(domain string) Endpoint {
	return Endpoint{
		domain: domain,
	}
}

func (e *Endpoint) SetToken(token string) {
	e.token = token
}

func (e *Endpoint) Request(method string, uri string, body interface{}) (*http.Response, error) {

	if !strings.HasPrefix("/", uri) {
		uri = "/" + uri
	}

	path := e.domain + uri

	json_data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	if method == http.MethodGet {
		req, err = http.NewRequest(method, path, nil)
	} else {
		req, err = http.NewRequest(method, path, bytes.NewBuffer(json_data))
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if e.token != "" {
		req.Header.Add("Authentication", "Bearer "+e.token)
	}

	client := http.Client{}

	return client.Do(req)
}
