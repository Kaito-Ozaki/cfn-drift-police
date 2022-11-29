package awsutil

/**
SqsMessageBodyは、このprojectでSQSにメッセージを送る際の、メッセージボディ内容を規定する構造体です。
*/
type SqsMessageBody struct {
	// CloudFormationのスタックに、ドリフト検出をかけた際に発行されるID
	StackDriftDetectionId string
	// CloudFormationのスタック名
	StackName string
}
