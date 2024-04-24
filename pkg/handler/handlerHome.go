package handler

import (
	"net/http"

	logger "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3"
	logger "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.Home
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	homeID, err := h.services.IHome.CreateHome(id.(int), input)
	if err != nil {
		logger.Log("Error", "CreateHome", "Error create home:", err, id.(int), input)
		return
	}

	_, err = h.services.IAccessHome.AddOwner(id.(int), homeID)
	if err != nil {
		logger.Log("Error", "AddOwner", "Error add owner:", err, id.(int), homeID)
		return
	}

	c.Set("homeID", homeID)
	c.Next()

	c.JSON(http.StatusOK, map[string]interface{}{
		"homeId": homeID,
	})

	logger.Log("Info", "", "A home has been created", nil)
}

func (h *Handler) deleteHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	err := h.services.IHome.DeleteHome(id.(int))
	if err != nil {
		logger.Log("Error", "DeleteHome", "Error delete home:", err, id.(int))
		return
	}

	logger.Log("Info", "", "A home has been deleted", nil)
}

type getAllListHomeResponse struct {
	Data []pkg.Home `json:"data"`
}

func (h *Handler) listHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	homeListUser, err := h.services.IHome.ListUserHome(id.(int))
	if err != nil {
		logger.Log("Error", "ListUserHome", "Error get user:", err, id.(int))
		return
	}

	c.JSON(http.StatusOK, getAllListHomeResponse{
		Data: homeListUser,
	})

	logger.Log("Info", "", "The list of users has been received", nil)
}

func (h *Handler) updateHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.Home
	err := c.BindJSON(&input)
	if err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	input.OwnerID, ok = id.(int)
	if !ok {
		logger.Log("Warning", "*.(int)", "Error convert to int", nil, id)
		return
	}

	err = h.services.IHome.UpdateHome(input)
	if err != nil {
		logger.Log("Error", "UpdateHome", "Error update home:", err, "")
		return
	}

	logger.Log("Info", "", "A home has been deleted", nil)
}
