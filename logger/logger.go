package logger

import (
	"log/slog"
)

func InitLogger(debugIsEnabled bool) *slog.Logger {
	loglevel := slog.LevelError
	if debugIsEnabled {
		loglevel = slog.LevelDebug
	}

	lgr := slog.New(slog.Default().Handler())

	slog.SetLogLoggerLevel(loglevel)
	slog.SetDefault(lgr)

	return lgr
}
