package container

import (
	"context"
	"fmt"

	appConfig "github.com/test-tzs/nomraeite/internal/pkg/config"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"github.com/test-tzs/nomraeite/internal/pkg/dbconn"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"

	"gorm.io/gorm"
)

type LogBatchConfig struct {
	LogLevel     string
	LogDirectory string
}

type BatchService struct {
	Ctx       context.Context
	Logger    logger.Logger
	DB        *gorm.DB
	AppConfig *appConfig.Config
}

func DefaultLogConfig(appConfig *appConfig.Config) *LogBatchConfig {
	return &LogBatchConfig{
		LogLevel:     appConfig.LogLevel,
		LogDirectory: appConfig.LogDirectory,
	}
}

func NewBatchContainer() (*BatchService, error) {
	return NewBatchContainerWithConfig()
}

func NewBatchContainerWithConfig() (*BatchService, error) {
	appConfig := appConfig.GetConfig()
	logConfig := DefaultLogConfig(appConfig)
	cliLogger := logger.InitCLILogger(&logger.CLILoggerConfig{
		LogLevel:     logConfig.LogLevel,
		LogDirectory: logConfig.LogDirectory,
	})

	db, err := dbconn.NewConnection(cliLogger)
	if err != nil {
		cliLogger.Error("Failed to connect to database", map[string]any{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	ctx := context.Background()
	ctx, dbSetErr := database.SetDB(ctx, db)
	if dbSetErr != nil {
		cliLogger.Error("Failed to set DB in context:", map[string]any{
			"error": dbSetErr.Error(),
		})
		return nil, fmt.Errorf("failed to set DB in context: %w", dbSetErr)
	}

	return &BatchService{
		DB:        db,
		Ctx:       ctx,
		Logger:    cliLogger,
		AppConfig: appConfig,
	}, nil
}

// Close releases resources when finished
func (s *BatchService) Close() error {
	sqlDB, err := s.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
