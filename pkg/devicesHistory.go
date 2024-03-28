package devicesHistory

type DevicesHistory struct {
	Id               int     `json:"-"`
	DevicesID        int     `json:"devicesId"`
	TimeWork         int     `json:"timework"` //second
	AverageIndicator float64 `json:"average"`
	EnergyConsumed   int     `json:"energy"`
}
