package mocks

import "github.com/cloudposse/ssm-syncronizer/internal/service/aws/awssvciface"

type AccountServiceMock struct {
	GetAccountResponse        string
	GetRegionResponse         string
	GetEnabledRegionsResponse []string
	GetAccountDetailsResponse *awssvciface.GetAccountDetailsResponse
}

func (a *AccountServiceMock) GetAccount() (string, error) {
	return a.GetAccountResponse, nil
}

func (a *AccountServiceMock) GetEnabledRegions() ([]string, error) {
	return a.GetEnabledRegionsResponse, nil
}

func (a *AccountServiceMock) GetRegion() (string, error) {
	return a.GetRegionResponse, nil
}

func (a *AccountServiceMock) GetAccountDetails() (*awssvciface.GetAccountDetailsResponse, error) {
	if a.GetAccountDetailsResponse != nil {
		return a.GetAccountDetailsResponse, nil
	}

	account, err := a.GetAccount()
	if err != nil {
		return nil, err
	}

	region, err := a.GetRegion()
	if err != nil {
		return nil, err
	}

	enabledRegions, err := a.GetEnabledRegions()
	if err != nil {
		return nil, err
	}

	return &awssvciface.GetAccountDetailsResponse{
		Account:        account,
		CurrentRegion:  region,
		EnabledRegions: enabledRegions,
	}, nil
}
