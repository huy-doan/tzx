package object

import (
	"strings"

	"github.com/test-tzs/nomraeite/internal/pkg/config"
)

type PayinFileType int

const (
	PayinReport             PayinFileType = 1 // 入金取引明細集計レポート
	PayinTransactionSummary PayinFileType = 2 // 入金取引明細集計
	PayinTransactionDetail  PayinFileType = 3 // 入金取引明細
)

// DetectPayinFileType returns the PayinFileType based on the remote path and config
// Pass a struct with PaypayPayinTransactionPath, PayinSummaryPath, PaypayPayinReportPath fields
func DetectPayinFileType(targetFolder string, fileName string, cfg config.Config) PayinFileType {
	if strings.Contains(targetFolder, cfg.PaypayPayinReportPath) {
		return PayinReport
	}
	if strings.Contains(targetFolder, cfg.PaypayPayinTransactionPath) {
		parts := strings.SplitN(fileName, "_入金_", 2)
		if len(parts) < 2 {
			return 0
		}
		before := parts[0]

		if strings.Contains(before, "_") {
			// <屋号>_<加盟店ID>_入金_<入金日(YYYYMMDD)>_<入金処理日(YYYYMMDD)>_v2.zip
			return PayinTransactionDetail
		} else {
			// <グループ名>_入金_<入金日(YYYYMMDD)_<入金処理日(YYYYMMDD)>_v2.zip
			return PayinTransactionSummary
		}
	}
	return 0
}
