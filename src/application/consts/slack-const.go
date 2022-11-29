package consts

const (
	// slackユーザーネーム：インフラアラート
	SlackAppUserName string = "kake-infra-alert"
	// slackアイコン：アラート
	AlertIcon string = ":rotating_light:"
	// チャンネル名
	SlackChannelName string = "SLACK_CHANNEL_NAME"
	// トークンの格納場所
	SlackTokenStore string = "SLACK_TOKEN_STORE"
	// ドリフト検出失敗メッセージタイトル
	DetectFailureSlackMessageTitle string = "ドリフト検出に失敗しました！ 異常なスタックかもしれないので確認してください！"
	// ドリフト検出失敗メッセージテンプレート {0}スタック名：{1}検出失敗したリソースリスト：{2}リージョン：{3}スタックID
	DetectFailureSlackMessageTemplate string = "スタック名　　　： %s\n失敗したリソース： %s\n詳細  　　　   　　：<https://ap-northeast-1.console.aws.amazon.com/cloudformation/home?region=%s#/stacks/drifts?stackId=%s | 詳細はこちら>"
	// ドリフト検出失敗メッセージ改行時の空白
	DetectFailureLineFeedBlank string = "　　　　　　　　　"
	// ドリフト通知メッセージタイトル
	DriftSlackMessageTitle string = "ドリフトが検出されました！ スタックを確認してください！"
	// ドリフト通知メッセージテンプレート {0}スタック名：{1}リージョン：{2}スタックID
	DriftSlackMessageTemplate string = "スタック名： %s\n詳細     　　：<https://ap-northeast-1.console.aws.amazon.com/cloudformation/home?region=%s#/stacks/drifts?stackId=%s | 詳細はこちら>"
)
