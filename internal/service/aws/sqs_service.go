package awssvc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type SQSService struct {
	client   sqsiface.SQSAPI
	session  *session.Session
	QueueUrl string
}

func NewSQSService(queueUrl string, session *session.Session, cfgs ...*aws.Config) *SQSService {
	return &SQSService{
		client:   sqs.New(session, cfgs...),
		session:  session,
		QueueUrl: queueUrl,
	}
}

func (s *SQSService) SendMessage(message string, messageGroupId string) error {
	if _, err := s.client.SendMessage(&sqs.SendMessageInput{
		QueueUrl:       &s.QueueUrl,
		MessageBody:    &message,
		MessageGroupId: &messageGroupId,
	}); err != nil {
		return err
	}

	return nil
}

func (s *SQSService) SetQueueUrl(url string) error {
	s.QueueUrl = url
	return nil
}
