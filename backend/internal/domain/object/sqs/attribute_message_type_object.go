package object

type AttributeMessageType string

const (
	ApplyReviewResults AttributeMessageType = "ApplyReviewResults" // 決済申請結果
	BankTransferStatus AttributeMessageType = "BankTransferStatus" // 銀行振込ステータス
)
