package object

import "strconv"

// 振込ステータス
type GmoAozoraTransferStatus string

const (
	GmoAozoraTransferStatusApplication      GmoAozoraTransferStatus = "2"  // 申請中
	GmoAozoraTransferStatusReverted         GmoAozoraTransferStatus = "3"  // 差戻
	GmoAozoraTransferStatusWithdrawn        GmoAozoraTransferStatus = "4"  // 取下げ
	GmoAozoraTransferStatusExpired          GmoAozoraTransferStatus = "5"  // 期限切れ
	GmoAozoraTransferStatusCancelled        GmoAozoraTransferStatus = "8"  // 承認取消/予約取消
	GmoAozoraTransferStatusReserved         GmoAozoraTransferStatus = "11" // 予約中
	GmoAozoraTransferStatusInProcess        GmoAozoraTransferStatus = "12" // 手続中
	GmoAozoraTransferStatusRetrying         GmoAozoraTransferStatus = "13" // リトライ中
	GmoAozoraTransferStatusCompleted        GmoAozoraTransferStatus = "20" // 手続済
	GmoAozoraTransferStatusFailed           GmoAozoraTransferStatus = "40" // 手続不成立

	//Only for transfer status
	GmoAozoraTransferStatusRefund           GmoAozoraTransferStatus = "22" // 資金返却
	GmoAozoraTransferStatusRefundProcessing GmoAozoraTransferStatus = "24" // 組戻手続中
	GmoAozoraTransferStatusRefunded         GmoAozoraTransferStatus = "25" // 組戻済
	GmoAozoraTransferStatusRefundFailed     GmoAozoraTransferStatus = "26" // 組戻不成立

	//Only for bulk transfer status
	GmoAozoraTransferStatusCannotProcess    GmoAozoraTransferStatus = "30" // 不能・組戻あり
)

func (g GmoAozoraTransferStatus) Value() string {
	return string(g)
}

func (g GmoAozoraTransferStatus) ToInt() int {
	status, err := strconv.Atoi(g.Value())
	if err != nil {
		return 0
	}
	return status
}
