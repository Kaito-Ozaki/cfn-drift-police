package cfncli

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"

	"cfn-drift-police/src/application/consts"
	cfndto "cfn-drift-police/src/application/dto/cloudformation"
	comutil "cfn-drift-police/src/util/commons"
)

type InMemoryCloudFormationClient struct {
	svc cloudformation.CloudFormation
}

func NewInMemoryCloudFormationClient() CloudFormationClient {
	return InMemoryCloudFormationClient{}
}

func (cli InMemoryCloudFormationClient) ListStacks(in cfndto.ListStacksInput) (*cfndto.ListStacksOutput, error) {
	ss := []cfndto.StackSummary{}
	ss = append(ss, cfndto.StackSummary{
		StackName: "inmemory-cfn-drift-police",
	})
	out := cfndto.ListStacksOutput{
		StackSummaries: ss,
	}
	return &out, nil
}

func (cli InMemoryCloudFormationClient) DetectStackDrift(in cfndto.DetectStackDriftInput) (*cfndto.DetectStackDriftOutput, error) {
	out := cfndto.DetectStackDriftOutput{
		StackDriftDetectionId: "arn:aws:cloudformation:ap-northeast-1:000000000000:stack/test-cfn-drift-police/aa67b910-dfe2-11ec-b933-068c0ea753ab",
	}
	return &out, nil
}

func (cli InMemoryCloudFormationClient) DescribeStackDriftDetectionStatus(in cfndto.DescribeStackDriftDetectionStatusInput) (*cfndto.DescribeStackDriftDetectionStatusOutput, error) {
	out := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       "DETECTION_COMPLETE",
		DetectionStatusReason: nil,
		StackDriftStatus:      comutil.StringCtoP(consts.CfnDriftStatusDrifted),
		StackId:               "arn%3Aaws%3Acloudformation%3Aap-northeast-1%3A000000000000%3Astack%2Ftest-cfn-drift-police%2F0dfbbf70-e7d7-11ec-8705-0633c5fc036d",
	}
	return &out, nil
}
