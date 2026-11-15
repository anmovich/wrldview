package main

import (
	"log/slog"

	"github.com/anmovich/wrldview/loger"
)

func main(){
	loger.New()
	slog.Info("Main")
	slog.Debug("Main start work")
}
