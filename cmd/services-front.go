package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"services-front/config"
	"services-front/pkg"
	"services-front/pkg/handler"
	"services-front/pkg/service"
	"services-front/pkg/storage"
)

func main() {
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("init config error: %s", err.Error())
	}

	db, err := storage.NewClient(viper.GetString("db"))
	if err != nil {
		logrus.Fatalf("init db error, %s", err.Error())
	}

	storages := storage.NewStorage(&db)
	services := service.NewService(storages)
	handlers := handler.NewHandler(services)

	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("httpPort"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error http server, %s", err.Error())
	}
}
