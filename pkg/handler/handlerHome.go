package handler

import (
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

	id := 0
	idHome, err := h.services.IHome.CreateHome(id, input)
	if err != nil {
		// *TODO log
		return
	}

	_ = idHome
}

func (h *Handler) deleteHome(c *gin.Context) {
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

	idUser := 0
	err := h.services.IHome.DeleteHome(idUser, input)
	if err != nil {
		// *TODO log
		return
	}

	_ = idUser
}

func (h *Handler) updateHome(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }

	idUser := 0
	var input pkg.Home
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	err := h.services.IHome.UpdateHome(idUser, input)
	if err != nil {
		// *TODO log
		return
	}
}


