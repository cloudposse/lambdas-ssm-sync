package model

type SQSMessageAttributes struct {
	ApproximateReceiveCount          string `json:"ApproximateReceiveCount"`
	SentTimestamp                    string `json:"SentTimestamp"`
	SenderId                         string `json:"SenderId"`
	ApproximateFirstReceiveTimestamp string `json:"ApproximateFirstReceiveTimestamp"`
}
