package model

import (
	incomingModel "github.com/test-tzs/nomraeite/internal/domain/model/webhook/gmo-aozora"
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
)

type RawJsonData = utils.JSONValue[incomingModel.GMOAozoraWebhookBodyData]
