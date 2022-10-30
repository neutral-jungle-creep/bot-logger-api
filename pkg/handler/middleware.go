package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	userCtx     = "userId"
	authHandler = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	handler := c.GetHeader(authHandler)
	if handler == "" {
		logrus.Info(c.GetHeader(authHandler), h.service.Authorization)
		NewExceptResp(c, http.StatusUnauthorized, "ошибка1 авторизации")
		return
	}

	headerParts := strings.Split(handler, " ")
	if len(headerParts) != 2 {
		NewExceptResp(c, http.StatusUnauthorized, "ошибка2 авторизации")
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		NewExceptResp(c, http.StatusUnauthorized, "ошибка3 авторизации")
		return
	}

	c.Set(userCtx, userId)
}
