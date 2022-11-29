package sqscli

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/pkg/errors"

	"cfn-drift-police/src/application/consts"
	sqsdto "cfn-drift-police/src/application/dto/sqs"
)

type DefaultSqsClient struct {
	svc sqs.SQS
}

func NewDefaultSqsClient() SqsClient {
	sess := session.Must(session.NewSession())
	svc := sqs.New(sess, aws.NewConfig().WithRegion(os.Getenv(consts.AwsRegion)))

	return DefaultSqsClient{
		svc: *svc,
	}
}

func (cli DefaultSqsClient) SendMessage(in sqsdto.SendMessageInput) error {
	req := sqs.SendMessageInput{
		QueueUrl:    &in.QueueUrl,
		MessageBody: &in.MessageBody,
	}

	_, err := cli.svc.SendMessage(&req)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (cli DefaultSqsClient) ReceiveMessage(in sqsdto.ReceiveMessageInput) (*sqsdto.ReceiveMessageOutput, error) {
	req := sqs.ReceiveMessageInput{
		QueueUrl:            &in.QueueUrl,
		MaxNumberOfMessages: in.MaxNumberOfMessages,
	}

	res, err := cli.svc.ReceiveMessage(&req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if res.Messages == nil {
		return &sqsdto.ReceiveMessageOutput{Messages: nil}, nil
	} else {
		messages := []sqsdto.Message{}
		for _, m := range res.Messages {
			messages = append(messages, sqsdto.Message{
				Body:          *m.Body,
				ReceiptHandle: *m.ReceiptHandle,
			})
		}
		out := sqsdto.ReceiveMessageOutput{
			Messages: &messages,
		}

		return &out, nil
	}
}

func (cli DefaultSqsClient) DeleteMessage(in sqsdto.DeleteMessageInput) error {
	req := sqs.DeleteMessageInput{
		QueueUrl:      &in.QueueUrl,
		ReceiptHandle: &in.ReceiptHandle,
	}

	_, err := cli.svc.DeleteMessage(&req)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
