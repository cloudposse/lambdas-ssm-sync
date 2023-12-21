package awsutil

import (
	"fmt"
	"strings"
)

type SSMParameterARNParts struct {
	SourceRegion  string
	SourceAccount string
	PathAccount   string
	PathRegion    string
}

func GetSSMParameterARNParts(input string) (SSMParameterARNParts, error) {
	parts := strings.Split(input, ":")
	if len(parts) < 6 || parts[0] != "arn" || parts[1] != "aws" || parts[2] != "ssm" {
		return SSMParameterARNParts{}, fmt.Errorf("the specified value %s is not an ssm parameter ARN", input)
	}

	sourceRegion := parts[3]
	sourceAccount := parts[4]
	//arn:aws:ssm:us-east-2:123456789012:parameter/MyExampleParameter
	pathParts := strings.Split(parts[5], "/")
	if len(pathParts) < 4 {
		return SSMParameterARNParts{}, fmt.Errorf("the specified value %s does not match the /terraform/shared-state/{account}/{region} pattern", input)
	}

	pathAccount := pathParts[3]
	pathRegion := pathParts[4]

	return SSMParameterARNParts{
		SourceRegion:  sourceRegion,
		SourceAccount: sourceAccount,
		PathAccount:   pathAccount,
		PathRegion:    pathRegion,
	}, nil
}
