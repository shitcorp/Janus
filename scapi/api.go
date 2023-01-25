package scapi

import (
	"github.com/dghubble/sling"
	"github.com/shitcorp/janus/scapi/website"
)

type Api struct {
	sling   *sling.Sling
	Website *scapiWebsite.Website
}

var baseUrl = "https://api.starcitizen-api.com/"

func CreateApiClient(apiToken string, userAgent string) *Api {
	if userAgent == "" {
		userAgent = "SC-API Golang Client"
	}

	s := sling.New().Base(baseUrl).Patch(apiToken).Path("v1").Set("User-Agent", userAgent).Set("Accept", "application/json")

	api := &Api{sling: s, Website: scapiWebsite.CreateWebsite(s)}

	return api
}
