package sqsdto

/**
ReceiveMessageOutputは、ReceiveMessageのoutputを定義するDTOです。
*/
type ReceiveMessageOutput struct {
	// SQSメッセージ
	Messages *[]Message
}

/**
Messageは、SQS Messageが持つ情報をまとめた構造体です。
*/
type Message struct {
	// メッセージの内容
	Body string
	// receiptHandle
	ReceiptHandle string
}
