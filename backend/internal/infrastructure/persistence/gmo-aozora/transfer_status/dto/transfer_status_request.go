package dto

import (
	"net/url"

	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora/transfer_status"
)

type TransferStatusRequestDTO struct {
	AccountID             string
	QueryKeyClass         string
	ApplyNo               string
	DateFrom              string
	DateTo                string
	NextItemKey           string
	RequestTransferStatus []string
	RequestTransferClass  string
	RequestTransferTerm   string
}

func (dto *TransferStatusRequestDTO) ToQueryParams() url.Values {
	params := url.Values{}

	params.Set("accountId", dto.AccountID)
	params.Set("queryKeyClass", dto.QueryKeyClass)

	if dto.ApplyNo != "" {
		params.Set("applyNo", dto.ApplyNo)
	}

	if dto.DateFrom != "" {
		params.Set("dateFrom", dto.DateFrom)
	}

	if dto.DateTo != "" {
		params.Set("dateTo", dto.DateTo)
	}

	if dto.NextItemKey != "" {
		params.Set("nextItemKey", dto.NextItemKey)
	}

	if len(dto.RequestTransferStatus) > 0 {
		for _, status := range dto.RequestTransferStatus {
			params.Add("requestTransferStatus", status)
		}
	}

	if dto.RequestTransferClass != "" {
		params.Set("requestTransferClass", dto.RequestTransferClass)
	}

	if dto.RequestTransferTerm != "" {
		params.Set("requestTransferTerm", dto.RequestTransferTerm)
	}

	return params
}

func ToTransferStatusRequestDTO(m model.TransferStatusRequest) TransferStatusRequestDTO {
	dto := TransferStatusRequestDTO{
		AccountID:     m.AccountID,
		QueryKeyClass: m.QueryKeyClass.Value(),
		ApplyNo:       m.ApplyNo,
		DateFrom:      m.DateFrom,
		DateTo:        m.DateTo,
		NextItemKey:   m.NextItemKey,
	}

	if len(m.RequestTransferStatus) > 0 {
		dto.RequestTransferStatus = make([]string, len(m.RequestTransferStatus))
		for i, transferStatus := range m.RequestTransferStatus {
			dto.RequestTransferStatus[i] = transferStatus.Value()
		}
	}

	if m.RequestTransferClass != "" {
		dto.RequestTransferClass = m.RequestTransferClass.Value()
	}

	if m.RequestTransferTerm != "" {
		dto.RequestTransferTerm = m.RequestTransferTerm.Value()
	}

	return dto
}
