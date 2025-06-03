package application

import (
	"fmt"
	"os"
	"time"

	"github.com/makeshop-jp/master-console/batch/infrastructure/container"
	payoutPersistence "github.com/makeshop-jp/master-console/batch/infrastructure/persistence/payout"
	transactionPersistence "github.com/makeshop-jp/master-console/batch/infrastructure/persistence/transaction"
	gmoAozoraTask "github.com/makeshop-jp/master-console/batch/task/gmo-aozora"
	gmoAozorUC "github.com/makeshop-jp/master-console/batch/usecase/gmo-aozora"
	adapterAozora "github.com/makeshop-jp/master-console/internal/infrastructure/adapter/gmo-aozora"
	connectedServiceTokenPersistence "github.com/makeshop-jp/master-console/internal/infrastructure/persistence/connected_service_token"
	"github.com/makeshop-jp/master-console/internal/pkg/database"
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

	logger.Info("======= Start Aozora Transfer Request Batch =======", nil)
	defer logger.Info("======= End Aozora Transfer Request Batch =======", nil)

	ctx := batchService.Ctx
	ctx, dbSetErr := database.SetDB(ctx, db)
	if dbSetErr != nil {
		logger.Error("Failed to set DB in context", map[string]any{
			"error": dbSetErr.Error(),
		})
		os.Exit(1)
	}

	start := time.Now()
	payoutRepository := payoutPersistence.NewPayoutPersistence(db)
	payoutRecordRepository := payoutPersistence.NewPayoutRecordPersistence(db)
	transactionRepository := transactionPersistence.NewTransactionPersistence(db)
	connectedServiceTokenRepository := connectedServiceTokenPersistence.NewConnectedServiceTokenPersistence(db)
	apiClient := adapterAozora.NewApiClient(logger)

	usecase := gmoAozorUC.NewTransferRequestUsecase(
		logger,
		payoutRepository,
		payoutRecordRepository,
		transactionRepository,
		apiClient,
	)

	task := gmoAozoraTask.NewTransferRequestTask(
		usecase,
		connectedServiceTokenRepository,
		payoutRepository,
		payoutRecordRepository,
		logger,
	)
	err = task.Do(ctx)
	if err != nil {
		logger.Error("Failed to execute transfer request usecase", map[string]any{
			"error": err.Error(),
		})
		os.Exit(1)
	}

	elapsed := time.Since(start)
	logger.Info(fmt.Sprintf("Batch completed successfully in %s", elapsed), nil)
}
