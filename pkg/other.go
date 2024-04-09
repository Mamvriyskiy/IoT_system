package pkg

type ClientHome struct {
	Username     string `db:"login"`
	AccessStatus string `db:"accessstatus"`
	AccessLevel  int    `db:"accesslevel"`
}
