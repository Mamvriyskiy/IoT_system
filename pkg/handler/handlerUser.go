package handler

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input pkg.User

	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	//input.CreateUser()
}

func (h *Handler) signIn(c *gin.Context) {

}
