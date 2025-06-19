package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/sqs"
)

type BankTransferStatusMessage struct {
	MessageAttributes object.MessageAttribute
	MessageBody       object.BodyTransferStatus
}

func (m *BankTransferStatusMessage) GetMessageAttributes() object.MessageAttribute {
	return m.MessageAttributes
}

func (m *BankTransferStatusMessage) GetMessageBody() any {
	return m.MessageBody
}
