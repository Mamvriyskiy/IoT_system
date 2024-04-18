package pkg

type ClientHome struct {
	Username     string `db:"login"`
	AccessStatus string `db:"accessstatus"`
	AccessLevel  int    `db:"accesslevel"`
}

type AddUserHome struct {
	AccessLevel int    `json: "accessLevel"`
	Email       string `json: "email"`
}

type AddHistory struct {
	Name             string  `db:"name"             json:"name"`
	TimeWork         int     `db:"timework"         json:"timework"`
	AverageIndicator float64 `db:"averageindicator" json:"average"`
	EnergyConsumed   int     `db:"energyconsumed"   json:"energy"`
}
