package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/sqs"
)

type BaseSQSMessage interface {
	GetMessageAttributes() object.MessageAttribute
	GetMessageBody() any
}
