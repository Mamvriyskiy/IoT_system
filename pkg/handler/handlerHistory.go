package handler

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func (h *Handler) createDeviceHistory(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }
	fmt.Println("+")
	var input pkg.DevicesHistory
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	deviceID := 1
	idHistory, err := h.services.IHistoryDevice.CreateDeviceHistory(deviceID, input)
	fmt.Println(err)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"idHistory":idHistory,
	})
}

func (h *Handler) updateDeviceHistory(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }

	var input pkg.DevicesHistory
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	idDevice := 0
	err := h.services.IHistoryDevice.UpdateDeviceHistory(idDevice, input)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) getDeviceHistory(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }

	idDevice := 0
	input, err := h.services.IHistoryDevice.GetDeviceHistory(idDevice)
	if err != nil {
		// *TODO log
		return
	}

	_ = input
}
