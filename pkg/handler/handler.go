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
		home.POST("/create-home", h.createHome)
		home.DELETE("/delete-home", h.deleteHome)
		home.PUT("/update-home", h.updateHome)
	}

	return router
}
