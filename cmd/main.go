package main

import (
	"log/slog"

	"github.com/anmovich/wrldview/internal/repository"
	"github.com/anmovich/wrldview/internal/service"
	"github.com/anmovich/wrldview/internal/transport"
	"github.com/anmovich/wrldview/pkg/config"
	"github.com/anmovich/wrldview/pkg/loger"
)

func main() {
	config.InitConfig()
	slog.Info("config init")
	loger.New()
	slog.Info("logger init")
	repository := repository.New()
	service := service.New(*repository)
	handler := transport.New(*service)
	server := transport.NewServer(handler)
	server.Run()
}
