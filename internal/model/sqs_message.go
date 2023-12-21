package model

type SQSMessage struct {
	MessageId         string                 `json:"messageId"`
	ReceiptHandle     string                 `json:"receiptHandle"`
	Body              string                 `json:"body"`
	Attributes        SQSMessageAttributes   `json:"attributes"`
	MessageAttributes map[string]interface{} `json:"messageAttributes"`
	Md5OfBody         string                 `json:"md5OfBody"`
	EventSource       string                 `json:"eventSource"`
	EventSourceARN    string                 `json:"eventSourceARN"`
	AwsRegion         string                 `json:"awsRegion"`
}
