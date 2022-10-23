package main

import (
	"log"
	"services-front/pkg"
	"services-front/pkg/handler"
	"services-front/pkg/service"
	"services-front/pkg/storage"
)

func main() {
	storages := storage.NewStorage()
	services := service.NewService(storages)
	handlers := handler.NewHandler(services)

	srv := new(pkg.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server, %s", err.Error())
	}
}
