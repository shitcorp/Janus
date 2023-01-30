package scapiWebsite

import (
	"fmt"
	"net/http"

	"github.com/rotisserie/eris"
	scapiUtils "github.com/shitcorp/janus/scapi/utils"
)

type OrgResponse struct {
	scapiUtils.BaseRes
	Data OrgData
}

type OrgData struct {
	Archetype  string
	Banner     string
	Commitment string
	Focus      OrgDataFocusItems
	Headline   OrgDataHeadline
	Href       string
	Lang       string
	Logo       string
	Members    uint32
	Name       string
	Recruiting bool
	Roleplay   bool
	Sid        string
	Url        string
}

type OrgDataFocusItems struct {
	Primary   OrgDataFocus
	Secondary OrgDataFocus
}

type OrgDataFocus struct {
	Image string
	Name  string
}

type OrgDataHeadline struct {
	Html      string
	Plaintext string
}

func (w *Website) Org(sid string) (*http.Response, *OrgResponse, error) {
	org := new(OrgResponse)
	res, err := w.sling.Path("auto/").Get(fmt.Sprintf("organization/%s", sid)).ReceiveSuccess(org)

	var noOrg OrgData
	if org.Data == noOrg && org.Success != 0 {
		err = eris.New("SC API Error: no data")
	}

	return res, org, eris.Wrap(err, "Star Citizen API stats endpoint")
}
