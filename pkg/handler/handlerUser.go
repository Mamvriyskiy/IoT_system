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

	id, err := h.services.IUser.CreateUser(input)
	if err != nil {
		// *TODO log
		return
	}

	_ = id
}

func (h *Handler) signIn(c *gin.Context) {
	var input pkg.User
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	cmp, id, err := h.services.IUser.CheckUser(input)
	if err != nil {
		// *TODO log
		return
	}

	_ = id
	_ = cmp
}
