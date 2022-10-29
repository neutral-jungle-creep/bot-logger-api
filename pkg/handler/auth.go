package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"services-front/pkg/service"
	"services-front/pkg/service/dto"
	"services-front/pkg/storage"
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

	db, err := storage.NewConnect(viper.GetString("dbLink"))
	if err != nil {
		c.Error(err)
		return
	}
	defer db.Close(context.Background())

	authStorage := storage.NewPgAuthStorage(db)
	authService := service.NewAuthService(authStorage)
	userDto := dto.NewUser(input.Id, input.Username, input.Password)
	if err := authService.Registration(userDto); err != nil {
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

	db, err := storage.NewConnect(viper.GetString("dbLink"))
	if err != nil {
		c.Error(err)
		return
	}
	defer db.Close(context.Background())

	authStorage := storage.NewPgAuthStorage(db)
	authService := service.NewAuthService(authStorage)
	userDto := dto.NewUser(input.Id, input.Username, input.Password)

	token, err := authService.ReturnToken(userDto)
	if err != nil {
		NewExceptResp(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
