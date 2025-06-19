package object

type TransactionRecordType int

// TransactionRecordType constants defined from the database comments
const (
	TransactionRecordTypeDeposit     TransactionRecordType = iota + 1 // 入金
	TransactionRecordTypeFee                                          // 手数料
	TransactionRecordTypeTransferFee                                  // 振込手数料
)

func (t TransactionRecordType) String() string {
	switch t {
	case TransactionRecordTypeDeposit:
		return "入金"
	case TransactionRecordTypeFee:
		return "Paypay手数料"
	case TransactionRecordTypeTransferFee:
		return "振込手数料"
	default:
		return "不明"
	}
}
