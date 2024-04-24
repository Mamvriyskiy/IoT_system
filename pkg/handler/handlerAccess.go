package handler

import (
	"net/http"

	logger "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.AddUserHome
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	idAccess, err := h.services.IAccessHome.AddUser(userID.(int), input.AccessLevel, input.Email)
	if err != nil {
		logger.Log("Error", "AddUser", "Error create access:", 
			err, userID.(int), input.AccessLevel, input.Email)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessID": idAccess,
	})

	logger.Log("Info", "", "The user has been granted access", nil)
}

func (h *Handler) deleteUser(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.AddUserHome
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IAccessHome.DeleteUser(userID.(int), input.Email)
	if err != nil {
		logger.Log("Error", "DeleteUser", "Error delete access:", err, userID.(int), input.Email)
		return
	}

	logger.Log("Info", "", "The user's access was deleted", nil)
}

func (h *Handler) updateLevel(c *gin.Context) {
	idUser := 2
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IAccessHome.UpdateLevel(idUser, input)
	if err != nil {
		logger.Log("Error", "UpdateLevel", "Error update access:", err, idUser, input)
		return
	}

	logger.Log("Info", "", "A level has been update", nil)
}

func (h *Handler) updateStatus(c *gin.Context) {
	idUser := 2
	var input pkg.AccessHome
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IAccessHome.UpdateStatus(idUser, input)
	if err != nil {
		logger.Log("Error", "UpdateStatus", "Error update access:", err, idUser, input)
		return
	}

	logger.Log("Info", "", "A status has been update", nil)
}

type getAllListUserResponse struct {
	Data []pkg.ClientHome `json:"data"`
}

func (h *Handler) getListUserHome(c *gin.Context) {
	homeID := 1
	listUser, err := h.services.IAccessHome.GetListUserHome(homeID)
	if err != nil {
		logger.Log("Error", "GetListUserHome", "Error get access:", err, homeID)
		return
	}

	c.JSON(http.StatusOK, getAllListUserResponse{
		Data: listUser,
	})

	logger.Log("Info", "", "The list of users has been received", nil)
}
