package ssmcli

import ssmdto "cfn-drift-police/src/application/dto/ssm"

/**
SsmClientは、AWS Systems Managerクライアントのインタフェースとなる構造体です。
*/
type SsmClient interface {
	/**
	GetParameterは、SSMパラメータストアから変数を取得する関数です。
	*/
	GetParameter(in ssmdto.GetParameterInput) (*ssmdto.GetParameterOutput, error)
}
