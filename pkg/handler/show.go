package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) messages(c *gin.Context) {
	messages, err := h.service.Show.ShowAllMessages()
	if err != nil {
		NewExceptResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, messages)
}
