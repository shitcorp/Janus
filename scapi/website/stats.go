package scapiWebsite

import (
	scapiUtils "github.com/shitcorp/janus/scapi/utils"
	"net/http"
)

type StatsResponse struct {
	scapiUtils.BaseRes
	Data StatsData
}

type StatsData struct {
	CurrentLive string `json:"current_live"`
	Fans        uint32
	Fleet       uint32
	Funds       uint64
}

func (w *Website) Stats() (*http.Response, *StatsResponse, error) {
	stats := new(StatsResponse)
	res, err := w.sling.Path("auto").Get("stats").ReceiveSuccess(stats)

	return res, stats, err
}
