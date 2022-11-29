package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	conf "cfn-drift-police/src/application/config"
	"cfn-drift-police/src/application/consts"
	ssmdto "cfn-drift-police/src/application/dto/ssm"
	log "cfn-drift-police/src/application/logger"
	comutil "cfn-drift-police/src/util/commons"
)

// ロガーを生成
var logger = log.NewAppLoger()

/**
mainは、文字通りのmain関数です。lambdaのハンドラを起動し、対象となる処理を実行します。
*/
func main() {
	lambda.Start(Handler)
}

/**
Handlerは、alertLambdaの処理本体です。AlertUsecaseを生成し、ユースケースを実行します。
*/
func Handler() {
	ssmCli := conf.InitSsmClientByStage()

	req := ssmdto.GetParameterInput{
		Name:              os.Getenv(consts.SlackTokenStore),
		RequireDecryption: comutil.BoolCtoP(consts.RequireDecryption),
	}
	res, err := ssmCli.GetParameter(req)
	if err != nil {
		logger.Fatal(consts.LOG0018, err)
	}

	au := conf.InitAlertUsecaseByStage(res.Parameter.Value)
	au.Execute()
}
