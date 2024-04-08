package pkg

type ClientHome struct {
	Username     string `db:"login"`
	AccessLevel  int    `db:"accesslevel"`
	AccessStatus string `db:"accessstatus"`
}
