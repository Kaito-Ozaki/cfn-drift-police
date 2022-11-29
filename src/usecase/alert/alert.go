package alert

import (
	"encoding/json"
	"fmt"
	"os"

	cfncli "cfn-drift-police/src/application/client/cloudformation"
	slackcli "cfn-drift-police/src/application/client/slack"
	sqscli "cfn-drift-police/src/application/client/sqs"
	"cfn-drift-police/src/application/consts"
	cfndto "cfn-drift-police/src/application/dto/cloudformation"
	sqsdto "cfn-drift-police/src/application/dto/sqs"
	log "cfn-drift-police/src/application/logger"
	awsutil "cfn-drift-police/src/util/aws"
	comutil "cfn-drift-police/src/util/commons"

	"github.com/pkg/errors"
	"github.com/slack-go/slack"
)

// ロガーを生成
var logger = log.NewAppLoger()

/**
AlertUsecaseは、CloudFormationスタックへのドリフト検出結果を取得し、ドリフトが発生した場合は、通知を行うというユースケースを所持する構造体です。
*/
type AlertUsecase struct {
	// CloudFormationクライアント
	cfnCli cfncli.CloudFormationClient
	// SQSクライアント
	sqsCli sqscli.SqsClient
	// Slackクライアント
	slackCli slackcli.SlackClient
}

/**
NewAlertUsecaseは、AlertUsecaseを生成するための、初期化関数です。
*/
func NewAlertUsecase(cfnCli cfncli.CloudFormationClient, sqsCli sqscli.SqsClient, slackCli slackcli.SlackClient) AlertUsecase {
	return AlertUsecase{
		cfnCli:   cfnCli,
		sqsCli:   sqsCli,
		slackCli: slackCli,
	}
}

/**
Executeは、CloudFormationスタックへのドリフト検出結果を取得し、ドリフトが発生した場合は、通知を行う関数です。
*/
func (au *AlertUsecase) Execute() {
	logger.Info(consts.LOG0009)
	messages, err := au.getAllSqsMessages(os.Getenv(consts.QueueUrl))
	if err != nil {
		logger.Fatal(consts.LOG0011, err)
	}
	logger.Info(consts.LOG0010)

	logger.Info(consts.LOG0012)
	for _, m := range *messages {
		sqsMessageData := awsutil.SqsMessageBody{}
		if err := json.Unmarshal([]byte(m.Body), &sqsMessageData); err != nil {
			err = errors.WithStack(err)
			logger.Error(consts.LOG0014, err)

			err = au.deleteSqsMessage(m.ReceiptHandle, os.Getenv(consts.QueueUrl))
			if err != nil {
				logger.ErrorWithParams(consts.LOG0017, err, consts.Unknown)
			}
			continue
		}

		stackDetectionResult, err := au.getStackDriftDetectionResult(sqsMessageData.StackDriftDetectionId)
		if err != nil {
			logger.Error(consts.LOG0015, err)

			err = au.deleteSqsMessage(m.ReceiptHandle, os.Getenv(consts.QueueUrl))
			if err != nil {
				logger.ErrorWithParams(consts.LOG0017, err, sqsMessageData.StackName)
			}
			continue
		}

		if stackDetectionResult.DetectionStatus == consts.CfnDetectionStatusFailed {
			dsr := cfndto.DetectionStatusReason{}
			if err := json.Unmarshal([]byte(*stackDetectionResult.DetectionStatusReason), &dsr); err == nil {
				rs := []string{}
				for _, f := range dsr.Failures {
					rs = append(rs, f.Resource)
				}

				err := au.postDetectFailureMessage(sqsMessageData.StackName, stackDetectionResult.StackId, rs)
				if err != nil {
					logger.ErrorWithParams(consts.LOG0016, err, sqsMessageData.StackName)
				}
			} else {
				logger.ErrorWithParams(consts.LOG0019, err, sqsMessageData.StackName)
			}
		}

		if stackDetectionResult.StackDriftStatus != nil && *stackDetectionResult.StackDriftStatus == consts.CfnDriftStatusDrifted {
			err := au.postDriftMessage(sqsMessageData.StackName, stackDetectionResult.StackId)
			if err != nil {
				logger.ErrorWithParams(consts.LOG0016, err, sqsMessageData.StackName)
			}
		}

		err = au.deleteSqsMessage(m.ReceiptHandle, os.Getenv(consts.QueueUrl))
		if err != nil {
			logger.ErrorWithParams(consts.LOG0017, err, sqsMessageData.StackName)
		}
	}
	logger.Info(consts.LOG0013)
}

