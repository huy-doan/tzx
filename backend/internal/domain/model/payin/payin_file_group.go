package model

import (
	"time"

	paymentProviderModel "github.com/test-tzs/nomraeite/internal/domain/model/payment_provider"
	basedatetime "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	commonObject "github.com/test-tzs/nomraeite/internal/domain/object/common"
)

type PayinFileGroup struct {
	ID                int
	FileGroupName     string
	PaymentProviderID int
	ImportTargetDate  *time.Time
	ImportedAt        *time.Time
	PaymentProvider   *paymentProviderModel.PaymentProvider
	PayinFile         *[]PayinFile
	basedatetime.BaseColumnTimestamp
}

type PaginatedPayinFileGroupResult struct {
	Items []*PayinFileGroup
	commonObject.Pagination
}
