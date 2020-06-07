package steam

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client defines the Steam Client.
type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	IPlayerService *IPlayerService
}

const (
	// APIBase is the URL prefix for the Steamworks Web API
	APIBase = "https://api.steampowered.com/"
)

// NewClient returns a new Client with the given API key.
func NewClient(apikey string, httpClient *http.Client) *Client {
	c := &Client{}
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c.common = service{
		apikey: apikey,
		client: httpClient,
	}
	c.IPlayerService = (*IPlayerService)(&c.common)
	return c
}

type service struct {
	apikey string
	client *http.Client
}

func (s *service) get(ctx context.Context, endpoint string, query url.Values) ([]byte, error) {
	url := APIBase + endpoint
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request err :%w", err)
	}
	query.Set("key", s.apikey)
	request.URL.RawQuery = query.Encode()

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := s.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http client do err :%w", err)
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read body err :%w", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("response code expects 200, but got %d, response body %s",
			response.StatusCode, bodyBytes)
	}

	return bodyBytes, nil
}
