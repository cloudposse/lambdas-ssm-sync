package awssvc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/cloudposse/ssm-syncronizer/internal/service/aws/awssvciface"
)

type AccountInfoService struct {
	currentRegion string
	ec2Client     ec2iface.EC2API
	stsClient     stsiface.STSAPI
}

func NewAccountService(currentRegion string, ec2Client ec2iface.EC2API, stsClient stsiface.STSAPI) *AccountInfoService {
	return &AccountInfoService{
		currentRegion: currentRegion,
		ec2Client:     ec2Client,
		stsClient:     stsClient,
	}
}

func (service *AccountInfoService) GetAccount() (string, error) {
	id, err := service.stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}

	return *id.Account, nil
}

// Get all enabled regions in the account
func (service *AccountInfoService) GetEnabledRegions() ([]string, error) {
	regions := []string{}
	regionOutput, err := service.ec2Client.DescribeRegions(&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)})
	if err != nil {
		return nil, err
	}

	for _, region := range regionOutput.Regions {
		regions = append(regions, *region.RegionName)
	}

	return regions, nil
}

func (service *AccountInfoService) GetRegion() (string, error) {
	return service.currentRegion, nil
}

func (service *AccountInfoService) GetAccountDetails() (*awssvciface.GetAccountDetailsResponse, error) {
	account, err := service.GetAccount()
	if err != nil {
		return nil, err
	}

	currentRegion, err := service.GetRegion()
	if err != nil {
		return nil, err
	}

	enabledRegions, err := service.GetEnabledRegions()
	if err != nil {
		return nil, err
	}

	return &awssvciface.GetAccountDetailsResponse{
		Account:        account,
		CurrentRegion:  currentRegion,
		EnabledRegions: enabledRegions,
	}, nil
}
