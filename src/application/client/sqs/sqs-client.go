package sqscli

import sqsdto "cfn-drift-police/src/application/dto/sqs"

/**
SqsClientは、Amazon Simple Queue Serviceクライアントのインタフェースとなる構造体です。
*/
type SqsClient interface {
	/**
	SendMessageは、SQSにメッセージを送信する関数です。
	*/
	SendMessage(in sqsdto.SendMessageInput) error
	/**
	ReceiveMessageは、SQSからメッセージを受信する関数です。
	*/
	ReceiveMessage(in sqsdto.ReceiveMessageInput) (*sqsdto.ReceiveMessageOutput, error)
	/**
	DeleteMessageは、SQS上の対象メッセージを取得する関数です。
	*/
	DeleteMessage(in sqsdto.DeleteMessageInput) error
}
