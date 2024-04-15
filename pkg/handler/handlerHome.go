package handler

import (
	"fmt"
	"net/http"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/gin-gonic/gin"
	//"strconv"
)

func (h *Handler) createHome(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		fmt.Println(ok)
		// *TODO: log
		return
	}

	var input pkg.Home
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		// *TODO: log
		return
	}

	ownerID := id.(int)
	idHome, err := h.services.IHome.CreateHome(ownerID, input)
	if err != nil {
		fmt.Println(err, "===")
		// *TODO log
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"homeId": idHome,
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

func (h *Handler) updateHome(c *gin.Context) {
	id, ok := c.Get("homeID")
	fmt.Println("UpdateOK:", ok)
	if !ok {
		// *TODO: log
		return
	}

	var input pkg.Home
	if err := c.BindJSON(&input); err != nil {
		// *TODO: log
		return
	}

	fmt.Println("HomeID:", id)
	input.ID = id.(int)

	err := h.services.IHome.UpdateHome(input)
	if err != nil {
		// *TODO log
		return
	}
}
