package cfndto

/**
ListStacksOutputは、ListStacksのoutputを定義するDTOです。
*/
type ListStacksOutput struct {
	// 前回取得時に返却された、次回取得用のtoken
	NextToken *string
	// CloudFormationスタックの情報
	StackSummaries []StackSummary
}

/**
StackSummaryは、CloudFormationスタックの情報をまとめた構造体です。
*/
type StackSummary struct {
	// CloudFormationスタック名
	StackName string
}

/**
DetectStackDriftOutputは、DetectStackDriftのoutputを定義するDTOです。
*/
type DetectStackDriftOutput struct {
	// ドリフト検出時に発行されたID
	StackDriftDetectionId string
}

/**
DescribeStackDriftDetectionStatusOutputは、DescribeStackDriftDetectionStatusのoutputを定義するDTOです。
*/
type DescribeStackDriftDetectionStatusOutput struct {
	// ドリフト検出ステータス
	DetectionStatus string
	// ドリフト検出失敗理由
	DetectionStatusReason *string
	// ドリフト状態ステータス
	StackDriftStatus *string
	// スタックID
	StackId string
}

/**
DetectionStatusReasonは、ドリフト検出を失敗した理由をまとめた構造体です。
*/
type DetectionStatusReason struct {
	// 失敗理由概要
	Summary string `json:"Summary"`
	// 失敗
	Failures []DetectionStatusFailure `json:"Failures"`
}

/**
DetectionStatusFailureは、ドリフト検出失敗の内容をまとめた構造体です。
*/
type DetectionStatusFailure struct {
	// 失敗したリソース
	Resource string `json:"Resource"`
	// 失敗した理由
	FailureReason string `json:"FailureReason"`
}
