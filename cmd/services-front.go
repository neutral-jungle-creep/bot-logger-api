package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"services-front/config"
	"services-front/pkg"
	"services-front/pkg/handler"
)

func main() {
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("init config error: %s", err.Error())
	}

	handlers := handler.NewHandler()

	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("httpPort"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error http server, %s", err.Error())
	}
}
