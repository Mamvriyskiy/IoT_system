package pkg

type DevicesHistory struct {
	ID               int     `db:"historyDev"       json:"-"`
	TimeWork         int     `db:"timework"         json:"timework"`
	AverageIndicator float64 `db:"averageindicator" json:"average"`
	EnergyConsumed   int     `db:"energyconsumed"   json:"energy"`
}
