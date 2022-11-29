package cfncli

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/pkg/errors"

	"cfn-drift-police/src/application/consts"
	cfndto "cfn-drift-police/src/application/dto/cloudformation"
)

type DefaultCloudFormationClient struct {
	svc cloudformation.CloudFormation
}

func NewDefaultCloudFormationClient() CloudFormationClient {
	sess := session.Must(session.NewSession())
	svc := cloudformation.New(sess, aws.NewConfig().WithRegion(os.Getenv(consts.AwsRegion)))

	return DefaultCloudFormationClient{
		svc: *svc,
	}
}

func (cli DefaultCloudFormationClient) ListStacks(in cfndto.ListStacksInput) (*cfndto.ListStacksOutput, error) {
	req := cloudformation.ListStacksInput{
		StackStatusFilter: in.StackStatusFilter,
		NextToken:         in.NextToken,
	}

	res, err := cli.svc.ListStacks(&req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	stackSummaries := []cfndto.StackSummary{}
	for _, ss := range res.StackSummaries {
		stackSummaries = append(stackSummaries, cfndto.StackSummary{
			StackName: *ss.StackName,
		})
	}

	out := cfndto.ListStacksOutput{
		NextToken:      res.NextToken,
		StackSummaries: stackSummaries,
	}

	return &out, nil
}

func (cli DefaultCloudFormationClient) DetectStackDrift(in cfndto.DetectStackDriftInput) (*cfndto.DetectStackDriftOutput, error) {
	req := cloudformation.DetectStackDriftInput{
		StackName: &in.StackName,
	}

	res, err := cli.svc.DetectStackDrift(&req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	out := cfndto.DetectStackDriftOutput{
		StackDriftDetectionId: *res.StackDriftDetectionId,
	}

	return &out, nil
}

func (cli DefaultCloudFormationClient) DescribeStackDriftDetectionStatus(in cfndto.DescribeStackDriftDetectionStatusInput) (*cfndto.DescribeStackDriftDetectionStatusOutput, error) {
	req := cloudformation.DescribeStackDriftDetectionStatusInput{
		StackDriftDetectionId: &in.StackDriftDetectionId,
	}

	res, err := cli.svc.DescribeStackDriftDetectionStatus(&req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	out := cfndto.DescribeStackDriftDetectionStatusOutput{
		DetectionStatus:       *res.DetectionStatus,
		DetectionStatusReason: res.DetectionStatusReason,
		StackDriftStatus:      res.StackDriftStatus,
		StackId:               *res.StackId,
	}
	return &out, nil
}
