package logger

import (
	"log/slog"
	"os"
)

func InitLogger(debugIsEnabled bool) *slog.Logger {
	loglevel := slog.LevelError
	if debugIsEnabled {
		loglevel = slog.LevelDebug
	}

	// lgr := slog.New(
	// 	slog.NewJSONHandler(
	// 		os.Stdout,
	// 		&slog.HandlerOptions{
	// 			Level: loglevel,
	// 		},
	// 	),
	// )

	lgr := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				// AddSource: true,
				Level: loglevel,
			},
		),
	)

	slog.SetDefault(lgr)

	return lgr
}
