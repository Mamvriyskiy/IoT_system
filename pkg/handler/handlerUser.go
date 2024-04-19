package handler

import (
	"net/http"

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

	// c.Next()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Password string `json:"password"`
	Username string `json:"login"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	token, err := h.services.IUser.GenerateToken(input.Username, input.Password)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
