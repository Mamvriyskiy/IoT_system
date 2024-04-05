package pkg

type Devices struct {
	Name             string `json:"name"`
	TypeDevice       string `json:"typeDevice"`
	Status           string `json:"status"`
	Brand            string `json:"brand"`
	deviceID         int    `json:"-" db:"deviceID"`
	PowerConsumption int    `json:"powerConsumption"`
	MinParameter     int    `json:"minParametr"`
	MaxParameter     int    `json:"maxParametr"`
}
