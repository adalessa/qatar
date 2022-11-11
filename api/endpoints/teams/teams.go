package teams

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adalessa/qatar/pkg/endpoint"
)

type TeamsEndpoint struct {
	endpoint endpoint.Endpoint
}

const GetTeamsURI = "/api/v1/team"

func NewEndpoint(
	domain string,
	token string,
) *TeamsEndpoint {
	endpoint := endpoint.New(domain)
	endpoint.SetToken(token)

	return &TeamsEndpoint{endpoint}
}

func (t *TeamsEndpoint) GetTeams() (*GetTeamsResponse, error) {
	resp, err := t.endpoint.Request(http.MethodGet, GetTeamsURI, nil)

	if err != nil {
		return nil, fmt.Errorf("%w error getting the teams", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w error reading the response getting the teams", err)
	}

	teamsResponse := GetTeamsResponse{}

	return &teamsResponse, json.Unmarshal(body, &teamsResponse)
}
