package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"services-front/configs"
	"services-front/pkg"
	"services-front/pkg/handler"
)

func main() {
	if err := configs.InitConfig("../configs"); err != nil {
		logrus.Fatalf("init configs error: %s", err.Error())
	}

	handlers := handler.NewHandler()

	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("httpPort"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error http server, %s", err.Error())
	}
}
