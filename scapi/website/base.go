package scapiWebsite

import (
	"github.com/dghubble/sling"
)

type Website struct {
	sling *sling.Sling
}

func CreateWebsite(sling *sling.Sling) *Website {
	return &Website{sling: sling}
}
