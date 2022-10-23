package main

import (
	"log"
	"services-front/pkg"
	"services-front/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(pkg.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server, %s", err.Error())
	}
}
