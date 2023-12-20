package mocks

import "github.com/aws/aws-sdk-go/service/ssm"

type CurrentAccountSyncServiceMock struct {
	SSMServiceMock
	GetParameterOutput ssm.GetParameterOutput
}
