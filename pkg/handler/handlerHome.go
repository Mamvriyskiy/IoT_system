package handler

import (
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createHome(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }
	var input pkg.Home
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	id := 1
	idHome, err := h.services.IHome.CreateHome(id, input)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"homeId": idHome,
	})
}

func (h *Handler) deleteHome(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }
	homeID := 1
	err := h.services.IHome.DeleteHome(homeID)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) updateHome(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }

	var input pkg.Home
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	input.ID = 1

	err := h.services.IHome.UpdateHome(input)
	if err != nil {
		// *TODO log
		return
	}
}
