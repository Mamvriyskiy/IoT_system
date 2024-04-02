package handler

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(c *gin.Context) {
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	idAccess, err := h.services.IAccessHome.AddUser(input)
	if err != nil {
		// *TODO log
		return
	}

	_ = idAccess
}

func (h *Handler) deleteUser(c *gin.Context) {
	idUser := 0
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	err := h.services.IAccessHome.DeleteUser(idUser, input)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) updateLevel(c *gin.Context) {
	idUser := 0
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	err := h.services.IAccessHome.UpdateLevel(idUser, input)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) updateStatus(c *gin.Context) {
	idUser := 0
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	err := h.services.IAccessHome.UpdateStatus(idUser, input)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) getListUserHome(c *gin.Context) {
	homeId := 0
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	listUser, err := h.services.IAccessHome.GetListUserHome(homeId, input)
	if err != nil {
		// *TODO log
		return
	}

	_ = listUser
}
