package pkg

type AccessHome struct {
	Id           int  `json:"-"`
	HomeID       int  `json:"homeid"`
	AccessStatus bool `json:"status"`
	AccessLevel  int  `json:"level"`
}
