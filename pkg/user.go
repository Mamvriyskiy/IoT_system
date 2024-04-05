package pkg

type User struct {
	Password string `json:"password"`
	Username string `json:"login"`
	Email    string `json:"email"`
	ID       int    `json:"-" db:"clientid"`
}
