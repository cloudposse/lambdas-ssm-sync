package awssvc

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func NewEC2Client(sess *session.Session) *ec2.EC2 {
	return ec2.New(session.Must(sess, nil))
}
