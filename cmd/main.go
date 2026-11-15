package main

import (
	"log/slog"

	"github.com/anmovich/wrldview/loger"
	"github.com/anmovich/wrldview/pkg/config"
)

func main(){
	config.InitConfig()
	slog.Info("config init")
	loger.New()
	slog.Info("logger init")
}
