package pkg

type AccessHome struct {
	ID           int  `json:"-" db:"accessID"`
	AccessStatus string `json:"status"`
	AccessLevel  int  `json:"level"`
}
