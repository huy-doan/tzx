package service

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	modelApproval "github.com/test-tzs/nomraeite/internal/domain/model/approval"
	modelApprovalStage "github.com/test-tzs/nomraeite/internal/domain/model/approval_stage"
	modelApprovalWorkflowStage "github.com/test-tzs/nomraeite/internal/domain/model/approval_workflow_stage"
	payoutModel "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	transactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	objectApproval "github.com/test-tzs/nomraeite/internal/domain/object/approval"
	objectApprovalStage "github.com/test-tzs/nomraeite/internal/domain/object/approval_stage"
	objectApprovalWorkflow "github.com/test-tzs/nomraeite/internal/domain/object/approval_workflow"
	bankAccountObject "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"
	objectPayout "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	transactionRecordTypeObject "github.com/test-tzs/nomraeite/internal/domain/object/transaction_record"
	approvalRepo "github.com/test-tzs/nomraeite/internal/domain/repository/approval"
	approvalStageRepo "github.com/test-tzs/nomraeite/internal/domain/repository/approval_stage"
	approvalWorkflowStageRepo "github.com/test-tzs/nomraeite/internal/domain/repository/approval_workflow_stage"
	payoutRepo "github.com/test-tzs/nomraeite/internal/domain/repository/payout"
	transactionRepo "github.com/test-tzs/nomraeite/internal/domain/repository/transaction"
	userRepo "github.com/test-tzs/nomraeite/internal/domain/repository/user"
	email "github.com/test-tzs/nomraeite/internal/infrastructure/adapter/email"
	config "github.com/test-tzs/nomraeite/internal/pkg/config"
	logger "github.com/test-tzs/nomraeite/internal/pkg/logger"
)

type ApprovalWorkflowService interface {
	ApprovePayout(ctx context.Context, input *inputdata.TransferApprovalInputData, currentStage *modelApprovalWorkflowStage.ApprovalWorkflowStage) *modelApproval.ApprovalResult
	ValidateUserPermissions(ctx context.Context, userID int, stageID int) error
	GetApprovalWorkflowStageForPayout(ctx context.Context, payoutID int) (*modelApprovalWorkflowStage.ApprovalWorkflowStage, error)
	SendMailToApprover(ctx context.Context, payoutID int, workflowID int, stageID int) error
}

type approvalWorkflowServiceImpl struct {
	approvalRepo          approvalRepo.ApprovalRepository
	approvalStageRepo     approvalStageRepo.ApprovalStageRepository
	workflowStageRepo     approvalWorkflowStageRepo.ApprovalWorkflowStageRepository
	payoutRepo            payoutRepo.PayoutRepository
	payoutRecordRepo      payoutRepo.PayoutRecordRepository
	userRepo              userRepo.UserRepository
	transactionRepo       transactionRepo.TransactionRepository
	transactionRecordRepo transactionRepo.TransactionRecordRepository
	mailService           *email.MailService
	logger                logger.Logger
	config                *config.Config
}

func NewApprovalWorkflowService(
	approvalRepository approvalRepo.ApprovalRepository,
	approvalStageRepository approvalStageRepo.ApprovalStageRepository,
	workflowStageRepository approvalWorkflowStageRepo.ApprovalWorkflowStageRepository,
	payoutRepository payoutRepo.PayoutRepository,
	payoutRecordRepository payoutRepo.PayoutRecordRepository,
	userRepository userRepo.UserRepository,
	transactionRepository transactionRepo.TransactionRepository,
	transactionRecordRepository transactionRepo.TransactionRecordRepository,
	mailService *email.MailService,
) ApprovalWorkflowService {
	return &approvalWorkflowServiceImpl{
		approvalRepo:          approvalRepository,
		approvalStageRepo:     approvalStageRepository,
		workflowStageRepo:     workflowStageRepository,
		payoutRepo:            payoutRepository,
		payoutRecordRepo:      payoutRecordRepository,
		userRepo:              userRepository,
		transactionRepo:       transactionRepository,
		transactionRecordRepo: transactionRecordRepository,
		mailService:           mailService,
		logger:                logger.GetLogger(),
		config:                config.GetConfig(),
	}
}

