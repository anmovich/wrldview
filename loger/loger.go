package loger

import (
	"log/slog"
	"os"

)

func New() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Debug("Logger start correctly")
}
