package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqsAdapter "github.com/test-tzs/nomraeite/internal/domain/adapter/sqs"
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/sqs"
	"github.com/test-tzs/nomraeite/internal/infrastructure/adapter/util"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
)

type SQSAdapterConfig struct {
	QueueURL string
	util.AWSConfig
}

type SQSAdapterImpl struct {
	config SQSAdapterConfig
	client *sqs.Client
	logger logger.Logger
}

const (
	maxBatchSize = 10
)

func NewSQSAdapter(cfg SQSAdapterConfig) (sqsAdapter.SQSAdapter, error) {
	awsCfg, err := util.MakeAWSConfig(cfg.AWSConfig)
	if err != nil {
		return nil, err
	}
	client := sqs.NewFromConfig(awsCfg)

	return &SQSAdapterImpl{
		config: cfg,
		client: client,
		logger: logger.GetLogger(),
	}, nil
}

func (c *SQSAdapterImpl) SendBankTransferMessage(message *model.BankTransferStatusMessage) error {
	sMInput, err := utils.ToSQSMessageInput(message)
	if err != nil {
		return err
	}

	return c.sendMessage(sMInput)
}

func (c *SQSAdapterImpl) SendApplicationReviewResultMessage(message *model.ReviewResultMessage) error {
	sMInput, err := utils.ToSQSMessageInput(message)
	if err != nil {
		return err
	}

	return c.sendMessage(sMInput)
}

func (c *SQSAdapterImpl) SendApplicationReviewResultBatchMessage(messages []*model.ReviewResultMessage) error {
	sMInput, err := utils.ToSQSMessageBatchInput(messages)
	if err != nil {
		return err
	}

	return c.sendMessageBatch(sMInput)
}

func (c *SQSAdapterImpl) sendMessage(message *sqs.SendMessageInput) error {
	message.QueueUrl = &c.config.QueueURL
	_, err := c.client.SendMessage(context.Background(), message)
	if err != nil {
		c.logger.LogErrorWithContext(err, "Failed to send message to SQS", map[string]any{
			"queue_url":    &c.config.QueueURL,
			"message_body": *message.MessageBody,
		})
		return err
	}

	return nil
}

func (c *SQSAdapterImpl) splitBatch(messages *sqs.SendMessageBatchInput) []*sqs.SendMessageBatchInput {
	chunks := utils.Chunk(messages.Entries, maxBatchSize)

	batches := make([]*sqs.SendMessageBatchInput, 0, len(chunks))
	for _, chunk := range chunks {
		batch := &sqs.SendMessageBatchInput{
			QueueUrl: &c.config.QueueURL,
			Entries:  chunk,
		}
		batches = append(batches, batch)
	}

	return batches
}

func (c *SQSAdapterImpl) sendMessageBatch(messages *sqs.SendMessageBatchInput) error {
	messages.QueueUrl = &c.config.QueueURL
	batches := c.splitBatch(messages)
	for _, batch := range batches {
		_, err := c.client.SendMessageBatch(context.Background(), batch)
		if err != nil {
			c.logger.LogErrorWithContext(err, "Failed to send message batch to SQS", map[string]any{
				"queue_url": &c.config.QueueURL,
				"messages":  messages,
			})
		}
	}

	return nil
}