/**
getAllSqsMessagesは、対象となるSQSから全てのメッセージを取得する関数です。
*/
func (au *AlertUsecase) getAllSqsMessages(queueUrl string) (*[]sqsdto.Message, error) {
	messages := []sqsdto.Message{}
	res, err := au.getSqsMessages(queueUrl)
	if err != nil {
		return nil, err
	}
	for res != nil {
		messages = append(messages, *res...)
		res, err = au.getSqsMessages(queueUrl)
		if err != nil {
			return nil, err
		}
	}
	return &messages, nil
}

/**
getSqsMessagesは、対象となるSQSから10件のメッセージを取得する関数です。
*/
func (au *AlertUsecase) getSqsMessages(queueUrl string) (*[]sqsdto.Message, error) {
	req := sqsdto.ReceiveMessageInput{
		QueueUrl:            queueUrl,
		MaxNumberOfMessages: comutil.Int64CtoP(consts.SqsMaxNumberOfReceiveMessage),
	}

	res, err := au.sqsCli.ReceiveMessage(req)
	if err != nil {
		return nil, err
	}
	return res.Messages, nil
}

/**
getStackDetectionStatusは、ドリフト検出が実行されたCloudFormationスタックから、ドリフトステータスを取得する関数です。
*/
func (au *AlertUsecase) getStackDriftDetectionResult(stackDriftDetectionId string) (*cfndto.DescribeStackDriftDetectionStatusOutput, error) {
	req := cfndto.DescribeStackDriftDetectionStatusInput{
		StackDriftDetectionId: stackDriftDetectionId,
	}

	res, err := au.cfnCli.DescribeStackDriftDetectionStatus(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/**
postDetectFailureMessageは、slackにドリフト検出に失敗したリソースに関するメッセージを投稿する関数です。
*/
func (au *AlertUsecase) postDetectFailureMessage(stackName string, stackId string, resources []string) error {
	rsString := resources[0]
	for i, _ := range resources {
		if i == 0 {
			continue
		}
		rsString = rsString + consts.LineFeed + consts.DetectFailureLineFeedBlank + resources[i]
	}

	at := slack.Attachment{
		Color: consts.ColorRed,
		Text:  fmt.Sprintf(consts.DetectFailureSlackMessageTemplate, stackName, rsString, os.Getenv(consts.AwsRegion), stackId),
	}

	err := au.slackCli.PostMessage(
		os.Getenv(consts.SlackChannelName),
		slack.MsgOptionUsername(consts.SlackAppUserName),
		slack.MsgOptionIconEmoji(consts.AlertIcon),
		slack.MsgOptionText(consts.DetectFailureSlackMessageTitle, true),
		slack.MsgOptionAttachments(at),
	)
	if err != nil {
		return err
	}
	return nil
}

/**
postDriftMessageは、slackにドリフトしているスタックに関するメッセージを投稿する関数です。
*/
func (au *AlertUsecase) postDriftMessage(stackName string, stackId string) error {
	at := slack.Attachment{
		Color: consts.ColorRed,
		Text:  fmt.Sprintf(consts.DriftSlackMessageTemplate, stackName, os.Getenv(consts.AwsRegion), stackId),
	}

	err := au.slackCli.PostMessage(
		os.Getenv(consts.SlackChannelName),
		slack.MsgOptionUsername(consts.SlackAppUserName),
		slack.MsgOptionIconEmoji(consts.AlertIcon),
		slack.MsgOptionText(consts.DriftSlackMessageTitle, true),
		slack.MsgOptionAttachments(at),
	)
	if err != nil {
		return err
	}
	return nil
}

/**
deleteSqsMessageは、対象となるメッセージをSQS上から削除する関数です。
*/
func (au *AlertUsecase) deleteSqsMessage(receiptHandle string, queueUrl string) error {
	req := sqsdto.DeleteMessageInput{
		QueueUrl:      queueUrl,
		ReceiptHandle: receiptHandle,
	}

	err := au.sqsCli.DeleteMessage(req)
	if err != nil {
		return err
	}
	return nil
}
