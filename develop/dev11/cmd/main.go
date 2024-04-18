package main

import (
	"dev11/cmd/initStart"
	"dev11/internal/handler"
	"dev11/internal/repository"
	"dev11/internal/server"
	"dev11/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"log"
)

func main() {
	if err := initStart.InitConfig(); err != nil {
		log.Fatalf("Error init config: %v\n", err)
		return
	}
	db := initStart.InitDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.Routers()); err != nil {
		log.Fatalf("Не запустился сервер")
		return
	}

}
