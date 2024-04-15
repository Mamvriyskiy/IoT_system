package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"fmt"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return 
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		fmt.Println("invalid auth header")
		return
	}

	userID, err := h.services.IUser.ParseToken(headerParts[1])
	fmt.Println("===", userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Set("userID", userID)
}

