package handler

import (
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDeviceHistory(c *gin.Context) {
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

	deviceID := 1
	idHistory, err := h.services.IHistoryDevice.CreateDeviceHistory(deviceID, input)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"idHistory": idHistory,
	})
}

type getAlllistResponse struct {
	Data []pkg.DevicesHistory `json:"Data"`
}

func (h *Handler) getDeviceHistory(c *gin.Context) {
	// id, ok := c.Get(userCtx)
	// if !ok {
	// 	// *TODO: log
	// 	return
	// }

	idDevice := 1
	input, err := h.services.IHistoryDevice.GetDeviceHistory(idDevice)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, getAlllistResponse{
		Data: input,
	})
}
