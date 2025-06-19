package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/sqs"
)

type ReviewResultMessage struct {
	MessageAttributes object.MessageAttribute
	MessageBody       object.BodyApplicationReviewResult
}

func (m *ReviewResultMessage) GetMessageAttributes() object.MessageAttribute {
	return m.MessageAttributes
}

func (m *ReviewResultMessage) GetMessageBody() any {
	return m.MessageBody
}
