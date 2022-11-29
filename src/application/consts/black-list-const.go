package consts

var (
	// ドリフト検出の実行対象外としたい、CloudFormationスタックのスタック名リスト
	BlackList = []string{
		"black-list-sample-stack",
	}
)
