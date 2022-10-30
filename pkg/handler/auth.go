package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"services-front/pkg/service/dto"
)

type UserRequest struct {
	Id       int64  `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input UserRequest

	if err := c.BindJSON(&input); err != nil {
		NewExceptResp(c, http.StatusBadRequest, err.Error())
		return
	}

	userDto := dto.NewUser(input.Id, input.Username, input.Password)
	if err := h.service.Authorization.Registration(userDto); err != nil {
		NewExceptResp(c, http.StatusInternalServerError, err.Error())
	}

	logrus.Infof("регистрация пользователя %s прошла успешно", userDto.Username)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ответ": "регистрация прошла успешно",
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input UserRequest

	if err := c.BindJSON(&input); err != nil {
		NewExceptResp(c, http.StatusBadRequest, err.Error())
		return
	}

	userDto := dto.NewUser(input.Id, input.Username, input.Password)

	token, err := h.service.Authorization.ReturnToken(userDto)
	if err != nil {
		NewExceptResp(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
