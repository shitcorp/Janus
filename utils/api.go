package utils

import (
	"github.com/shitcorp/janus/scapi"
)

var Api = scapi.CreateApiClient(Config.ScApiToken, "Janus Bot")
