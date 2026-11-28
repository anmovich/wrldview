package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/anmovich/wrldview/internal/repository/postgres"
	"github.com/anmovich/wrldview/internal/service"
	"github.com/anmovich/wrldview/internal/transport"
	"github.com/anmovich/wrldview/pkg/config"
	"github.com/anmovich/wrldview/pkg/loger"
	"github.com/anmovich/wrldview/pkg/storage/garage"
)

func main() {
	//Init .env
	config.InitConfig()
	slog.Info("config init")

	//Init loger
	loger.New()
	slog.Info("logger init")

	//Init file storage
	storage := garage.NewGarage(os.Getenv("GARAGELINK"), os.Getenv("ACCESSEDKEYGARAGE"), os.Getenv("SECRETKEYGARAGE"), "bookcover")

	//Init repository
	repository := postgres.New(context.Background())

	//Init service
	service := service.New(*repository, storage)

	//Init handlers
	handler := transport.New(*service)

	//Init server and run
	server := transport.NewServer(handler)
	server.Run()

}
