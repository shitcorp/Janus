package scapi

import (
	"fmt"

	"github.com/dghubble/sling"
	scapiWebsite "github.com/shitcorp/janus/scapi/website"
)

type Api struct {
	sling   *sling.Sling
	Website *scapiWebsite.Website
}

var baseUrl = "https://api.starcitizen-api.com/"

// CreateApiClient creates an API client for the Star Citizen API
func CreateApiClient(apiToken string, userAgent string) *Api {
	if userAgent == "" {
		userAgent = "SC-API Golang Client"
	}

	s := sling.New().Base(baseUrl).Path(fmt.Sprintf("%s/v1/", apiToken)).Set("User-Agent", userAgent).Set("Accept", "application/json")

	api := &Api{sling: s, Website: scapiWebsite.CreateWebsite(s)}

	return api
}
