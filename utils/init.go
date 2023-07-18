// Description: Generic utils functions
// Author: Pixie79
// ============================================================================
// package utils

package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

var (
	Logger      *slog.Logger    // Logger is the default logger
	logLevel    string          // logLevel is the log level
	Err         error           // Err is the default error
	Hostname, _ = os.Hostname() // Hostname is the default hostname
)

// init loads the .env file and sets the log level
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		slog.Debug("No .env file found")
	}
	logLevel = GetEnv("LOG_LEVEL", "INFO")
	Logger = initLog()
}

// initLog initializes the logger
func initLog() *slog.Logger {
	switch logLevel {
	case "DEBUG":
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
		logger.Info(fmt.Sprintf("Loglevel set to: %s", logLevel))
		slog.SetDefault(logger)
		return logger
	default:
		opts := &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
		logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
		logger.Info(fmt.Sprintf("Loglevel set to: %s", logLevel))
		slog.SetDefault(logger)
		return logger
	}
}
