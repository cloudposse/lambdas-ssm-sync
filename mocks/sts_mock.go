package mocks

import (
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

type STSClientMock struct {
	stsiface.STSAPI
	GetCallerIdentityOutout sts.GetCallerIdentityOutput
}

func (m STSClientMock) GetCallerIdentity(*sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	return &m.GetCallerIdentityOutout, nil
}
