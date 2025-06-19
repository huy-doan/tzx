package dto

import (
	"net/url"
	"strconv"

	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora/transfer_status"
)

type BulkTransferStatusRequestDTO struct {
	AccountID             string
	QueryKeyClass         string
	DetailInfoNecessity   *bool
	BulktransferItemKey   *string
	ApplyNo               *string
	DateFrom              *string
	DateTo                *string
	NextItemKey           *string
	RequestTransferStatus []string
	RequestTransferClass  *string
	RequestTransferTerm   *string
}

func (dto *BulkTransferStatusRequestDTO) ToQueryParams() url.Values {
	params := url.Values{}

	params.Set("accountId", dto.AccountID)
	params.Set("queryKeyClass", dto.QueryKeyClass)

	if dto.DetailInfoNecessity != nil {
		params.Set("detailInfoNecessity", strconv.FormatBool(*dto.DetailInfoNecessity))
	}

	if dto.BulktransferItemKey != nil {
		params.Set("bulktransferItemKey", *dto.BulktransferItemKey)
	}

	if dto.ApplyNo != nil {
		params.Set("applyNo", *dto.ApplyNo)
	}

	if dto.DateFrom != nil {
		params.Set("dateFrom", *dto.DateFrom)
	}

	if dto.DateTo != nil {
		params.Set("dateTo", *dto.DateTo)
	}

	if dto.NextItemKey != nil {
		params.Set("nextItemKey", *dto.NextItemKey)
	}

	if len(dto.RequestTransferStatus) > 0 {
		for _, status := range dto.RequestTransferStatus {
			params.Add("requestTransferStatus", status)
		}
	}

	if dto.RequestTransferClass != nil {
		params.Set("requestTransferClass", *dto.RequestTransferClass)
	}

	if dto.RequestTransferTerm != nil {
		params.Set("requestTransferTerm", *dto.RequestTransferTerm)
	}

	return params
}

func ToBulkTransferStatusRequestDTO(request model.BulkTransferStatusRequest) BulkTransferStatusRequestDTO {
	dto := BulkTransferStatusRequestDTO{
		AccountID:           request.AccountID,
		QueryKeyClass:       request.QueryKeyClass.Value(),
		DetailInfoNecessity: request.DetailInfoNecessity,
		BulktransferItemKey: request.BulktransferItemKey,
		ApplyNo:             request.ApplyNo,
		DateFrom:            request.DateFrom,
		DateTo:              request.DateTo,
		NextItemKey:         request.NextItemKey,
	}

	if len(request.RequestTransferStatus) > 0 {
		dto.RequestTransferStatus = make([]string, len(request.RequestTransferStatus))
		for i, status := range request.RequestTransferStatus {
			dto.RequestTransferStatus[i] = string(status)
		}
	}

	if request.RequestTransferClass != nil {
		classStr := string(*request.RequestTransferClass)
		dto.RequestTransferClass = &classStr
	}

	if request.RequestTransferTerm != nil {
		termStr := string(*request.RequestTransferTerm)
		dto.RequestTransferTerm = &termStr
	}

	return dto
}
