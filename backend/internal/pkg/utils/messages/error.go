package messages

const (
	// Error Code
	CodeValidationError      = "VALIDATION_ERROR"
	CodeNotFound             = "NOT_FOUND"
	CodeUnauthorized         = "UNAUTHORIZED"
	CodeForbidden            = "FORBIDDEN"
	CodeBadRequest           = "BAD_REQUEST"
	CodeInternalError        = "INTERNAL_ERROR"
	CodeDatabaseError        = "DATABASE_ERROR"
	CodeExternalServiceError = "EXTERNAL_SERVICE_ERROR"
	CodeUnknownError         = "UNKNOWN_ERROR"
	CodeServiceUnavailable   = "SERVICE_UNAVAILABLE"
	CodeTimeout              = "TIMEOUT"
	CodeTooManyRequests      = "TOO_MANY_REQUESTS"
	CodeUnsupportedMediaType = "UNSUPPORTED_MEDIA_TYPE"
	CodeMethodNotAllowed     = "METHOD_NOT_ALLOWED"
	CodeNotAcceptable        = "NOT_ACCEPTABLE"

	// Error Types
	TypeValidationError      = "VALIDATION"
	TypeAuthorizationError   = "AUTHORIZATION"
	TypeDatabaseError        = "DATABASE"
	TypeExternalServiceError = "EXTERNAL"
	TypeServerError          = "SERVER"
	TypeClientError          = "CLIENT"
	TypeNotFoundError        = "NOT_FOUND"
	TypeUnknownError         = "UNKNOWN"
	TypeServiceUnavailable   = "SERVICE_UNAVAILABLE"
	TypeTimeoutError         = "TIMEOUT"

	// Error Messages
	MsgInvalidCredentials = "メールアドレスまたはパスワードが無効です"
	MsgEmailAlreadyExists = "このメールアドレスは既に登録されています"
	MsgUserNotFound       = "ユーザーが見つかりません"
	MsgInvalidPassword    = "現在のパスワードが正しくありません"
	MsgAccountLocked      = "アカウントがロックされています"
	MsgBadRequest         = "不正なリクエストです"

	// Route and API errors
	MsgRouteNotFound    = "リクエストされたURLが見つかりません"
	MsgMethodNotAllowed = "このリクエストメソッドは許可されていません"

	// Role Error Messages
	MsgListRolesError                  = "ロール一覧の取得に失敗しました"
	MsgGetRoleError                    = "ロールの取得に失敗しました"
	MsgRoleNotFoundError               = "ロールが見つかりません"
	MsgCreateRoleError                 = "ロールの作成に失敗しました"
	MsgUpdateRoleError                 = "ロールの更新に失敗しました"
	MsgDeleteRoleError                 = "ロールの削除に失敗しました"
	MsgBatchUpdateRolePermissionsError = "ロール権限の一括更新に失敗しました"
	MsgIDRequiredError                 = "IDが必要です"

	// Permission Error Messages
	MsgListPermissionsError = "権限一覧の取得に失敗しました"

	// Screen Error Messages
	MsgListScreensError = "画面一覧の取得に失敗しました"

	// Auth middleware error messages
	MsgUnauthorized       = "認証ヘッダーが必要です"
	MsgInvalidToken       = "トークンの検証エラー"
	MsgForbidden          = "このリソースにアクセスする権限がありません"
	MsgTokenBlacklisted   = "無効または失効したトークン"
	MsgInvalidTokenFormat = "認証ヘッダーの形式が正しくありません"
	MsgLogoutFailed       = "ログアウトに失敗しました"

	// Auth related error messages
	MsgLoginFailed      = "ログインに失敗しました"
	MsgUnauthenticated  = "認証されていません"
	MsgMFAVerifyFailed  = "認証コードの検証に失敗しました"
	MsgResendCodeFailed = "認証コードの再送信に失敗しました"
	// user related error messages
	MsgCreateUserFailed = "ユーザーを登録できませんでした"
	MsgUpdateUserFailed = "ユーザーを更新できませんでした"
	MsgDeleteUserFailed = "ユーザーを削除できませんでした"
	MsgGetUserFailed    = "ユーザーを取得できませんでした"

	// Single Transfer Approval related error messages
	MsgTransferApprovalFailed = "承認の処理に失敗しました"
	// Transfer related error messages
	MsgTransferRequestFailed           = "振込依頼が失敗しました。問題が解決しない場合システム管理者までお問い合わせ下さい"
	MsgTransferRequestFailedValidation = "振込依頼が失敗しました。振込の内容をもう一度確認してください"
	MsgTransferRequestAvoidDuplicate   = "多重振り込みを避けるため、振り込みをしていないことを確認するまでは再実行しないでください"

	// Payment Provider Review related error messages
	MsgInvalidPaymentProviderReviewCSVFile = "無効な決済プロバイダ審査CSVファイルです"
	MsgPaymentProviderIDRequired           = "決済プロバイダIDが必要です"
	MsgFailedToGetLastReviewByShopIDs      = "決済プロバイダ審査の最新レビューを取得できませんでした"

	// Payment Provider related error messages
	MsgListPaymentProvidersError = "決済プロバイダ一覧の取得に失敗しました"

	// CSV related error messages
	MsgFileNotFound             = "ファイルが見つかりません"
	MsgInvalidPaymentProviderID = "決済プロバイダIDが無効です: %s, エラー: %v"
	MsgFileNotCSV               = "ファイルがCSVではありません"
	MsgInvalidFileNameForReview = "審査ステータス検出用のファイル名が無効です"
	MsgCSVHeaderMismatch        = "CSVヘッダーが必要なヘッダーと一致しません。審査ステータス: %s。必要なヘッダー: %v。実際のヘッダー: %v"
	MsgFailedToCreateCSVDecoder = "CSVデコーダーの作成に失敗しました"
	MsgFailedToParseDate        = "日付の解析に失敗しました '%s': %v"
	MsgFailedToDecodeCSVRecord  = "CSVレコードのデコードに失敗しました"
	MsgCSVCallbackError         = "CSVコールバックエラー"
	MsgFailedToReadCSV          = "CSVの読み取りに失敗しました"
)
