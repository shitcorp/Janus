package scapiWebsite

import (
	"net/http"
	"reflect"

	"github.com/rotisserie/eris"
	scapiUtils "github.com/shitcorp/janus/scapi/utils"
)

type ShipQuery struct {
	Name           string `url:"name,omitempty"`
	Classification string `url:"classification,omitempty"`

	LengthMin int `url:"length_min,omitempty"`
	LengthMax int `url:"length_max,omitempty"`

	CrewMin int `url:"crew_min,omitempty"`
	CrewMax int `url:"crew_max,omitempty"`

	PriceMin int `url:"price_min,omitempty"`
	PriceMax int `url:"price_max,omitempty"`

	MassMin int `url:"mass_min,omitempty"`
	MassMax int `url:"mass_max,omitempty"`

	PageMax int `url:"page_max,omitempty"`

	Id int `url:"id,omitempty"`
}

type ShipsResponse struct {
	scapiUtils.BaseRes
	Data []ShipsData
}

type ShipsData struct {
	AfterburnerSpeed       string `json:"afterburner_speed"`
	Beam                   string
	CargoCapacity          string
	ChassisId              string `json:"chassis_id"`
	Compiled               ShipsDataCompiled
	Description            string
	Focus                  string
	Height                 string
	Id                     string
	Length                 string
	Manufacturer           ShipsDataManufacturer
	ManufacturerId         string `json:"manufacturer_id"`
	Mass                   string
	MaxCrew                string `json:"max_crew"`
	Media                  []ShipsDataMedia
	MinCrew                string
	Name                   string
	PitchMax               string `json:"pitch_max"`
	Price                  float32
	ProductionNote         string `json:"production_note"`
	ProductionStatus       string `json:"production_status"`
	RollMax                string `json:"roll_max"`
	ScmSpeed               string `json:"scm_speed"`
	Size                   string
	TimeModified           string `json:"time_modified"`
	TimeModifiedUnfiltered string `json:"time_modified.unfiltered"`
	Type                   string
	Url                    string
	XAxisAccel             string `json:"xaxis_acceleration"`
	YawMax                 string `json:"yaw_max"`
	YAxisAccel             string `json:"yaxis_acceleration"`
	ZAxisAccel             string `json:"zaxis_acceleration"`
}

type ShipsDataCompiled struct {
	// missing items
	// not documented
}

type ShipsDataManufacturer struct {
	// missing items
	// not documented
}

type ShipsDataMedia struct {
	// missing items
	// not documented
	Id           string
	Images       ShipsDataMediaImages
	SourceUrl    string `json:"source_url"`
	TimeModified string `json:"time_modified"`
}

type ShipsDataMediaImages struct {
	// missing items
	// not documented
	Avatar string
	Banner string
	Cover  string
}

func (w *Website) Ships(params ShipQuery) (*http.Response, *ShipsResponse, error) {
	ship := new(ShipsResponse)
	res, err := w.sling.Path("live/").Get("ships").QueryStruct(params).ReceiveSuccess(ship)

	// fmt.Sprintf("ships/%s", params.Name)

	var noShip ShipsData
	if reflect.DeepEqual(ship.Data, noShip) && ship.Success != 0 {
		err = eris.New("SC API Error: no data")
	}

	return res, ship, eris.Wrap(err, "Star Citizen API stats endpoint")
}
