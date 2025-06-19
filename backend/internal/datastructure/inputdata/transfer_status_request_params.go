package inputdata

import (
	"time"
)

type TransferStatusRequestParams struct {
	AccessToken         string
	AccountID           string
	NextItemKey         string
	SendingDate         time.Time
}
