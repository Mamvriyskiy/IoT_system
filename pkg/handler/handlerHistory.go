package handler

import (
	"crypto/rand"
	"math/big"
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
	logger "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3"
)

func generateRandomInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		logger.Log("Error", "rand.Int", "Error rand number:", err, rand.Reader, big.NewInt(int64(max)))
		return 0
	}
	return int(n.Int64())
}

func generateRandomFloat(max float64) float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max*100)))
	if err != nil {
		logger.Log("Error", "rand.Int", "Error rand number:", err, rand.Reader, big.NewInt(int64(max)))
		return 0.0
	}
	return float64(n.Int64()) / 100.0
}

func (h *Handler) createDeviceHistory(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var input pkg.AddHistory
	if err := c.BindJSON(&input); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	history := pkg.AddHistory{
		Name:             input.Name,
		TimeWork:         generateRandomInt(101),
		AverageIndicator: generateRandomFloat(100),
		EnergyConsumed:   generateRandomInt(101),
	}

	idHistory, err := h.services.IHistoryDevice.CreateDeviceHistory(id.(int), history)
	if err != nil {
		logger.Log("Error", "CreateDeviceHistory", "Error create history:", err, id.(int), history)
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
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userID")
		return
	}

	var info pkg.AddHistory
	if err := c.BindJSON(&info); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	input, err := h.services.IHistoryDevice.GetDeviceHistory(id.(int), info.Name)
	if err != nil {
		logger.Log("Error", "GetDeviceHistory", "Error get history:", err, id.(int), info.Name)
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: input,
	})
}
