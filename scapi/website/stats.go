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
	CurrentEtf  string `json:"current_etf"`
	CurrentLive string `json:"current_live"`
	CurrentPtu  string `json:"current_ptu"`
	Fans        uint32
	Fleet       uint32
	Funds       float64
}

func (w *Website) Stats() (*http.Response, *StatsResponse, error) {
	stats := new(StatsResponse)
	res, err := w.sling.Path("auto/").Get("stats").ReceiveSuccess(stats)

	return res, stats, err
}
