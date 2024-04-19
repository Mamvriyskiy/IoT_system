package handler

import (
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		// *TODO: log
		return
	}

	var input pkg.Home
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	homeID, err := h.services.IHome.CreateHome(id.(int), input)
	if err != nil {
		// *TODO log
		return
	}

	_, err = h.services.IAccessHome.AddOwner(id.(int), homeID)
	if err != nil {
		// *TODO log
		return
	}

	c.Set("homeID", homeID)
	c.Next()

	c.JSON(http.StatusOK, map[string]interface{}{
		"homeId": homeID,
	})
}

func (h *Handler) deleteHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		// *TODO: log
		return
	}

	err := h.services.IHome.DeleteHome(id.(int))
	if err != nil {
		// *TODO log
		return
	}
}

type getAllListHomeResponse struct {
	Data []pkg.Home `json:"data"`
}

func (h *Handler) listHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		// *TODO: log
		return
	}

	homeListUser, err := h.services.IHome.ListUserHome(id.(int))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, getAllListHomeResponse{
		Data: homeListUser,
	})
}

func (h *Handler) updateHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		// *TODO: log
		return
	}

	var input pkg.Home
	err := c.BindJSON(&input)
	if err != nil {
		// *TODO: log
		return
	}

	if err != nil {
		// *TODO log
		return
	}

	input.OwnerID, ok = id.(int)
	if !ok {
		return
	}

	err = h.services.IHome.UpdateHome(input)
	if err != nil {
		// *TODO log
		return
	}
}
