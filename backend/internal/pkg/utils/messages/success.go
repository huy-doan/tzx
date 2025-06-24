package messages

const (
	MsgListAuditLogsSuccess    = "監査ログの一覧を取得しました"
	MsgGetAuditLogUsersSuccess = "ユーザー一覧を取得しました"

	// Role Success Messages
	MsgListRolesSuccess                  = "ロール一覧を取得しました"
	MsgGetRoleSuccess                    = "ロールを取得しました"
	MsgCreateRoleSuccess                 = "ロールが作成されました"
	MsgUpdateRoleSuccess                 = "ロールが更新されました"
	MsgDeleteRoleSuccess                 = "ロールが削除されました"
	MsgBatchUpdateRolePermissionsSuccess = "ロール権限の一括更新が完了しました"

	// Permission Success Messages
	MsgListPermissionsSuccess = "権限一覧を取得しました"

	// Screen Success Messages
	MsgListScreensSuccess = "画面一覧を取得しました"

	// Auth related success messages
	MsgLoginSuccess        = "ログインしました"
	MsgLogoutSuccess       = "ログアウトしました"
	MsgMFARequired         = "2FA認証が必要です"
	MsgAuthCodeSentSuccess = "認証コードが正常に送信されました"

	// merchant related success messages
	MsgListMerchantsSuccess = "加盟店一覧を取得しました"

	// notification related success messages
	MsgListNotificationsSuccess = "通知リストを取得しました"
	// dashboard transaction success messages
	TransactionSummarySuccess = "ダッシュボードの取引件数を取得できました。"
	// payout related success messages
	MsgListPayoutsSuccess = "出金履歴の取得に成功しました"

	// payin related success messages
	MsgListPayinsSuccess = "支払い履歴が正常に取得されました"

	// transaction related success messages
	MsgListTransferRequestsSuccess = "振込依頼の一覧を取得しました"
	// payin related success messages
	MsgListPayinSummarySuccess = "支払い履歴が正常に取得されました"

	// payin detail related success messages
	MsgListPayinFilesSuccess = "支払い詳細履歴が正常に取得されました"

	// Virtual Account related success messages
	MsgListVirtualAccountsSuccess  = "仮想アカウントは正常に取得されました"
	MsgCreateVirtualAccountSuccess = "バーチャルアカウントを登録しました"
	MsgUpdateVirtualAccountSuccess = "バーチャルアカウントを更新しました"

	// User related success messages
	MsgListUsersSuccess  = "ユーザー一覧を取得しました"
	MsgCreateUserSuccess = "ユーザーを登録しました"
	MsgUpdateUserSuccess = "ユーザーを更新しました"
	MsgDeleteUserSuccess = "ユーザーを削除しました"
	MsgGetUserSuccess    = "ユーザーを取得しました"

	MsgListShopReviewsSuccess = "ショップ一覧を取得しました"

	// Payment Provider Review related success messages
	MsgPaymentProviderReviewCSVUploaded = "決済プロバイダ審査CSVがアップロードされました"
	MsgPaymentProviderReviewCSVImported = "決済プロバイダ審査CSVがインポートされました"

	// Payment Provider related success messages
	MsgListPaymentProvidersSuccess = "決済プロバイダ一覧を取得しました"

	// Aozora Connect Success Messages
	MsgAozoraConnectionStatusSuccess = "GMOあおぞらネット銀行の接続状態を取得しました"
	MsgAozoraConnectionRevokeSuccess = "GMOあおぞらネット銀行の接続を解除しました"
	MsgAozoraCallbackSuccess         = "GMOあおぞらネット銀行との接続が完了しました"

	// Approval Success Messages
	MsgTransferApprovalSuccess = "承認の処理に成功しました"
	// Transaction related success messages
	MsgGetTransactionDetailSuccess = "取引詳細を取得しました"
	MsgListTransfersSuccess        = "取引履歴を取得しました"

	// Payout related success messages
	MsgCreatePayoutSuccess = "出金が正常に作成されました"

	// Payout related success messages
	MsgPayoutNotFound = "出金が見つかりません"
)
