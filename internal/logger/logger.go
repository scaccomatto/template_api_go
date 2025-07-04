package logger

import (
	"log/slog"
	"os"
)

var L *slog.Logger

func init() {
	// You could use slog.NewJSONHandler for machine-friendly logs
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		//AddSource: true,
	})
	L = slog.New(handler)

	slog.SetDefault(L)
}
