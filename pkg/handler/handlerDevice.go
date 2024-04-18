package handler

import (
	"fmt"
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDevice(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		// *TODO: log
		return
	}

	var input pkg.Devices
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	idDevice, err := h.services.IDevice.CreateDevice(id.(int), &input)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"idDevice": idDevice,
	})
}

func (h *Handler) deleteDevice(c *gin.Context) {
	id, ok := c.Get("userID")
	fmt.Println(id, ok)
	if !ok {
		// *TODO: log
		return
	}

	fmt.Println("1")
	var input pkg.Devices
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	fmt.Println("2")
	err := h.services.IDevice.DeleteDevice(id.(int), input.Name)
	if err != nil {
		// *TODO log
		return
	}
}
