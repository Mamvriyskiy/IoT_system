package handler

import (
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

	idDevice, err := h.services.IDevice.CreateDevice(input)
	if err != nil {
		// *TODO log
		return
	}

	_ = idDevice
}

func (h *Handler) deleteDevice(c *gin.Context) {
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

	idDevice := 0
	err := h.services.IDevice.DeleteDevice(idDevice, input)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) updateDevice(c *gin.Context) {
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

	idDevice := 0
	err := h.services.IDevice.UpdateDevice(idDevice, input)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) addHomeDevice(c *gin.Context) {
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

	idDevice := 0
	idHome := 0
	err := h.services.IDevice.AddHomeDevice(idHome, idDevice, input)
	if err != nil {
		// *TODO log
		return
	}
}

func (h *Handler) deleteHomeDevice(c *gin.Context) {
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

	idDevice := 0
	idHome := 0
	err := h.services.IDevice.DeleteHomeDevice(idHome, idDevice, input)
	if err != nil {
		// *TODO log
		return
	}
}
