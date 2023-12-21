package mocks

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cloudposse/ssm-syncronizer/internal/service/aws/awssvciface"
)

func NewAWSSessionMock(region string) *awssvciface.AWSSession {
	return &awssvciface.AWSSession{
		Config: &aws.Config{
			Region: aws.String(region),
		},
	}
}
