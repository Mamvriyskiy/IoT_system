package handler

import (
	"net/http"
	"math/rand"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDeviceHistory(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		// *TODO: log
		return
	}

	var input pkg.AddHistory
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	history := pkg.AddHistory{
		Name : input.Name,
		TimeWork : rand.Intn(101),
		AverageIndicator : rand.Float64() * 100,
		EnergyConsumed :  rand.Intn(101),
	}

	idHistory, err := h.services.IHistoryDevice.CreateDeviceHistory(id.(int), history)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"idHistory": idHistory,
	})
}

type getAllListResponse struct {
	Data []pkg.DevicesHistory `json:"data"`
}

func (h *Handler) getDeviceHistory(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		// *TODO: log
		return
	}

	var info pkg.AddHistory
	if err := c.BindJSON(&info); err != nil {
		// *TODO: log
		return
	}

	input, err := h.services.IHistoryDevice.GetDeviceHistory(id.(int), info.Name)
	if err != nil {
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: input,
	})
}
