package service

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/cloudposse/ssm-syncronizer/internal/model"
	"github.com/cloudposse/ssm-syncronizer/internal/service/aws/awssvciface"
	"github.com/cloudposse/ssm-syncronizer/internal/util"
)

type OrchestratorService struct {
	AccountService awssvciface.AccountInfoService
	ConfigPath     string
	SSMService     awssvciface.SSMService
	SQSService     awssvciface.SQSService
}

type enabledAccount struct {
	Account  string `json:"account"`
	QueueUrl string `json:"queue_url"`
}

func NewOrchestratorService(accountService awssvciface.AccountInfoService, configPath string, ssmService awssvciface.SSMService, sqsService awssvciface.SQSService) *OrchestratorService {
	return &OrchestratorService{
		AccountService: accountService,
		ConfigPath:     configPath,
		SSMService:     ssmService,
		SQSService:     sqsService,
	}
}

func (s *OrchestratorService) getEnabledListenerAccounts() ([]enabledAccount, error) {
	region, err := s.AccountService.GetRegion()
	if err != nil {
		return nil, err
	}

	accountsEnabledPath := strings.Join([]string{s.ConfigPath, "listeners", region}, "/")

	var accounts []enabledAccount

	input := &ssm.GetParametersByPathInput{
		Path: aws.String(accountsEnabledPath),
	}

	params, err := s.SSMService.GetParametersByPathPages(input)
	for _, param := range params {
		account := strings.Split(*param.Name, fmt.Sprintf("%s/", accountsEnabledPath))[1]
		accounts = append(accounts, enabledAccount{Account: account, QueueUrl: *param.Value})
	}

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *OrchestratorService) sendToRemoteListener(account string, queue string, event model.ParameterStoreChangeEvent) error {
	message, err := util.Marshal(event)
	if err != nil {
		return err
	}

	err = s.SQSService.SetQueueUrl(queue)
	if err != nil {
		return err
	}

	s.SQSService.SendMessage(string(message), event.Detail.Name)
	return nil
}

func (s *OrchestratorService) Sync(event model.ParameterStoreChangeEvent) error {
	accounts, err := s.getEnabledListenerAccounts()
	if err != nil {
		return err
	}

	for _, account := range accounts {
		err := s.sendToRemoteListener(account.Account, account.QueueUrl, event)
		if err != nil {
			return err
		}
	}

	return nil
}
