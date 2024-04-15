package pkg

type AccessHome struct {
	ID           int  `json:"-"`
	HomeID       int  `json:"homeid"`
	AccessStatus bool `json:"status"`
	AccessLevel  int  `json:"level"`
}
