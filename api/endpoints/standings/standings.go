package standings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StandingsEndpoint struct {
	domain string
	token  string
}

const GetByGroup = "api/v1/standings/%s"

func NewEndpoint(
	domain string,
	token string,
) *StandingsEndpoint {
	return &StandingsEndpoint{
		domain: domain,
		token:  token,
	}
}

func (e *StandingsEndpoint) GetByGroup(group string) (*StandingResponse, error) {
	path := fmt.Sprintf("%s/%s", e.domain, fmt.Sprintf(GetByGroup, group))
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("%w error creating the request to get the stansings by group", err)
	}

	req.Header.Add("Authorization", e.token)

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("%w error getting the standings", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w error reading the response getting the stansings", err)
	}

	standingResponse := StandingResponse{}

	return &standingResponse, json.Unmarshal(body, &standingResponse)
}
