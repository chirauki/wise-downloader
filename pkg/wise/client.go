package wise

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

const (
	wiseApiEndpoint   = "api.transferwise.com"
	wiseProfilesPath  = "/v1/profiles"
	wiseTransfersPath = "/v1/transfers"
)

type WiseClient struct {
	cli *resty.Client
}

func NewClient(token string) *WiseClient {
	client := resty.New()
	client.SetBaseURL("https://" + wiseApiEndpoint)
	client.SetAuthScheme("Bearer")
	client.SetAuthToken(token)

	return &WiseClient{
		cli: client,
	}
}

func (w WiseClient) GetProfiles() ([]Profile, error) {
	var p []Profile
	resp, err := w.cli.R().Get(wiseProfilesPath)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(resp.Body(), &p)
	if err != nil {
		return p, err
	}
	return p, nil
}