func (s *approvalWorkflowServiceImpl) ApprovePayout(
	ctx context.Context,
	input *inputdata.TransferApprovalInputData,
	currentStage *modelApprovalWorkflowStage.ApprovalWorkflowStage,
) *modelApproval.ApprovalResult {
	payoutID := input.PayoutID
	userID := input.UserID
	action := objectApproval.ApprovalAction(input.Action)
	approvalObj, err := s.approvalRepo.GetByPayoutID(ctx, payoutID)
	if err != nil {
		return &modelApproval.ApprovalResult{
			PayoutID: payoutID,
			Error:    fmt.Sprintf("failed to get approval: %v", err),
		}
	}

	// Skip approval if the payout has been rejected
	if approvalObj.ApprovalStatus == objectApproval.ApprovalStatusRejected {
		return &modelApproval.ApprovalResult{
			PayoutID:     payoutID,
			ApprovalID:   approvalObj.ID,
			Status:       "Rejected",
			CurrentStage: currentStage.Level,
			Error:        "Approval has been rejected",
		}
	}

	// Create approval stage
	if result := s.createApprovalStage(
		ctx, approvalObj, currentStage, userID, action, input.Note, payoutID,
	); result != nil {
		return result
	}

	// Calculate new approval status
	newStatus, nextStage := s.calculateNewApprovalStatus(ctx, approvalObj, action, currentStage)
	approvalObj.SetStatus(newStatus)

	if err := s.approvalRepo.Update(ctx, approvalObj); err != nil {
		return &modelApproval.ApprovalResult{
			PayoutID:   payoutID,
			ApprovalID: approvalObj.ID,
			Error:      fmt.Sprintf("failed to update approval status: %v", err),
		}
	}

	// Update payout with approval status
	if result := s.updatePayout(ctx, payoutID, approvalObj.ID, currentStage.Level, action); result != nil {
		return result
	}

	isLatestStage := nextStage == nil
	if isLatestStage && newStatus == objectApproval.ApprovalStatusApproved && action == objectApproval.ApprovalActionApproved {
		if err := s.payoutRecordRepo.UpdateTransferStatusByPayoutID(ctx, payoutID, objectPayout.PayoutRecordStatusWaitingTransfer); err != nil {
			return &modelApproval.ApprovalResult{
				PayoutID:   payoutID,
				ApprovalID: approvalObj.ID,
				Error:      fmt.Sprintf("failed to update payout record transfer status: %v", err),
			}
		}
	}

	if action == objectApproval.ApprovalActionRejected {
		if err := s.createTransactionsFromPayout(ctx, payoutID); err != nil {
			return &modelApproval.ApprovalResult{
				PayoutID:   payoutID,
				ApprovalID: approvalObj.ID,
				Error:      fmt.Sprintf("failed to create new transactions for rejected payout: %v", err),
			}
		}
	}

	return &modelApproval.ApprovalResult{
		PayoutID:     payoutID,
		ApprovalID:   approvalObj.ID,
		Status:       newStatus.String(),
		CurrentStage: currentStage.Level,
		NextStage:    nextStage,
	}
}

func (s *approvalWorkflowServiceImpl) calculateNewApprovalStatus(
	ctx context.Context,
	approvalObj *modelApproval.Approval,
	action objectApproval.ApprovalAction,
	currentStage *modelApprovalWorkflowStage.ApprovalWorkflowStage,
) (objectApproval.ApprovalStatus, *int) {
	if action == objectApproval.ApprovalActionRejected {
		return objectApproval.ApprovalStatusRejected, nil
	}

	nextStage, err := s.workflowStageRepo.GetNextStage(ctx, currentStage)
	if err != nil || nextStage == nil {
		return objectApproval.ApprovalStatusApproved, nil
	}

	return objectApproval.ApprovalStatusWaitApproval, &nextStage.Level
}

func (s *approvalWorkflowServiceImpl) ValidateUserPermissions(ctx context.Context, userID, stageID int) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to find user %d: %w", userID, err)
	}

	stage, err := s.workflowStageRepo.GetByID(ctx, stageID)
	if err != nil {
		return fmt.Errorf("failed to get workflow stage %d: %w", stageID, err)
	}

	if !user.HasPermission(stage.PermissionID) {
		return fmt.Errorf("user %d does not have permission %d for stage %d", userID, stage.PermissionID, stageID)
	}

	return nil
}

