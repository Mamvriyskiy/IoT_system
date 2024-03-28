package devices

type Devices struct {
	Id               int    `json:"-"`
	Name             string `json:"name"`
	typeDevice       string `json:"type"`
	status           string `json:"status"`
	brand            string `json:"brand"`
	powerСonsumption uint   `json:"power"`
	minParameter     int    `json:"minl"`
	maxParameter     int    `json:"maxl"`
}
