package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		api.GET("/show-messages")
	}
	return router
}
