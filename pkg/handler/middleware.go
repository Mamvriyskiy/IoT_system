package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	//logger "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return
	}

	userID, err := h.services.IUser.ParseToken(headerParts[1])
	if err != nil {
		return
	}

	c.Set("userID", userID)
}
