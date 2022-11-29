package ssmdto

/**
GetParameterOutputは、GetParameterのoutputを定義するDTOです。
*/
type GetParameterOutput struct {
	// パラメーター
	Parameter Parameter
}

/**
Parameterは、SSM Parameter Storeに格納されている変数の情報をまとめた構造体です。
*/
type Parameter struct {
	// 変数の値
	Value string
}
