package application

import (
	"fmt"
	"os"
	"time"

	gmoAozoraService "github.com/test-tzs/nomraeite/batch/domain/service/gmo-aozora"
	"github.com/test-tzs/nomraeite/batch/infrastructure/container"
	gmoAozoraTask "github.com/test-tzs/nomraeite/batch/task/gmo-aozora"
	gmoAozoraUC "github.com/test-tzs/nomraeite/batch/usecase/gmo-aozora"
	adapterAozora "github.com/test-tzs/nomraeite/internal/infrastructure/adapter/gmo-aozora"
	adapterSQS "github.com/test-tzs/nomraeite/internal/infrastructure/adapter/sqs"
	connectedServiceTokenPersistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/connected_service_token"
	prTransferredPersistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout"

	adapterUtil "github.com/test-tzs/nomraeite/internal/infrastructure/adapter/util"
)

func Execute() {
	batchService, err := container.NewBatchContainer()
	if err != nil {
		fmt.Printf("Failed to initialize batch container: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if closeErr := batchService.Close(); closeErr != nil {
			fmt.Printf("Failed to close batch service: %v\n", closeErr)
			os.Exit(1)
		}
	}()

	logger := batchService.Logger
	db := batchService.DB

	logger.Info("======= Start Aozora Transfer Status Batch =======", nil)
	defer logger.Info("======= End Aozora Transfer Status Batch =======", nil)

	start := time.Now()

	prTransferredRepository := prTransferredPersistence.NewPayoutRecordTransferredPersistence(db)
	connectedServiceTokenRepository := connectedServiceTokenPersistence.NewConnectedServiceTokenPersistence(db)

	apiClient := adapterAozora.NewApiClient(logger)
	sqsClient, sqsErr := adapterSQS.NewSQSAdapter(adapterSQS.SQSAdapterConfig{
		QueueURL: batchService.AppConfig.MakeshopSQSQueueURL,
		AWSConfig: adapterUtil.AWSConfig{
			AccessKeyID:     batchService.AppConfig.AwsAccessKeyID,
			SecretAccessKey: batchService.AppConfig.AwsSecretAccessKey,
			Region:          batchService.AppConfig.AwsRegion,
		},
	})
	if sqsErr != nil {
		logger.Error("Failed to initialize SQS adapter", map[string]any{
			"error": sqsErr.Error(),
		})
		os.Exit(1)
	}

	transferStatusService := gmoAozoraService.NewTransferStatusService(
		logger,
		apiClient,
		sqsClient,
		prTransferredRepository,
	)

	transferStatusUsecase := gmoAozoraUC.NewTransferStatusUsecase(
		logger,
		prTransferredRepository,
		connectedServiceTokenRepository,
		apiClient,
		transferStatusService,
	)

	task := gmoAozoraTask.NewTransferStatusTask(
		transferStatusUsecase,
		prTransferredRepository,
		connectedServiceTokenRepository,
		logger,
		batchService.AppConfig,
	)

	err = task.Do(batchService.Ctx)
	if err != nil {
		logger.Error("Failed to execute transfer status task", map[string]any{
			"error": err.Error(),
		})
		os.Exit(1)
	}

	elapsed := time.Since(start)
	logger.Info(fmt.Sprintf("Batch completed successfully in %s", elapsed), nil)
}
