package check

import (
	"encoding/json"
	"os"

	cfncli "cfn-drift-police/src/application/client/cloudformation"
	sqscli "cfn-drift-police/src/application/client/sqs"
	"cfn-drift-police/src/application/consts"
	cfndto "cfn-drift-police/src/application/dto/cloudformation"
	sqsdto "cfn-drift-police/src/application/dto/sqs"
	log "cfn-drift-police/src/application/logger"
	awsutil "cfn-drift-police/src/util/aws"
	comutil "cfn-drift-police/src/util/commons"
)

// ロガーを生成
var logger = log.NewAppLoger()

/**
CheckUsecaseは、CloudFormationのスタックに対して、ドリフトが発生していないかの検出を実行し、実行した際のIDを後処理に伝えるというユースケースを所持する構造体です。
*/
type CheckUsecase struct {
	// CloudFormationクライアント
	cfnCli cfncli.CloudFormationClient
	// SQSクライアント
	sqsCli sqscli.SqsClient
}

/**
NewCheckUsecaseは、CheckUsecaseを生成するための、初期化関数です。
*/
func NewCheckUsecase(cfnCli cfncli.CloudFormationClient, sqsCli sqscli.SqsClient) CheckUsecase {
	return CheckUsecase{
		cfnCli: cfnCli,
		sqsCli: sqsCli,
	}
}

/**
Executeは、CloudFormationのスタックに対して、ドリフトが発生していないかの検出を実行し、実行した際のIDおよび対象となったスタック名を後工程に送信する関数です。
*/
func (cu *CheckUsecase) Execute() {
	targetStackStatusList := []string{
		consts.CfnStackStatusCreateComplete,
		consts.CfnStackStatusCreateInProgress,
		consts.CfnStackStatusCreateFailed,
		consts.CfnStackStatusDeleteFailed,
		consts.CfnStackStatusDeleteInProgress,
		consts.CfnStackStatusReviewInProgress,
		consts.CfnStackStatusRollbackComplete,
		consts.CfnStackStatusRollbackFailed,
		consts.CfnStackStatusRollbackInProgress,
		consts.CfnStackStatusUpdateComplete,
		consts.CfnStackStatusUpdateCompleteCleanupInProgress,
		consts.CfnStackStatusUpdateFailed,
		consts.CfnStackStatusUpdateInProgress,
		consts.CfnStackStatusUpdateRollbackComplete,
		consts.CfnStackStatusUpdateRollbackComplateCleanupInProgress,
		consts.CfnStackStatusUpdateRollbackFailed,
		consts.CfnStackStatusUpdateRollbackInProgress,
		consts.CfnStackStatusImportInProgress,
		consts.CfnStackStatusImportComplete,
		consts.CfnStackStatusImportRollbackInProgress,
		consts.CfnStackStatusImportRollbackFailed,
		consts.CfnStackStatusImportRollbackComplete,
	}
	logger.Info(consts.LOG0001)
	stacks, err := cu.getAllStackList(targetStackStatusList)
	if err != nil {
		logger.Fatal(consts.LOG0003, err)
	}
	stackNames := cu.extractStackName(*stacks)
	targetStackNames := comutil.DeleteByList(stackNames, consts.BlackList)
	logger.Info(consts.LOG0002)

	logger.Info(consts.LOG0004)
	for _, v := range targetStackNames {
		stackDriftDetectionId, err := cu.detectStackDrift(v)
		if err != nil {
			logger.ErrorWithParams(consts.LOG0006, err, v)
			continue
		}

		sqsMessageData := awsutil.SqsMessageBody{
			StackDriftDetectionId: *stackDriftDetectionId,
			StackName:             v,
		}
		jsonData, _ := json.Marshal(sqsMessageData)

		err = cu.sendSqsMessage(string(jsonData), os.Getenv(consts.QueueUrl))
		if err != nil {
			logger.ErrorWithParams(consts.LOG0008, err, v)
		}
	}
	logger.Info(consts.LOG0005)
}

/**
getAllStackListは、指定したステータスのCloudFormationスタックの一覧を取得する関数です。
*/
func (cu *CheckUsecase) getAllStackList(targetStackStatusList []string) (*[]cfndto.StackSummary, error) {
	stackLists := []cfndto.StackSummary{}
	res, nextToken, err := cu.getStackList(targetStackStatusList, nil)
	if err != nil {
		return nil, err
	}
	stackLists = append(stackLists, *res...)
	for nextToken != nil {
		res, nextToken, err = cu.getStackList(targetStackStatusList, nextToken)
		if err != nil {
			return nil, err
		}
		stackLists = append(stackLists, *res...)
	}
	return &stackLists, nil
}

/**
getStackListは、指定したステータスのCloudFormationを10件取得する関数です。
*/
func (cu *CheckUsecase) getStackList(targetStackStatusList []string, nextToken *string) (*[]cfndto.StackSummary, *string, error) {
	targetStackStatusListToP := comutil.StringSliceToStringPSlice(targetStackStatusList)
	req := cfndto.ListStacksInput{
		StackStatusFilter: targetStackStatusListToP,
	}
	res, err := cu.cfnCli.ListStacks(req)
	if err != nil {
		return nil, nil, err
	}
	return &res.StackSummaries, res.NextToken, nil
}

/**
extractStackNameは、CloudFormationスタックのサマリ情報から、スタック名だけを抽出し、スタック名のListを返す関数です。
*/
func (cu *CheckUsecase) extractStackName(in []cfndto.StackSummary) []string {
	stackNames := []string{}
	for _, v := range in {
		stackNames = append(stackNames, v.StackName)
	}
	return stackNames
}

/**
detectStackDriftは、CloudFormationのスタックに対して、ドリフト検出を実行し、実行時に生成されたdetectionIDを返す関数です。
*/
func (cu *CheckUsecase) detectStackDrift(stackName string) (*string, error) {
	req := cfndto.DetectStackDriftInput{
		StackName: stackName,
	}

	res, err := cu.cfnCli.DetectStackDrift(req)

	if err != nil {
		return nil, err
	}
	return &res.StackDriftDetectionId, nil
}

/**
sendSqsMessageは、SQSにメッセージを送信する関数です。
*/
func (cu *CheckUsecase) sendSqsMessage(messageBody string, queueUrl string) error {
	req := sqsdto.SendMessageInput{
		MessageBody: messageBody,
		QueueUrl:    queueUrl,
	}

	err := cu.sqsCli.SendMessage(req)
	if err != nil {
		return err
	}
	return nil
}
