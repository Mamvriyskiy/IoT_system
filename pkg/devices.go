package pkg

type Devices struct {
	Name             string `json:"name"`
	TypeDevice       string `json:"type"`
	Status           string `json:"status"`
	Brand            string `json:"brand"`
	ID               int    `json:"-"`
	PowerConsumption int    `json:"power"`
	MinParameter     int    `json:"minl"`
	MaxParameter     int    `json:"maxl"`
}
