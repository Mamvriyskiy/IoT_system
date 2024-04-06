package handler

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func (h *Handler) addUser(c *gin.Context) {
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		// *TODO: log
		return
	}

	userID := 1
	homeID := 1
	idAccess, err := h.services.IAccessHome.AddUser(homeID, userID, input)
	fmt.Println(err)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessID":idAccess,
	})
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
	homeID := 0
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	listUser, err := h.services.IAccessHome.GetListUserHome(homeID, input)
	if err != nil {
		// *TODO log
		return
	}

	_ = listUser
}
