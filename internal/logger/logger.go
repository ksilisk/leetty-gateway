package logger

import (
	"log/slog"
	"os"
)

const loggerLevelEnvVar = "LEETTY_GATEWAY_LOGGER_LEVEL"

var Logger = slog.New(slog.NewJSONHandler(os.Stdout, getLoggerOpts()))

func getLoggerOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true,
		Level:     getLoggerLevel(),
	}
}

func getLoggerLevel() slog.Level {
	var path, present = os.LookupEnv(loggerLevelEnvVar)
	if !present {
		return slog.LevelDebug
	}
	switch path {
	case "DEBUG", "debug":
		return slog.LevelDebug
	case "INFO", "info":
		return slog.LevelInfo
	case "WARNING", "warning":
		return slog.LevelWarn
	case "ERROR", "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
