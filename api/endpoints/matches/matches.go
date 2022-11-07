package matches

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MatchesEndpoint struct {
	domain string
	token  string
}

const GetByMatchDayIDURI = "api/v1/bymatch/%s"
const GetByIDURI = "api/v1/match/%s"
const PostSearchByDate = "api/v1/bydate"

func NewEndpoint(
	domain string,
	token string,
) *MatchesEndpoint {
	return &MatchesEndpoint{
		domain: domain,
		token:  token,
	}
}
func (t *MatchesEndpoint) GetById(ID string) (*MatchesResponse, error) {
	return t.getRequest(fmt.Sprintf(GetByIDURI, ID))
}

func (t *MatchesEndpoint) GetByMatchDayID(MatchDayID string) (*MatchesResponse, error) {
	return t.getRequest(fmt.Sprintf(GetByMatchDayIDURI, MatchDayID))
}

func (t *MatchesEndpoint) GetByDate(date string) (*MatchesResponse, error) {
	path := fmt.Sprintf("%s/%s", t.domain, PostSearchByDate)
	content := struct {
		Date string `json:"date"`
	}{
		Date: date,
	}
	json_data, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, fmt.Errorf("%w error creating the request to get match by team id", err)
	}

	req.Header.Add("Authorization", t.token)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("%w error getting the match by team di", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w error reading the response getting the match by team id", err)
	}

	matchesResponse := MatchesResponse{}

	return &matchesResponse, json.Unmarshal(body, &matchesResponse)
}

func (t *MatchesEndpoint) getRequest(uri string) (*MatchesResponse, error) {
	path := fmt.Sprintf("%s/%s", t.domain, uri)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("%w error creating the request to get match by team id", err)
	}

	req.Header.Add("Authorization", t.token)

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("%w error getting the match by team di", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w error reading the response getting the match by team id", err)
	}

	matchesResponse := MatchesResponse{}

	return &matchesResponse, json.Unmarshal(body, &matchesResponse)
}
