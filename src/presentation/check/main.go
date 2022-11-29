package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	conf "cfn-drift-police/src/application/config"
)

/**
mainは、文字通りのmain関数です。lambdaのハンドラを起動し、対象となる処理を実行します。
*/
func main() {
	lambda.Start(Handler)
}

/**
Handlerは、checkLambdaの処理本体です。CheckUsecaseを生成し、ユースケースを実行します。
*/
func Handler() {
	cu := conf.InitCheckUsecaseByStage()
	cu.Execute()
}
