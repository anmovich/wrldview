package config

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func InitConfig(){
	godotenv.Load(".env")
	fmt.Println(os.Getenv("LOGFILE"))
}
