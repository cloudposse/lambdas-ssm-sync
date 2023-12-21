package awssvciface

import "github.com/aws/aws-sdk-go/aws"

type AWSSession struct {
	Config *aws.Config
}
