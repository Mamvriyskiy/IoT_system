package handler

import (
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(c *gin.Context) {
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	userID := 1
	homeID := 1
	idAccess, err := h.services.IAccessHome.AddUser(homeID, userID, input)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessID": idAccess,
	})
}

func (h *Handler) deleteUser(c *gin.Context) {
	idUser := 1

	err := h.services.IAccessHome.DeleteUser(idUser)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) updateLevel(c *gin.Context) {
	idUser := 2
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
	idUser := 2
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

type getAlllistUserResponse struct {
	Data []pkg.ClientHome `json:"Data"`
}

func (h *Handler) getListUserHome(c *gin.Context) {
	homeID := 1
	listUser, err := h.services.IAccessHome.GetListUserHome(homeID)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, getAlllistUserResponse{
		Data: listUser,
	})
}
