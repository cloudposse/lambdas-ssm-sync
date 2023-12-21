package awssvciface

import "github.com/aws/aws-sdk-go/service/ssm"

type SSMService interface {
	SetRegion(name string) error
	DeleteParameter(name string) (*ssm.DeleteParameterOutput, error)
	GetParameter(name string) (*ssm.GetParameterOutput, error)
	SetParameter(name string, paramType string, value string) (*ssm.PutParameterOutput, error)
}
