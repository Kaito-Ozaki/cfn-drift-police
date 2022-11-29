package sqscli

import (
	"github.com/aws/aws-sdk-go/service/sqs"

	sqsdto "cfn-drift-police/src/application/dto/sqs"
)

// mockで確実に結果を返すようにしてしまうと、SQSを全件取得するための処理（usecase/alert/alert.goのgetAllSqsMessagesなど）で無限ループが発生する
// そのため、無限ループが発生しないように、初回の呼び出しだけ結果を返し、次回の呼び出しでは結果を返さないよう、呼び出し回数を内部でカウントしておく
var receiveMessageCallCount int

type InMemorySqsClient struct {
	svc sqs.SQS
}

func NewInMemorySqsClient() SqsClient {
	receiveMessageCallCount = 0
	return InMemorySqsClient{}
}

func (cli InMemorySqsClient) SendMessage(in sqsdto.SendMessageInput) error {
	return nil
}

func (cli InMemorySqsClient) ReceiveMessage(in sqsdto.ReceiveMessageInput) (*sqsdto.ReceiveMessageOutput, error) {
	if receiveMessageCallCount == 0 {
		receiveMessageCallCount += 1
		ms := []sqsdto.Message{}
		ms = append(ms, sqsdto.Message{
			Body:          "{\"StackDriftDetectionId\":\"a0bd8b10-f77a-11ec-b078-0e3f99ca5407\",\"StackName\":\"cfn-drift-police-test\"}",
			ReceiptHandle: "AQEBpGB3SkHfpFiruWMPl2SGoWv3efebUE84CrvAcCJ4hnbSeV0NkkbECJsNVANFU6Whv+DSgww+7RpSVTb2DtU2i1dIuE/izK2bCXew==",
		})
		out := sqsdto.ReceiveMessageOutput{
			Messages: &ms,
		}
		return &out, nil
	} else {
		out := sqsdto.ReceiveMessageOutput{
			Messages: nil,
		}
		return &out, nil
	}
}

func (cli InMemorySqsClient) DeleteMessage(in sqsdto.DeleteMessageInput) error {
	return nil
}
