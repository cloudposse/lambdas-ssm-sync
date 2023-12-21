package awssvc

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/cloudposse/ssm-syncronizer/internal/util"
	"github.com/cloudposse/ssm-syncronizer/mocks"
	"github.com/stretchr/testify/assert"
)

func getEnabledRegions(t *testing.T) []string {
	regionsFile := "../../../fixtures/enabled-regions.json"
	regions, err := util.UnmarshalFile[[]string](regionsFile)
	if err != nil {
		t.Fatalf("Failed to unmarshal event %s", regionsFile)
	}
	return regions
}

func getEC2ClientMock() *mocks.EC2ClientMock {
	return &mocks.EC2ClientMock{DescribeRegionsOutput: ec2.DescribeRegionsOutput{
		Regions: []*ec2.Region{

			{
				Endpoint:    aws.String("ec2.ap-south-1.amazonaws.com"),
				RegionName:  aws.String("ap-south-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.eu-north-1.amazonaws.com"),
				RegionName:  aws.String("eu-north-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.eu-west-3.amazonaws.com"),
				RegionName:  aws.String("eu-west-3"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.eu-west-2.amazonaws.com"),
				RegionName:  aws.String("eu-west-2"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.eu-west-1.amazonaws.com"),
				RegionName:  aws.String("eu-west-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.ap-northeast-3.amazonaws.com"),
				RegionName:  aws.String("ap-northeast-3"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.ap-northeast-2.amazonaws.com"),
				RegionName:  aws.String("ap-northeast-2"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.ap-northeast-1.amazonaws.com"),
				RegionName:  aws.String("ap-northeast-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.ca-central-1.amazonaws.com"),
				RegionName:  aws.String("ca-central-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.sa-east-1.amazonaws.com"),
				RegionName:  aws.String("sa-east-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.ap-southeast-1.amazonaws.com"),
				RegionName:  aws.String("ap-southeast-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.ap-southeast-2.amazonaws.com"),
				RegionName:  aws.String("ap-southeast-2"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.eu-central-1.amazonaws.com"),
				RegionName:  aws.String("eu-central-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.us-east-1.amazonaws.com"),
				RegionName:  aws.String("us-east-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.us-east-2.amazonaws.com"),
				RegionName:  aws.String("us-east-2"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.us-west-1.amazonaws.com"),
				RegionName:  aws.String("us-west-1"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
			{
				Endpoint:    aws.String("ec2.us-west-2.amazonaws.com"),
				RegionName:  aws.String("us-west-2"),
				OptInStatus: aws.String("opt-in-not-required"),
			},
		},
	}}
}

func getSTSClientMock() *mocks.STSClientMock {
	return &mocks.STSClientMock{GetCallerIdentityOutout: sts.GetCallerIdentityOutput{
		Account: aws.String("123456789012"),
	}}
}

func TestGetEnabledRegions(t *testing.T) {
	svc := AccountInfoService{ec2Client: getEC2ClientMock()}

	regions, err := svc.GetEnabledRegions()

	assert.Equal(t, err, nil)
	assert.ElementsMatch(t, regions, getEnabledRegions(t))
}

func TestGetCallerIdentity(t *testing.T) {
	svc := AccountInfoService{stsClient: getSTSClientMock()}
	account, err := svc.GetAccount()

	assert.Equal(t, err, nil)
	assert.Equal(t, account, "123456789012")
}

func TestGetRegion(t *testing.T) {
	svc := AccountInfoService{currentRegion: "us-east-1"}
	region, err := svc.GetRegion()

	assert.Equal(t, err, nil)
	assert.Equal(t, region, "us-east-1")
}

func TestGetAccountDetails(t *testing.T) {
	svc := AccountInfoService{currentRegion: "us-east-1", ec2Client: getEC2ClientMock(), stsClient: getSTSClientMock()}
	details, err := svc.GetAccountDetails()

	assert.Equal(t, err, nil)
	assert.Equal(t, details.Account, "123456789012")
	assert.Equal(t, details.CurrentRegion, "us-east-1")
	assert.ElementsMatch(t, details.EnabledRegions, getEnabledRegions(t))
}
