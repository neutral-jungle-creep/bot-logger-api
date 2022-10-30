package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) messages(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ответ": userId,
	})
}
