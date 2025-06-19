package model

import (
	"time"

	// model "github.com/test-tzs/nomraeite/internal/domain/model/paypay"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payin"
)

// PayinFile represents the payin_file domain entity
type PayinFile struct {
	ID int
	util.BaseColumnTimestamp

	PaymentProviderID int
	PayinFileGroupID  int
	FileName          string
	FileContentKey    string
	PayinFileType     object.PayinFileType
	HasDataRecord     bool
	AddedManually     bool
	CreatedAt         time.Time
	ImportStatus      object.PayinFileStatus
	UploadStatus      object.PayinFileStatus

	PayinFileGroup *PayinFileGroup
	// PayPayPayinSummary     *model.PayPayPayinSummary
	// PayPayPayinDetail      *[]model.PayPayPayinDetail
	// PayPayPayinTransaction *[]model.PayPayPayinTransaction
}

type PaginatedPayinFileResult struct {
	Items []*PayinFile
	util.Pagination
}
