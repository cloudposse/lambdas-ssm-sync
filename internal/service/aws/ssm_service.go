package awssvc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type SSMService struct {
	client  ssmiface.SSMAPI
	session *session.Session
}

func NewSSMService(session *session.Session, cfgs ...*aws.Config) *SSMService {
	return &SSMService{
		client:  ssm.New(session, cfgs...),
		session: session,
	}
}

func (s *SSMService) SetRegion(regionName string) error {
	s.client = ssm.New(s.session, aws.NewConfig().WithRegion(regionName))
	return nil
}

func (s *SSMService) DeleteParameter(name string) (*ssm.DeleteParameterOutput, error) {
	return s.client.DeleteParameter(&ssm.DeleteParameterInput{
		Name: &name,
	})
}

func (s *SSMService) GetParameter(name string) (*ssm.GetParameterOutput, error) {
	return s.client.GetParameter(&ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: aws.Bool(true),
	})
}

func (s *SSMService) GetParametersByPathPages(input *ssm.GetParametersByPathInput) ([]*ssm.Parameter, error) {
	var params []*ssm.Parameter

	err := s.client.GetParametersByPathPages(input, func(page *ssm.GetParametersByPathOutput, lastPage bool) bool {
		params = append(params, page.Parameters...)
		return !lastPage
	})

	if err != nil {
		return nil, err
	}

	return params, nil
}

func (s *SSMService) SetParameter(name string, paramType string, value string) (*ssm.PutParameterOutput, error) {
	return s.client.PutParameter(&ssm.PutParameterInput{
		Name:      &name,
		Value:     &value,
		Overwrite: aws.Bool(true),
		Type:      aws.String(paramType),
	})
}
