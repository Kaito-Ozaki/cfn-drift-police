package sqsdto

/**
SendMessageInputは、SendMessageのinputを規定するDTOです。
*/
type SendMessageInput struct {
	// 送信するメッセージの内容
	MessageBody string
	// 対象となるキューのURL
	QueueUrl string
}

/**
ReceiveMessageInputは、ReceiveMessageのinputを規定するDTOです。
*/
type ReceiveMessageInput struct {
	// 対象となるキューのURL
	QueueUrl string
	// 受信するメッセージの最大数
	MaxNumberOfMessages *int64
}

/**
DeleteMessageInputは、DeleteMessageeのinputを規定するDTOです。
*/
type DeleteMessageInput struct {
	// 対象となるキューのURL
	QueueUrl string
	// 対象となるメッセージのreceiptHandle
	ReceiptHandle string
}
