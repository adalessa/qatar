package teams

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TeamsEndpoint struct {
	domain string
	token  string
}

const GetTeamsURI = "api/v1/team"

func NewEndpoint(
	domain string,
	token string,
) *TeamsEndpoint {
	return &TeamsEndpoint{
		domain: domain,
		token:  token,
	}
}

func (t *TeamsEndpoint) GetTeams() (*GetTeamsResponse, error) {
	path := fmt.Sprintf("%s/%s", t.domain, GetTeamsURI)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("%w error creating the request to get the teams", err)
	}

	req.Header.Add("Authorization", t.token)

	client := http.Client{}

	resp, err := client.Do(req)

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
