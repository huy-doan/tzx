package application

import (
	"log"
	"os"

	"github.com/test-tzs/nomraeite/batch/infrastructure/container"
	task "github.com/test-tzs/nomraeite/batch/task/gmo-aozora"
	adapterAozora "github.com/test-tzs/nomraeite/internal/infrastructure/adapter/gmo-aozora"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/connected_service_token"
	aozoraUS "github.com/test-tzs/nomraeite/internal/usecase/gmo-aozora"
)

func Execute() {
	batchService, err := container.NewBatchContainer()
	if err != nil {
		log.Fatalf("Failed to initialize batch container: %v", err)
	}
	defer func() {
		if err := batchService.Close(); err != nil {
			log.Printf("Failed to close batch service: %v", err)
		}
	}()

	logger := batchService.Logger
	logger.Info("======= Start RefreshAozoraToken =======", nil)
	defer logger.Info("======= Stop RefreshAozoraToken =======", nil)

	repository := persistence.NewConnectedServiceTokenPersistence(batchService.DB)
	apiClient := adapterAozora.NewApiClient(logger)
	if apiClient == nil {
		logger.Error("Failed to create Aozora auth client", nil)
		os.Exit(1)
	}
	aozoraUsecase := aozoraUS.NewAuthUseCase(repository, apiClient, logger)
	reconciliationTask := task.NewRefreshGmoAozoraTokenTask(aozoraUsecase, repository, logger)
	err = reconciliationTask.Do(batchService.Ctx)
	if err != nil {
		logger.Error("Failed to refresh Aozora token", map[string]any{
			"error": err.Error(),
		})
		os.Exit(1)
	}

	logger.Info("Successfully refreshed Aozora token", nil)
}
