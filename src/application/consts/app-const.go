package consts

const (
	// AWSリージョン
	AwsRegion string = "REGION"
	// appの実行ステージ
	Stage string = "STAGE"
	// 開発環境
	StageDev string = "dev"
	// AWSアカウントID
	AwsAccountId string = "ACCOUNT_ID"
	// SQSキューURL
	QueueUrl string = "QUEUE_URL"
	// SQSから取得するメッセージの件数
	SqsMaxNumberOfReceiveMessage int64 = 10
	// 取得したSQSメッセージが0件であること
	NoMessage int = 0
	// デコードをするか：する
	RequireDecryption bool = true
	// 赤色
	ColorRed string = "#D00000"
	// 改行コード
	LineFeed string = "\n"

	// CloudFormationスタックのステータス：CREATE_COMPLETE
	CfnStackStatusCreateComplete string = "CREATE_COMPLETE"
	// CloudFormationスタックのステータス：CREATE_IN_PROGRESS
	CfnStackStatusCreateInProgress string = "CREATE_IN_PROGRESS"
	// CloudFormationスタックのステータス：CREATE_FAILED
	CfnStackStatusCreateFailed string = "CREATE_FAILED"
	// CloudFormationスタックのステータス：DELETE_FAILED
	CfnStackStatusDeleteFailed string = "DELETE_FAILED"
	// CloudFormationスタックのステータス：DELETE_IN_PROGRESS
	CfnStackStatusDeleteInProgress string = "DELETE_IN_PROGRESS"
	// CloudFormationスタックのステータス：REVIEW_IN_PROGRESS
	CfnStackStatusReviewInProgress string = "REVIEW_IN_PROGRESS"
	// CloudFormationスタックのステータス：ROLLBACK_COMPLETE
	CfnStackStatusRollbackComplete string = "ROLLBACK_COMPLETE"
	// CloudFormationスタックのステータス：ROLLBACK_FAILED
	CfnStackStatusRollbackFailed string = "ROLLBACK_FAILED"
	// CloudFormationスタックのステータス：ROLLBACK_IN_PROGRESS
	CfnStackStatusRollbackInProgress string = "ROLLBACK_IN_PROGRESS"
	// CloudFormationスタックのステータス：UPDATE_COMPLETE
	CfnStackStatusUpdateComplete string = "UPDATE_COMPLETE"
	// CloudFormationスタックのステータス：UPDATE_COMPLETE_CLEANUP_IN_PROGRESS
	CfnStackStatusUpdateCompleteCleanupInProgress string = "UPDATE_COMPLETE_CLEANUP_IN_PROGRESS"
	// CloudFormationスタックのステータス：UPDATE_FAILED
	CfnStackStatusUpdateFailed string = "UPDATE_FAILED"
	// CloudFormationスタックのステータス：UPDATE_IN_PROGRESS
	CfnStackStatusUpdateInProgress string = "UPDATE_IN_PROGRESS"
	// CloudFormationスタックのステータス：UPDATE_ROLLBACK_COMPLETE
	CfnStackStatusUpdateRollbackComplete string = "UPDATE_ROLLBACK_COMPLETE"
	// CloudFormationスタックのステータス：UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS
	CfnStackStatusUpdateRollbackComplateCleanupInProgress string = "UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS"
	// CloudFormationスタックのステータス：UPDATE_ROLLBACK_FAILED
	CfnStackStatusUpdateRollbackFailed string = "UPDATE_ROLLBACK_FAILED"
	// CloudFormationスタックのステータス：UPDATE_ROLLBACK_IN_PROGRESS
	CfnStackStatusUpdateRollbackInProgress string = "UPDATE_ROLLBACK_IN_PROGRESS"
	// CloudFormationスタックのステータス：IMPORT_IN_PROGRESS
	CfnStackStatusImportInProgress string = "IMPORT_IN_PROGRESS"
	// CloudFormationスタックのステータス：IMPORT_COMPLETE
	CfnStackStatusImportComplete string = "IMPORT_COMPLETE"
	// CloudFormationスタックのステータス：IMPORT_ROLLBACK_IN_PROGRESS
	CfnStackStatusImportRollbackInProgress string = "IMPORT_ROLLBACK_IN_PROGRESS"
	// CloudFormationスタックのステータス：IMPORT_ROLLBACK_FAILED
	CfnStackStatusImportRollbackFailed string = "IMPORT_ROLLBACK_FAILED"
	// CloudFormationスタックのステータス：IMPORT_ROLLBACK_COMPLETE
	CfnStackStatusImportRollbackComplete string = "IMPORT_ROLLBACK_COMPLETE"
	// CloudFormationドリフト結果ステータス：DETECTION_FAILED
	CfnDetectionStatusFailed string = "DETECTION_FAILED"
	// CloudFormationドリフト状態ステータス：DRIFTED
	CfnDriftStatusDrifted string = "DRIFTED"
)
