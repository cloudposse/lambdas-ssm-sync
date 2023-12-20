package service

import (
	"fmt"

	"github.com/cloudposse/ssm-syncronizer/internal/model"
	"github.com/cloudposse/ssm-syncronizer/internal/service/aws/awssvciface"
	awsutil "github.com/cloudposse/ssm-syncronizer/internal/util/aws"
)

type CurrentAccountSyncService struct {
	AccountService awssvciface.AccountInfoService
	SSMService     awssvciface.SSMService
}

func NewCurrentAccountSyncService(accountService awssvciface.AccountInfoService, ssmService awssvciface.SSMService) *CurrentAccountSyncService {
	return &CurrentAccountSyncService{
		AccountService: accountService,
		SSMService:     ssmService,
	}
}

func (s *CurrentAccountSyncService) enrichEventWithParameterValue(event *model.ParameterStoreChangeEvent) (model.ParameterStoreChangeEvent, error) {
	value, err := s.SSMService.GetParameter(event.Resources[0])
	if err != nil {
		return *event, err
	}

	event.Detail.SetValue(value.String())

	return *event, nil
}

// isAllowedUpdate determines if the event source is the current account and region
func (s *CurrentAccountSyncService) isAllowedUpdate(currentAccount string, currentRegion string, event model.ParameterStoreChangeEvent) (bool, error) {
	paramArnParts, err := awsutil.GetSSMParameterARNParts(event.Resources[0])
	if err != nil {
		return false, err
	}

	// If the source of the event is the current account
	sourceAccount := paramArnParts.SourceAccount
	sourceAccountMatches := sourceAccount == currentAccount

	if !sourceAccountMatches {
		return false, fmt.Errorf("sync of parameter failed because the source account of the change (%s) is not the current account", sourceAccount)
	}

	// If the source of the event is the current region
	sourceRegion := paramArnParts.SourceRegion
	sourceRegionMatches := sourceRegion == currentRegion

	if !sourceRegionMatches {
		return false, fmt.Errorf("sync of parameter failed because the source region of the change (%s) is not the current region", sourceRegion)
	}

	// If the path of the parameter that is being updated matches the current account and region
	//
	// For security reasons, we don't want to allow local updates to parameters that are not in the SSM path for the
	// current account and region to be replicated to other regions and/or accounts.
	destinationAccount := paramArnParts.PathAccount
	destinationAccountMatches := destinationAccount == currentAccount

	if !destinationAccountMatches {
		return false, fmt.Errorf("sync of parameter failed because the destination path account (%s) does not match the source account of the change (%s)", destinationAccount, sourceAccount)
	}

	destinationRegion := paramArnParts.PathRegion
	destinationRegionMatches := destinationRegion == currentRegion

	if !destinationRegionMatches {
		return false, fmt.Errorf("sync of parameter failed because the destination path region (%s) does not match the source region of the change (%s)", destinationRegion, currentRegion)
	}

	currentRegionChange := sourceAccountMatches && sourceRegionMatches && destinationAccountMatches && destinationRegionMatches

	return currentRegionChange, nil
}

func (s *CurrentAccountSyncService) Sync(event model.ParameterStoreChangeEvent) error {
	eventWithValue, err := s.enrichEventWithParameterValue(&event)
	if err != nil {
		return err
	}

	details, err := s.AccountService.GetAccountDetails()
	if err != nil {
		return err
	}

	currentRegion := details.CurrentRegion
	isLocalUpdate, err := s.isAllowedUpdate(details.Account, currentRegion, eventWithValue)
	if err != nil {
		return err
	}

	for _, region := range details.EnabledRegions {
		// If the event is from the current region, then we don't need to update it.
		if isLocalUpdate {
			if region == currentRegion {
				fmt.Printf("Skipping update of %s in %s because the source of the change is the current region\n", event.Detail.Name, region)
				continue
			}
			s.SSMService.SetRegion(region)
			_, err := s.SSMService.SetParameter(event.Detail.Name, event.Detail.Type, event.Detail.Value)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("Skipping update of %s in %s because it is not in the current region\n", event.Detail.Name, region)
		}
	}

	return nil
}
