package mocks

type SQSServiceMock struct {
	SendMessageCalls int
	SetQueueUrlCalls int
}

func (s *SQSServiceMock) SendMessage(message string, messageGroupId string) error {
	s.SendMessageCalls = s.SendMessageCalls + 1
	return nil
}

func (s *SQSServiceMock) SetQueueUrl(url string) error {
	s.SetQueueUrlCalls = s.SetQueueUrlCalls + 1
	return nil
}
