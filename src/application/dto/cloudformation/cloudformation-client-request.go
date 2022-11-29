package cfndto

/**
ListStacksInputは、ListStacksのinputを定義するDTOです。
*/
type ListStacksInput struct {
	// 取得対象となるCloudFormationスタックのステータス一覧
	StackStatusFilter []*string
	// 前回取得時に返却された、次回取得用のtoken
	NextToken *string
}

/**
DetectStackDriftInputは、DetectStackDriftのinputを定義するDTOです。
*/
type DetectStackDriftInput struct {
	// CloudFormationスタック名
	StackName string
}

/**
DescribeStackDriftDetectionStatusInputは、DescribeStackDriftDetectionStatusのinputを定義するDTOです。
*/
type DescribeStackDriftDetectionStatusInput struct {
	// ドリフト検出時に発行されたID
	StackDriftDetectionId string
}
