package object

import "strconv"

// 振込ステータス
type GmoAozoraBulkTransferStatus string

const (
	GmoAozoraBulkTransferStatusApplication   GmoAozoraBulkTransferStatus = "2"  // 申請中
	GmoAozoraBulkTransferStatusReverted      GmoAozoraBulkTransferStatus = "3"  // 差戻
	GmoAozoraBulkTransferStatusWithdrawn     GmoAozoraBulkTransferStatus = "4"  // 取下げ
	GmoAozoraBulkTransferStatusExpired       GmoAozoraBulkTransferStatus = "5"  // 期限切れ
	GmoAozoraBulkTransferStatusCancelled     GmoAozoraBulkTransferStatus = "8"  // 承認取消/予約取消
	GmoAozoraBulkTransferStatusReserved      GmoAozoraBulkTransferStatus = "11" // 予約中
	GmoAozoraBulkTransferStatusInProcess     GmoAozoraBulkTransferStatus = "12" // 手続中
	GmoAozoraBulkTransferStatusRetrying      GmoAozoraBulkTransferStatus = "13" // リトライ中
	GmoAozoraBulkTransferStatusCompleted     GmoAozoraBulkTransferStatus = "20" // 手続済
	GmoAozoraBulkTransferStatusCannotProcess GmoAozoraBulkTransferStatus = "30" // 不能・組戻あり
	GmoAozoraBulkTransferStatusFailed        GmoAozoraBulkTransferStatus = "40" // 手続不成立
)

func (g GmoAozoraBulkTransferStatus) Value() string {
	return string(g)
}

func (g GmoAozoraBulkTransferStatus) ToInt() int {
	status, err := strconv.Atoi(g.Value())
	if err != nil {
		return 0
	}
	return status
}
