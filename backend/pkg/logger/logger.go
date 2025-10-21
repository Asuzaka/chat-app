package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func Init(isDev bool) {
	// Global time format
	zerolog.TimeFieldFormat = time.RFC3339

	if isDev {
		// Human-readable console output for local dev
		Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05", // black magic! #learnbetter
		})
		Logger.Info().Msg("Logger initialized in development mode")
	} else {
		// JSON structured logs for production
		Logger = log.With().Timestamp().Logger()
		Logger.Info().Msg("Logger initialized in production mode")
	}
}

// Shortcuts
func Info(msg string)  { Logger.Info().Msg(msg) }
func Warn(msg string)  { Logger.Warn().Msg(msg) }
func Error(msg string) { Logger.Error().Msg(msg) }
func Debug(msg string) { Logger.Debug().Msg(msg) }
