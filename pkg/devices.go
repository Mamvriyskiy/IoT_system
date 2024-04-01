package pkg

type Devices struct {
	ID               int    `json:"-"`
	Name             string `json:"name"`
	TypeDevice       string `json:"type"`
	Status           string `json:"status"`
	Brand            string `json:"brand"`
	PowerConsumption uint   `json:"power"`
	MinParameter     int    `json:"minl"`
	MaxParameter     int    `json:"maxl"`
}
