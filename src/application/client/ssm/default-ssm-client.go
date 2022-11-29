package ssmcli

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/pkg/errors"

	"cfn-drift-police/src/application/consts"
	ssmdto "cfn-drift-police/src/application/dto/ssm"
)

type DefaultSsmClient struct {
	svc ssm.SSM
}

func NewDefaultSsmClient() SsmClient {
	sess := session.Must(session.NewSession())
	svc := ssm.New(sess, aws.NewConfig().WithRegion(os.Getenv(consts.AwsRegion)))

	return DefaultSsmClient{
		svc: *svc,
	}
}

func (cli DefaultSsmClient) GetParameter(in ssmdto.GetParameterInput) (*ssmdto.GetParameterOutput, error) {
	req := ssm.GetParameterInput{
		Name:           &in.Name,
		WithDecryption: in.RequireDecryption,
	}

	res, err := cli.svc.GetParameter(&req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	out := ssmdto.GetParameterOutput{
		Parameter: ssmdto.Parameter{
			Value: *res.Parameter.Value,
		},
	}

	return &out, nil
}
