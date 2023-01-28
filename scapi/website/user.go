package scapiWebsite

import (
	"fmt"
	"net/http"

	"github.com/rotisserie/eris"
	scapiUtils "github.com/shitcorp/janus/scapi/utils"
)

type UserResponse struct {
	scapiUtils.BaseRes
	Data UserData
}

type UserData struct {
	Organization UserDataOrg
	Profile      UserDataProfile
}

type UserDataOrg struct {
	Image string
	Name  string
	Rank  string
	Sid   string
}

type UserDataProfile struct {
	Badge      string
	BadgeImage string `json:"badge_image"`
	Display    string
	Enlisted   string
	Fluency    []string
	Handle     string
	Id         string
	Image      string
	Page       UserDataProfilePage
}

type UserDataProfilePage struct {
	Title string
	Url   string
}

func (w *Website) User(handle string) (*http.Response, *UserResponse, error) {
	user := new(UserResponse)
	res, err := w.sling.Path("auto/").Get(fmt.Sprintf("user/%s", handle)).ReceiveSuccess(user)

	return res, user, eris.Wrap(err, "Star Citizen API stats endpoint")
}
