package mocks

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type EC2ClientMock struct {
	ec2iface.EC2API
	DescribeRegionsOutput ec2.DescribeRegionsOutput
}

func (m *EC2ClientMock) DescribeRegions(*ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	return &m.DescribeRegionsOutput, nil
}