// Get current stage for payout
func (s *approvalWorkflowServiceImpl) GetApprovalWorkflowStageForPayout(ctx context.Context, payoutID int) (*modelApprovalWorkflowStage.ApprovalWorkflowStage, error) {
	payout, err := s.payoutRepo.GetByID(ctx, payoutID)
	if err != nil && payout == nil {
		return nil, fmt.Errorf("failed to get payout for payout ID %d: %w", payoutID, err)
	}

	approval, err := s.approvalRepo.GetByPayoutID(ctx, payoutID)
	if err != nil {
		return nil, fmt.Errorf("failed to get approval for payout %d: %w", payoutID, err)
	}

	if approval.ApprovalStatus == objectApproval.ApprovalStatusPending {
		return s.workflowStageRepo.GetApprovalWorkflowStage(ctx, objectApprovalWorkflow.TRANSFER_APPROVAL_WORKFLOW, objectApprovalStage.TRANSFER_APPROVAL_FIRST_LEVEL)
	}

	currentStage, err := s.approvalStageRepo.GetCurrentStageByApprovalID(ctx, approval.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get current stage: %w", err)
	}

	currentWorkflowStage, err := s.workflowStageRepo.GetByID(ctx, currentStage.ApprovalWorkflowStageID)
	if err != nil {
		return nil, fmt.Errorf("failed to get current workflow stage: %w", err)
	}

	nextStage, err := s.workflowStageRepo.GetNextStage(ctx, currentWorkflowStage)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to get next stage: %w", err)
	}

	if nextStage == nil {
		return nil, fmt.Errorf("completed approval workflow")
	}

	return nextStage, nil
}

// Update payout with approval ID and status based on current stage
func (s *approvalWorkflowServiceImpl) updatePayout(ctx context.Context, payoutID int, approvalID int, stageLevel int, action objectApproval.ApprovalAction) *modelApproval.ApprovalResult {
	payoutModelApproval := &payoutModel.Payout{
		ID:         payoutID,
		ApprovalID: &approvalID,
	}

	if action == objectApproval.ApprovalActionApproved {
		switch stageLevel {
		case 1:
			payoutModelApproval.PayoutStatus = objectPayout.PayoutStatusApproving
		case 2:
			payoutModelApproval.PayoutStatus = objectPayout.PayoutStatusWaitingTransfer
		}
	}

	if err := s.payoutRepo.Update(ctx, payoutModelApproval); err != nil {
		return &modelApproval.ApprovalResult{
			PayoutID:   payoutID,
			ApprovalID: approvalID,
			Error:      fmt.Sprintf("failed to update payout approval_id: %v", err),
		}
	}

	return nil
}

// Create approval stage for the given approval
func (s *approvalWorkflowServiceImpl) createApprovalStage(ctx context.Context, approvalObj *modelApproval.Approval, currentStage *modelApprovalWorkflowStage.ApprovalWorkflowStage, userID int, action objectApproval.ApprovalAction, note string, payoutID int) *modelApproval.ApprovalResult {
	var approvalResult objectApprovalStage.ApprovalResult
	if action == objectApproval.ApprovalActionRejected {
		approvalResult = objectApprovalStage.ApprovalResultRejected
	} else {
		approvalResult = objectApprovalStage.ApprovalResultApproved
	}

	approvalStage := modelApprovalStage.NewApprovalStage(modelApprovalStage.ApprovalStageParams{
		ApprovalID:              approvalObj.ID,
		ApprovalWorkflowStageID: currentStage.ID,
		ApproverID:              userID,
		Note:                    note,
		ApprovalResult:          approvalResult,
	})

	if err := s.approvalStageRepo.Create(ctx, approvalStage); err != nil {
		return &modelApproval.ApprovalResult{
			PayoutID:   payoutID,
			ApprovalID: approvalObj.ID,
			Error:      fmt.Sprintf("failed to create approval stage: %v", err),
		}
	}

	return nil
}

// Send email notification to approvers
func (s *approvalWorkflowServiceImpl) SendMailToApprover(ctx context.Context, payoutID int, workflowID int, stageID int) error {
	users, err := s.userRepo.GetUsersWithApprovalStageWorkflow(ctx, workflowID, stageID)
	if err != nil {
		s.logger.Error("Failed to get users for approval notification", map[string]any{"error": err})
		return fmt.Errorf("failed to get approvers: %w", err)
	}

	if len(users) == 0 {
		s.logger.Warn("No approvers found for workflow stage", map[string]any{"workflowID": workflowID, "stageID": stageID})
		return nil
	}

	for _, user := range users {
		err := s.sendApprovalRequestEmail(user.Email, user.FullName, payoutID)
		if err != nil {
			s.logger.Error("Failed to send approval request email", map[string]any{"error": err, "userID": user.ID})
		}
	}

	return nil
}

