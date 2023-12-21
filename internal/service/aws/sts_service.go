package awssvc

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func NewSTSClient(sess *session.Session) *sts.STS {
	return sts.New(session.Must(sess, nil))
}
