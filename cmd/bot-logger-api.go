package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"services-front/configs"
	"services-front/pkg"
	"services-front/pkg/handler"
	"services-front/pkg/service"
	"services-front/pkg/storage"
)

func main() {
	if err := configs.InitConfig("../configs"); err != nil {
		logrus.Fatalf("init configs error: %s", err.Error())
	}

	db, err := storage.NewConnect(viper.GetString("dbLink"))
	if err != nil {
		logrus.Fatalf("connect db error: %s", err.Error())
	}
	defer db.Close(context.Background())

	stor := storage.NewStorage(db)
	serv := service.NewService(stor)
	handlers := handler.NewHandler(serv)

	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("httpPort"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error http server, %s", err.Error())
	}
}