// sendApprovalRequestEmail sends an approval request email to a specific approver
func (s *approvalWorkflowServiceImpl) sendApprovalRequestEmail(to string, name string, payoutID int) error {
	return s.mailService.SendEmail(email.EmailData{
		To:             []string{to},
		Subject:        email.SubjectApprovalRequest,
		TemplateFile:   email.TemplateFileApprovalRequest,
		TemplateFolder: email.TemplateFolderApproval,
		Data: map[string]any{
			"ToName":   name,
			"PayoutID": payoutID,
			"Domain":   s.config.FrontUrl,
		},
		ContentType: email.ContentTypeHTML,
	})
}

// createTransactions creates new transactions and transaction records when a payout is rejected
func (s *approvalWorkflowServiceImpl) createTransactionsFromPayout(ctx context.Context, payoutID int) error {
	originalTransactionDetails, err := s.transactionRepo.GetTransactionDetailsByPayoutID(ctx, payoutID)
	if err != nil {
		return fmt.Errorf("failed to get transaction details by payout ID %d: %w", payoutID, err)
	}

	if len(originalTransactionDetails) == 0 {
		return fmt.Errorf("no transactions found for payout ID %d", payoutID)
	}

	transactionModels := []*transactionModel.Transaction{}
	for _, originalDetail := range originalTransactionDetails {
		newTransaction := &transactionModel.Transaction{
			ShopID:            originalDetail.Merchant.ShopID,
			BankCode:          bankAccountObject.FromStringToBankCode(originalDetail.BankCode),
			BankBranch:        bankAccountObject.BankBranch(originalDetail.BranchName),
			BankBranchCode:    bankAccountObject.FromStringToBankBranchCode(originalDetail.BankBranchCode),
			AccountNumber:     bankAccountObject.FromStringToAccountNumber(originalDetail.AccountNumber),
			AccountHolder:     bankAccountObject.AccountHolder(originalDetail.AccountName),
			AccountHolderKana: bankAccountObject.FromStringToAccountHolderKana(originalDetail.AccountName),
			AccountKind:       bankAccountObject.OrdinaryAccount,
		}
		newTransaction.SetDraft()
		transactionModels = append(transactionModels, newTransaction)
	}

	// Bulk create new transactions
	createdTransactions, err := s.transactionRepo.BulkCreate(ctx, transactionModels)
	if err != nil {
		return fmt.Errorf("failed to bulk create transactions: %w", err)
	}

	// Create new transaction records for each transaction
	transactionRecordModels := []*transactionModel.TransactionRecord{}
	for i, newTransaction := range createdTransactions {
		originalDetail := originalTransactionDetails[i]
		for _, originalRecord := range originalDetail.TransactionRecords {
			switch originalRecord.TransactionRecordType {
			case transactionRecordTypeObject.TransactionRecordTypeDeposit:
				depositTransactionRecord := transactionModel.NewTransactionRecord(
					newTransaction.ID,
					originalRecord.MerchantID,
					originalRecord.PayinDetailID,
					originalRecord.Amount,
					transactionRecordTypeObject.TransactionRecordTypeDeposit,
				)
				transactionRecordModels = append(transactionRecordModels, depositTransactionRecord)
			case transactionRecordTypeObject.TransactionRecordTypeFee:
				feeTransactionRecord := transactionModel.NewTransactionRecord(
					newTransaction.ID,
					originalRecord.MerchantID,
					originalRecord.PayinDetailID,
					originalRecord.Amount,
					transactionRecordTypeObject.TransactionRecordTypeFee,
				)
				transactionRecordModels = append(transactionRecordModels, feeTransactionRecord)
			case transactionRecordTypeObject.TransactionRecordTypeTransferFee:
				transferFeeTransactionRecord := transactionModel.NewTransactionRecord(
					newTransaction.ID,
					originalRecord.MerchantID,
					originalRecord.PayinDetailID,
					originalRecord.Amount,
					transactionRecordTypeObject.TransactionRecordTypeTransferFee,
				)
				transactionRecordModels = append(transactionRecordModels, transferFeeTransactionRecord)
			default:
				return fmt.Errorf("unknown transaction record type: %d", originalRecord.TransactionRecordType)
			}
		}
	}

	// Bulk create new transaction records
	if err := s.transactionRecordRepo.BulkCreate(ctx, transactionRecordModels); err != nil {
		return fmt.Errorf("failed to bulk create transaction records: %w", err)
	}

	return nil
}
