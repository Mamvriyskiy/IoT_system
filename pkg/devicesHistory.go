package pkg

type DevicesHistory struct {
	ID               int     `json:"-" db:"historyDev"`
	TimeWork         int     `json:"timework"`
	AverageIndicator float64 `json:"average"`
	EnergyConsumed   int     `json:"energy"`
}
