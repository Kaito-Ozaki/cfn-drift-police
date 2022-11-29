package conf

import (
	"os"

	cfncli "cfn-drift-police/src/application/client/cloudformation"
	slackcli "cfn-drift-police/src/application/client/slack"
	sqscli "cfn-drift-police/src/application/client/sqs"
	ssmcli "cfn-drift-police/src/application/client/ssm"
	"cfn-drift-police/src/application/consts"
	"cfn-drift-police/src/usecase/alert"
	"cfn-drift-police/src/usecase/check"
)

/**
initProdAlertUsecaseは、prod環境版のAlertUsecaseを生成する関数です。
各クライアントを生成し、AlertUsecaseに注入します。
*/
func initProdAlertUsecase(token string) alert.AlertUsecase {
	cfnCli := cfncli.NewDefaultCloudFormationClient()
	sqsCli := sqscli.NewDefaultSqsClient()
	slackCli := slackcli.NewDefaultSlackClient(token)
	return alert.NewAlertUsecase(cfnCli, sqsCli, slackCli)
}

/**
initDevAlertUsecaseは、dev環境版のAlertUsecaseを生成する関数です。
各クライアントのモックを生成し、AlertUsecaseに注入します。
*/
func initDevAlertUsecase(token string) alert.AlertUsecase {
	cfnCli := cfncli.NewInMemoryCloudFormationClient()
	sqsCli := sqscli.NewInMemorySqsClient()
	slackCli := slackcli.NewInMemorySlackClient(token)
	return alert.NewAlertUsecase(cfnCli, sqsCli, slackCli)
}

/**
InitAlertUsecaseByStageは、appの実行ステージに応じて、対応するAlertUsecaseを生成する関数です。
*/
func InitAlertUsecaseByStage(token string) alert.AlertUsecase {
	stage := os.Getenv(consts.Stage)
	if stage != consts.StageDev {
		return initProdAlertUsecase(token)
	} else {
		return initDevAlertUsecase(token)
	}
}

/**
initProdCheckUsecaseは、prod環境版のCheckUsecaseを生成する関数です。
各クライアントを生成し、CheckUsecaseに注入します。
*/
func initProdCheckUsecase() check.CheckUsecase {
	cfnCli := cfncli.NewDefaultCloudFormationClient()
	sqsCli := sqscli.NewDefaultSqsClient()
	return check.NewCheckUsecase(cfnCli, sqsCli)
}

/**
initDevCheckUsecaseは、dev環境版のCheckUsecaseを生成する関数です。
各クライアントのモックを生成し、CheckUsecaseに注入します。
*/
func initDevCheckUsecase() check.CheckUsecase {
	cfnCli := cfncli.NewInMemoryCloudFormationClient()
	sqsCli := sqscli.NewInMemorySqsClient()
	return check.NewCheckUsecase(cfnCli, sqsCli)
}

/**
InitCheckUsecaseByStageは、appの実行ステージに応じて、対応するCheckUsecaseを生成する関数です。
*/
func InitCheckUsecaseByStage() check.CheckUsecase {
	stage := os.Getenv(consts.Stage)
	if stage != consts.StageDev {
		return initProdCheckUsecase()
	} else {
		return initDevCheckUsecase()
	}
}

/**
InitSsmClientByStageは、appの実行ステージに応じて、対応するSsmClientを生成する関数です。
*/
func InitSsmClientByStage() ssmcli.SsmClient {
	stage := os.Getenv(consts.Stage)
	if stage != consts.StageDev {
		return ssmcli.NewDefaultSsmClient()
	} else {
		return ssmcli.NewInMemorySsmClient()
	}
}
