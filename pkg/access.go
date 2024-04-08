package pkg

type AccessHome struct {
	ID           int    `db:"accessID" json:"-"`
	AccessStatus string `json:"status"`
	AccessLevel  int    `json:"level"`
}
