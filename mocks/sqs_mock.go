package mocks

type SQSServiceMock struct {
	SendMessageCalls int
}

func (m *SQSServiceMock) SendMessage(message string, messageGroupId string) error {
	return nil
}
