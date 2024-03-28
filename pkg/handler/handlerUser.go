package handler

import (
	"github.com/gin-gonic/gin"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/src/user"
)

func (h *Handler) signUp(c *gin.Context) {
	var input user.User

	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	h.ser
	//input.CreateUser()
}

func (h *Handler) signIn(c *gin.Context) {

}
