package ssmcli

import (
	"github.com/aws/aws-sdk-go/service/ssm"

	ssmdto "cfn-drift-police/src/application/dto/ssm"
)

type InMemorySsmClient struct {
	svc ssm.SSM
}

func NewInMemorySsmClient() SsmClient {
	return InMemorySsmClient{}
}

func (cli InMemorySsmClient) GetParameter(in ssmdto.GetParameterInput) (*ssmdto.GetParameterOutput, error) {
	out := ssmdto.GetParameterOutput{
		Parameter: ssmdto.Parameter{
			Value: "test-token",
		},
	}
	return &out, nil
}
