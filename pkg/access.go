package pkg

type AccessHome struct {
	ID           int  `json:"-" db:"accessID"`
	AccessStatus bool `json:"status"`
	AccessLevel  int  `json:"level"`
}
