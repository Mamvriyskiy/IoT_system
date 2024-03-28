package src

import (
	"github.com/gin-gonic/gin"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/src/user"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/src/handler"
)

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}
