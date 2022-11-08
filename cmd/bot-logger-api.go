package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"services-front/configs"
	"services-front/pkg"
	"services-front/pkg/handler"
	"services-front/pkg/service"
	"services-front/pkg/storage"
	"syscall"
)

func main() {
	if err := configs.InitConfig("../configs"); err != nil {
		logrus.Fatalf("init configs error: %s", err.Error())
	}

	db, err := storage.NewConnect(viper.GetString("dbLink"))
	if err != nil {
		logrus.Fatalf("connect db error: %s", err.Error())
	}

	stor := storage.NewStorage(db)
	serv := service.NewService(stor)
	handlers := handler.NewHandler(serv)

	srv := new(pkg.Server)

	go func() {
		if err := srv.Run(viper.GetString("httpPort"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error http server, %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error on server shutting down, %s", err.Error())
	}

	if err := db.Close(context.Background()); err != nil {
		logrus.Errorf("error on db connection close, %s", err.Error())
	}
}
