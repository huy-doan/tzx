package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	appConf "github.com/test-tzs/nomraeite/internal/pkg/config"
)

type CLILoggerConfig struct {
	LogLevel     string
	LogDirectory string
}

func setCliLogOutput(logger *logrus.Logger, config *CLILoggerConfig) *logrus.Logger {
	if !appConf.IsDevelopment() {
		logger.SetOutput(os.Stdout)
		return logger
	}

	// Create log directory
	if err := os.MkdirAll(config.LogDirectory, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create log directory: %v\n", err)
		logger.SetOutput(os.Stdout)
		return logger
	}

	// Create batch log directory
	batchLogDir := filepath.Join(config.LogDirectory, "batch")
	if err := os.MkdirAll(batchLogDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create batch log directory: %v\n", err)
		logger.SetOutput(os.Stdout)
		return logger
	}

	// Open CLI log file
	cliLogPath := filepath.Join(batchLogDir, "cli.log")
	file, err := os.OpenFile(cliLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open CLI log file: %v\n", err)
		logger.SetOutput(os.Stdout)
		return logger
	}

	// Set multi-writer for stdout and log file
	mw := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(mw)
	return logger
}

func InitCLILogger(config *CLILoggerConfig) Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	logger = setCliLogOutput(logger, config)

	return &loggerImpl{
		logger:       logger,
		traceID:      GenerateTraceID(),
		extraFields:  make(map[string]any),
		logDirectory: config.LogDirectory,
		filePrefix:   "cli",
	}
}
