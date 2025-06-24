package logger

import (
	"fmt"
	"io"
	"maps"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	appConf "github.com/test-tzs/nomraeite/internal/pkg/config"
)

// LogLevel represents logging levels
type LogLevel string

const (
	// DEBUG level for detailed information in development environment
	DEBUG LogLevel = "debug"
	// INFO level for general operational information
	INFO LogLevel = "info"
	// WARN level for warnings that don't cause errors but should be noted
	WARN LogLevel = "warn"
	// ERROR level for system errors
	ERROR LogLevel = "error"
)

// TraceIDKey is the context key for trace ID
const TraceIDKey = "trace_id"

// loggerImpl is the implementation of Logger interface
type loggerImpl struct {
	logger       *logrus.Logger
	traceID      string
	extraFields  map[string]any
	logDirectory string
	filePrefix   string // Prefix for log files, empty for web, "cli" for CLI
}

// Config holds the configuration for logger
type Config struct {
	LogLevel     string
	LogDirectory string
	EnableSQLLog bool
}

// Global singleton instance and initialization lock
var (
	instance      Logger
	instanceOnce  sync.Once
	instanceMutex sync.RWMutex
)

func setLogOutput(logger *logrus.Logger, config *Config) *logrus.Logger {
	if !appConf.IsLocal() {
		logger.SetOutput(os.Stdout)
		return logger
	}

	// Create log directory
	if err := os.MkdirAll(config.LogDirectory, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create log directory: %v\n", err)
		logger.SetOutput(os.Stdout)
		return logger
	}

	// Open log file
	logFile, err := os.OpenFile(filepath.Join(config.LogDirectory, "app.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
		logger.SetOutput(os.Stdout)
		return logger
	}

	// Set multi-writer for stdout and log file
	mw := io.MultiWriter(os.Stdout, logFile)
	logger.SetOutput(mw)
	return logger
}

// InitLogger initializes the global logger instance with the given config
// This should be called once at application startup
func InitLogger(config *Config) {
	instanceOnce.Do(func() {
		logger := logrus.New()

		level, err := logrus.ParseLevel(config.LogLevel)
		if err != nil {
			level = logrus.InfoLevel
		}
		logger.SetLevel(level)

		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})

		logger = setLogOutput(logger, config)

		instance = &loggerImpl{
			logger:      logger,
			traceID:     GenerateTraceID(),
			extraFields: make(map[string]any),
		}
	})
}

// GetLogger returns the global logger instance
// If the logger hasn't been initialized, it returns a default logger
func GetLogger() Logger {
	instanceMutex.RLock()
	defer instanceMutex.RUnlock()

	if instance == nil {
		appConf := appConf.LoadConfig()

		defaultConfig := &Config{
			LogLevel:     appConf.LogLevel,
			LogDirectory: appConf.LogDirectory,
			EnableSQLLog: appConf.EnableSQLLog,
		}

		return NewLogger(defaultConfig)
	}

	return instance
}

// NewLogger creates a new logger instance
// This should be used only for specific cases where a separate logger instance is needed
func NewLogger(config *Config) Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	setLogOutput(logger, config)

	return &loggerImpl{
		logger:      logger,
		traceID:     GenerateTraceID(),
		extraFields: make(map[string]any),
	}
}

// GenerateTraceID generates a unique trace ID for request tracking
func GenerateTraceID() string {
	return uuid.New().String()
}

// WithTraceID creates a new logger instance with the specified trace ID
func (l *loggerImpl) WithTraceID(traceID string) Logger {
	newLogger := &loggerImpl{
		logger:       l.logger,
		traceID:      traceID,
		extraFields:  make(map[string]any),
		logDirectory: l.logDirectory,
		filePrefix:   l.filePrefix,
	}

	// Copy extra fields
	maps.Copy(newLogger.extraFields, l.extraFields)

	return newLogger
}

// WithField adds a field to the logger
func (l *loggerImpl) WithField(key string, value any) Logger {
	newLogger := &loggerImpl{
		logger:       l.logger,
		traceID:      l.traceID,
		extraFields:  make(map[string]any),
		logDirectory: l.logDirectory,
		filePrefix:   l.filePrefix,
	}

	// Copy existing extra fields
	maps.Copy(newLogger.extraFields, l.extraFields)

	// Add new field
	newLogger.extraFields[key] = value

	return newLogger
}

// WithFields adds multiple fields to the logger
func (l *loggerImpl) WithFields(fields map[string]any) Logger {
	newLogger := &loggerImpl{
		logger:       l.logger,
		traceID:      l.traceID,
		extraFields:  make(map[string]any),
		logDirectory: l.logDirectory,
		filePrefix:   l.filePrefix,
	}

	// Copy existing extra fields
	maps.Copy(newLogger.extraFields, l.extraFields)

	// Add new fields
	maps.Copy(newLogger.extraFields, fields)

	return newLogger
}

// GetTraceID returns the current trace ID
func (l *loggerImpl) GetTraceID() string {
	return l.traceID
}

// makeFields adds common fields to all log entries
func (l *loggerImpl) makeFields(fields map[string]any) logrus.Fields {
	if fields == nil {
		fields = make(map[string]any)
	}

	// Add trace ID and timestamp
	fields[TraceIDKey] = l.traceID
	fields["timestamp"] = time.Now().UTC().Format(time.RFC3339)

	// Add extra fields
	maps.Copy(fields, l.extraFields)

	return logrus.Fields(fields)
}

// Debug logs a message at the DEBUG level
func (l *loggerImpl) Debug(msg string, fields map[string]any) {
	if l.logger.IsLevelEnabled(logrus.DebugLevel) {
		entry := &logrus.Entry{
			Logger:  l.logger,
			Data:    l.makeFields(fields),
			Time:    time.Now(),
			Level:   logrus.DebugLevel,
			Message: msg,
		}

		l.logger.WithFields(entry.Data).Debug(msg)
	}
}

// Info logs a message at the INFO level
func (l *loggerImpl) Info(msg string, fields map[string]any) {
	if l.logger.IsLevelEnabled(logrus.InfoLevel) {
		entry := &logrus.Entry{
			Logger:  l.logger,
			Data:    l.makeFields(fields),
			Time:    time.Now(),
			Level:   logrus.InfoLevel,
			Message: msg,
		}

		l.logger.WithFields(entry.Data).Info(msg)
	}
}

// Warn logs a message at the WARN level
func (l *loggerImpl) Warn(msg string, fields map[string]any) {
	if l.logger.IsLevelEnabled(logrus.WarnLevel) {
		entry := &logrus.Entry{
			Logger:  l.logger,
			Data:    l.makeFields(fields),
			Time:    time.Now(),
			Level:   logrus.WarnLevel,
			Message: msg,
		}

		l.logger.WithFields(entry.Data).Warn(msg)
	}
}

// Error logs a message at the ERROR level
func (l *loggerImpl) Error(msg string, fields map[string]any) {
	if l.logger.IsLevelEnabled(logrus.ErrorLevel) {
		entry := &logrus.Entry{
			Logger:  l.logger,
			Data:    l.makeFields(fields),
			Time:    time.Now(),
			Level:   logrus.ErrorLevel,
			Message: msg,
		}

		l.logger.WithFields(entry.Data).Error(msg)
	}
}
