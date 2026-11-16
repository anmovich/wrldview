package loger

import (
	"io"
	"log/slog"
	"os"
)

func New() {
	logpath := os.Getenv("LOGPATH")
	file, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		slog.Error("error", err.Error(), "Error in log")
	}

	logger := slog.New(slog.NewTextHandler(io.MultiWriter(os.Stdout, file), nil))

	slog.SetDefault(logger)
	slog.Debug("Logger start correctly")
}
