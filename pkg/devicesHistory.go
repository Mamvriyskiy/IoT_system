package pkg

type DevicesHistory struct {
	ID               int     `json:"-"`
	DevicesID        int     `json:"devicesId"`
	TimeWork         int     `json:"timework"`
	AverageIndicator float64 `json:"average"`
	EnergyConsumed   int     `json:"energy"`
}
