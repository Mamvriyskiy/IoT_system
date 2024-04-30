package handler

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()
	// *TODO: middlewear
	auth := router.Group("/auth")
	auth.POST("/sign-up", h.signUp)
	auth.POST("/sign-in", h.signIn)

	api := router.Group("/api", h.userIdentity)
	home := api.Group("/home")
	home.POST("/", h.createHome)
	home.DELETE("/", h.deleteHome)
	home.PUT("/", h.updateHome)
	home.GET("/", h.listHome)

	access := api.Group("/access")
	access.POST("/", h.addUser)
	access.DELETE("/", h.deleteUser)
	access.GET("/", h.getListUserHome)
	access.PUT("/level/", h.updateLevel)
	access.PUT("/status/", h.updateStatus)

	devices := api.Group("/device")
	devices.POST("/", h.createDevice)
	devices.DELETE("/", h.deleteDevice)

	deviceHistory := api.Group("/history")
	deviceHistory.POST("/", h.createDeviceHistory)
	deviceHistory.GET("/", h.getDeviceHistory)

	return router
}
