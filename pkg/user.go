package pkg

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int    `json:"-"`
}
