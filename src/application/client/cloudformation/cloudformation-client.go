package cfncli

import cfndto "cfn-drift-police/src/application/dto/cloudformation"

/**
CloufFormationClientは、AWS CloufFormationクライアントのインタフェースとなる構造体です。
*/
type CloudFormationClient interface {
	/**
	ListStacksは、CloudFormationスタックの情報を10件単位で取得する関数です。
	*/
	ListStacks(in cfndto.ListStacksInput) (*cfndto.ListStacksOutput, error)
	/**
	DetectStackDriftは、対象となるCloudFormationスタックに対して、ドリフト検出を実行する関数です。
	*/
	DetectStackDrift(in cfndto.DetectStackDriftInput) (*cfndto.DetectStackDriftOutput, error)
	/**
	DescribeStackDriftDetectionStatusは、対象となるドリフト検出の結果を取得する関数です。
	*/
	DescribeStackDriftDetectionStatus(in cfndto.DescribeStackDriftDetectionStatusInput) (*cfndto.DescribeStackDriftDetectionStatusOutput, error)
}
