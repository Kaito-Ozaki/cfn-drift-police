package ssmdto

/**
GetParameterInputは、GetParameterのinputを定義するDTOです。
*/
type GetParameterInput struct {
	// パラメータ名
	Name string
	// デコードが必要か
	RequireDecryption *bool
}
