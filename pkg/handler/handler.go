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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	home := router.Group("/home")
	{
		home.POST("/", h.createHome)
		home.DELETE("/:id", h.deleteHome)
		home.PUT("/:id", h.updateHome)

		access := router.Group(":id/access")
		{
			access.POST("/", h.addUser)
			access.DELETE("/:id", h.deleteUser)
			access.GET("/:id", h.getListUserHome)
			access.PUT("/level/:id", h.updateLevel)
			access.PUT("/status/:id", h.updateStatus)
		}

		devices := router.Group(":id/device") 
		{
			devices.POST("/", h.createDevice)
			devices.DELETE("/:id", h.deleteDevice)
			devices.PUT("/:id", h.updateDevice)
			devices.POST("/:id/add", h.addHomeDevice)
			devices.POST("/:id/del", h.deleteHomeDevice)

			deviceHistory := router.Group(":id/device/:id/history") 
			{
				deviceHistory.POST("/", h.createDeviceHistory)
				deviceHistory.PUT("/:id", h.updateDeviceHistory)
				deviceHistory.GET("/:id", h.getDeviceHistory)
			}
		}
	}

	return router
}
