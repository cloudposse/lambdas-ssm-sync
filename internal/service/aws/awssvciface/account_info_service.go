package awssvciface

type AccountInfoService interface {
	GetAccount() (string, error)
	GetEnabledRegions() ([]string, error)
	GetRegion() (string, error)
	GetAccountDetails() (*GetAccountDetailsResponse, error)
}

type GetAccountDetailsResponse struct {
	Account        string
	CurrentRegion  string
	EnabledRegions []string
}
