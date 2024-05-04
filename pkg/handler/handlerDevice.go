package handler

import (
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDevice(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.Devices
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	idDevice, err := h.services.IDevice.CreateDevice(id.(int), &input)
	if err != nil {
		logger.Log("Error", "CreateDevice", "Error create device:", err, id.(int), &input)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"idDevice": idDevice,
	})

	logger.Log("Info", "", "A device has been created", nil)
}

func (h *Handler) deleteDevice(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.Devices
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	err := h.services.IDevice.DeleteDevice(id.(int), input.Name)
	if err != nil {
		logger.Log("Error", "DeleteDevice", "Error delete device:", err, id.(int), input.Name)
		return
	}

	logger.Log("Info", "", "A device has been deleted", nil)
}
