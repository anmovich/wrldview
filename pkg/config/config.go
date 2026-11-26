package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println(os.Getenv("LOGPATH"))
}
