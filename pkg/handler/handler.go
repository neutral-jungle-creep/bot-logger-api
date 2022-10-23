package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/show")
	{
		api.GET("/messages", h.messages)
	}
	return router
}
