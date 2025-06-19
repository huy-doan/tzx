package object

import (
	sqsObject "github.com/test-tzs/nomraeite/internal/domain/object/sqs"
)

type ApplicationReviewStatus int

const (
	ApplicationReviewStatusProcessing            ApplicationReviewStatus = 1 // 審査中
	ApplicationReviewStatusApproved              ApplicationReviewStatus = 2 // 審査通過
	ApplicationReviewStatusRejected              ApplicationReviewStatus = 3 // 審査否認
	ApplicationReviewStatusReturnedForCorrection ApplicationReviewStatus = 4 // 審査不備理由記載
)

var filePrefixMapping = map[string]ApplicationReviewStatus{
	"審査中":      ApplicationReviewStatusProcessing,
	"審査不備理由記載": ApplicationReviewStatusReturnedForCorrection,
	"審査可決":     ApplicationReviewStatusApproved,
	"審査否決":     ApplicationReviewStatusRejected,
}

func (m ApplicationReviewStatus) String() string {
	switch m {
	case ApplicationReviewStatusProcessing:
		return "審査中"
	case ApplicationReviewStatusReturnedForCorrection:
		return "審査不備理由記載"
	case ApplicationReviewStatusApproved:
		return "審査通過"
	case ApplicationReviewStatusRejected:
		return "審査否認"
	default:
		return "不明"
	}
}

func (m ApplicationReviewStatus) Value() int {
	switch m {
	case ApplicationReviewStatusProcessing:
		return 1
	case ApplicationReviewStatusApproved:
		return 2
	case ApplicationReviewStatusRejected:
		return 3
	case ApplicationReviewStatusReturnedForCorrection:
		return 4
	default:
		return 0
	}
}

func (m ApplicationReviewStatus) SQSValue() sqsObject.PaymentApplicationReviewResult {
	switch m {
	case ApplicationReviewStatusProcessing:
		return sqsObject.PaymentApplicationReviewResultInProgress
	case ApplicationReviewStatusReturnedForCorrection:
		return sqsObject.PaymentApplicationReviewResultRequestChange
	case ApplicationReviewStatusApproved:
		return sqsObject.PaymentApplicationReviewResultApproved
	case ApplicationReviewStatusRejected:
		return sqsObject.PaymentApplicationReviewResultRejected
	default:
		return 0
	}
}

func (m ApplicationReviewStatus) IsPending() bool {
	return m == ApplicationReviewStatusProcessing
}

func (m ApplicationReviewStatus) IsReturnedForCorrection() bool {
	return m == ApplicationReviewStatusReturnedForCorrection
}

func (m ApplicationReviewStatus) IsApproved() bool {
	return m == ApplicationReviewStatusApproved
}

func (m ApplicationReviewStatus) IsRejected() bool {
	return m == ApplicationReviewStatusRejected
}

func DetectReviewStatusFromFileName(fileName string) (ApplicationReviewStatus, bool) {
	for prefix, status := range filePrefixMapping {
		if len(fileName) >= len(prefix) && fileName[:len(prefix)] == prefix {
			return status, true
		}
	}

	return 0, false
}

func (m ApplicationReviewStatus) GetRequiredHeaders() []string {
	baseHeaders := []string{
		"大手フラグ",
		"ID区分",
		"契約先申請番号",
		"申込み完了日",
		"加盟店接続企業",
		"法人名、個人事業主名",
		"屋号名",
		"申請店舗数",
		"PayOrigin18桁Id",
		"包括代理店_備考",
		"決済手数料率",
	}
	switch m {
	case ApplicationReviewStatusProcessing:
		return append(baseHeaders, "オンライン決済利用サイトURL", "営業確認ステータス変更日時", "[審査]営業向けコメント", "申込担当者_回答")
	case ApplicationReviewStatusReturnedForCorrection:
		return append(baseHeaders, "オンライン決済利用サイトURL", "営業確認ステータス変更日時", "[審査]営業向けコメント", "申込担当者_回答")
	case ApplicationReviewStatusApproved:
		return append(baseHeaders, "オンライン決済利用サイトURL", "審査可決日", "MID", "月間決済上限金額", "ID発行状況")
	case ApplicationReviewStatusRejected:
		return append(baseHeaders, "審査ステータス", "審査否決日")
	default:
		return nil
	}
}
