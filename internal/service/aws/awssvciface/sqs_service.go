package awssvciface

type SQSService interface {
	SendMessage(message string, messageGroupId string) error
	SetQueueUrl(queueUrl string) error
}
