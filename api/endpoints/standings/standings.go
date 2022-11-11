package standings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adalessa/qatar/pkg/endpoint"
)

type StandingsEndpoint struct {
	endpoint endpoint.Endpoint
}

const GetByGroup = "/api/v1/standings/%s"

func NewEndpoint(
	domain string,
	token string,
) *StandingsEndpoint {
	endpoint := endpoint.New(domain)
	endpoint.SetToken(token)

	return &StandingsEndpoint{endpoint}
}

func (e *StandingsEndpoint) GetByGroup(group string) (*StandingResponse, error) {
	path := fmt.Sprintf(GetByGroup, group)
	resp, err := e.endpoint.Request(http.MethodGet, path, nil)

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
