package matches

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adalessa/qatar/pkg/endpoint"
)

type MatchesEndpoint struct {
	endpoint endpoint.Endpoint
}

const GetByMatchDayIDURI = "/api/v1/bymatch/%s"
const GetByIDURI = "/api/v1/match/%s"
const PostSearchByDate = "/api/v1/bydate"

func NewEndpoint(
	domain string,
	token string,
) *MatchesEndpoint {
	endpoint := endpoint.New(domain)
	endpoint.SetToken(token)

	return &MatchesEndpoint{endpoint}
}
func (t *MatchesEndpoint) GetById(ID string) (*MatchesResponse, error) {
	matches, err := t.getRequest(fmt.Sprintf(GetByIDURI, ID))
	if err != nil {
		return nil, fmt.Errorf("%w error getting matches by id", err)
	}

	return matches, nil
}

func (t *MatchesEndpoint) GetByMatchDayID(MatchDayID string) (*MatchesResponse, error) {
	matches, err := t.getRequest(fmt.Sprintf(GetByMatchDayIDURI, MatchDayID))
	if err != nil {
		return nil, fmt.Errorf("%w error getting matches by id", err)
	}

	return matches, nil
}

func (t *MatchesEndpoint) GetByDate(date string) (*MatchesResponse, error) {
	requestBody := MatchesByDateRequest{
		Date: date,
	}

	resp, err := t.endpoint.Request(http.MethodPost, PostSearchByDate, requestBody)

	if err != nil {
		return nil, fmt.Errorf("%w error getting the match by team di", err)
	}

	return decode_response(resp)
}

func (t *MatchesEndpoint) getRequest(uri string) (*MatchesResponse, error) {
	resp, err := t.endpoint.Request(http.MethodGet, uri, nil)

	if err != nil {
		return nil, fmt.Errorf("%w error getting the match by team di", err)
	}

	return decode_response(resp)
}

func decode_response(r *http.Response) (*MatchesResponse, error) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	matchesResponse := MatchesResponse{}

	return &matchesResponse, json.Unmarshal(body, &matchesResponse)
}
