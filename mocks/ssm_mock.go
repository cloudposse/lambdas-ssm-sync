package mocks

import (
	"github.com/aws/aws-sdk-go/service/ssm"
)

type SetParameterCall struct {
	Name      string
	ParamType string
	Value     string
}

type SSMServiceMock struct {
	PutParameterOutput             ssm.PutParameterOutput
	GetParameterOutput             ssm.GetParameterOutput
	GetParametersByPathPagesOutput []*ssm.Parameter
	SetParameterCalls              []SetParameterCall
	SetRegionCalls                 []string
}

func (m *SSMServiceMock) SetRegion(name string) error {
	m.SetRegionCalls = append(m.SetRegionCalls, name)
	return nil
}

func (m *SSMServiceMock) DeleteParameter(name string) (*ssm.DeleteParameterOutput, error) {
	return &ssm.DeleteParameterOutput{}, nil
}

func (m *SSMServiceMock) GetParameter(name string) (*ssm.GetParameterOutput, error) {
	return &ssm.GetParameterOutput{}, nil
}

func (s *SSMServiceMock) GetParametersByPathPages(input *ssm.GetParametersByPathInput) ([]*ssm.Parameter, error) {
	return s.GetParametersByPathPagesOutput, nil
}

func (m *SSMServiceMock) SetParameter(name string, paramType string, value string) (*ssm.PutParameterOutput, error) {
	m.SetParameterCalls = append(m.SetParameterCalls, SetParameterCall{Name: name, ParamType: paramType, Value: value})
	return &ssm.PutParameterOutput{}, nil
}
