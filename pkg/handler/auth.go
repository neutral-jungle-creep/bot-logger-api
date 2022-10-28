package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"services-front/pkg/service"
	"services-front/pkg/service/dto"
	"services-front/pkg/storage"
)

type UserRequest struct {
	Id       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	resp := authService.LogIn(userDto)
	c.Error(resp)
}
