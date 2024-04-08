package handler

import (
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDevice(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }

	var input pkg.Devices
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	homeID := 1
	idDevice, err := h.services.IDevice.CreateDevice(homeID, input)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"idDevice": idDevice,
	})
}

func (h *Handler) deleteDevice(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }

	idDevice := 1
	err := h.services.IDevice.DeleteDevice(idDevice)
	if err != nil {
		// *TODO log
		return
	}
}
