package main

import (
	"github.com/spf13/viper"
	"log"
	"services-front/config"
	"services-front/pkg"
	"services-front/pkg/handler"
	"services-front/pkg/service"
	"services-front/pkg/storage"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("init config error: %s", err.Error())
	}

	storages := storage.NewStorage()
	services := service.NewService(storages)
	handlers := handler.NewHandler(services)

	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server, %s", err.Error())
	}
}
