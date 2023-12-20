package awsutil

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func UnmarshalSQSEvent[T interface{}](sqsMessage events.SQSMessage) (T, error) {
	var outputEvent T
	err := json.Unmarshal([]byte(sqsMessage.Body), &outputEvent)
	if err != nil {
		return *new(T), err
	}

	return outputEvent, nil
}
