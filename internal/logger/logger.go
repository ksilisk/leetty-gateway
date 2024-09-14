package logger

import (
	"log/slog"
	"os"
	"strings"
)

const loggerLevelEnvVar = "LEETTY_GATEWAY_LOGGER_LEVEL"

var Logger = slog.New(slog.NewJSONHandler(os.Stdout, getLoggerOpts()))

const (
	LogLevelDebug = "DEBUG"
	LogLevelInfo  = "INFO"
	LogLevelWarn  = "WARN"
	LogLevelError = "ERROR"
)

func getLoggerOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true,
		Level:     getLoggerLevel(),
	}
}

func getLoggerLevel() slog.Level {
	var path, present = os.LookupEnv(loggerLevelEnvVar)
	if !present {
		return slog.LevelInfo
	}
	switch strings.ToUpper(path) {
	case LogLevelDebug:
		return slog.LevelDebug
	case LogLevelInfo:
		return slog.LevelInfo
	case LogLevelWarn:
		return slog.LevelWarn
	case LogLevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
