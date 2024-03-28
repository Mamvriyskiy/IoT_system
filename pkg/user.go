package user

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username" binding:"required"` 
	Email    string `json:"email" binding:"required"` 
	Password string `json:"password" binding:"required"` 
}