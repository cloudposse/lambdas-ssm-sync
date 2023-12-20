package service

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/cloudposse/ssm-syncronizer/internal/model"
	"github.com/cloudposse/ssm-syncronizer/internal/util"
	"github.com/cloudposse/ssm-syncronizer/mocks"
	"github.com/stretchr/testify/assert"
)

func getEnabledRegions(t *testing.T) []string {
	regionsFile := "../../fixtures/enabled-regions.json"
	regions, err := util.UnmarshalFile[[]string](regionsFile)
	if err != nil {
		t.Fatalf("Failed to unmarshal event %s", regionsFile)
	}
	return regions
}

func getParameterValue(t *testing.T) ssm.GetParameterOutput {
	file := "../../fixtures/get-parameter-output.json"
	param, err := util.UnmarshalFile[ssm.GetParameterOutput](file)
	if err != nil {
		t.Fatalf("Failed to unmarshal %s", file)
	}
	return param
}

func TestCurrentAccountService_Sync(t *testing.T) {
	tests := []struct {
		name                      string
		eventFile                 string
		currentAccount            string
		currentRegion             string
		expectError               bool
		expectedSetParameterCalls int
		expectedSetRegionCalls    int
	}{
		{
			name:                      "UpdateFromSameAccountAndRegion",
			eventFile:                 "../../fixtures/param-store-event-same-account-and-region.json",
			currentAccount:            "123456789012",
			currentRegion:             "us-east-2",
			expectedSetParameterCalls: 16,
			expectedSetRegionCalls:    16,
			expectError:               false,
		},
		{
			name:                      "UpdateFromSameAccountDifferentRegion",
			eventFile:                 "../../fixtures/param-store-event-same-account-and-region.json",
			currentAccount:            "123456789012",
			currentRegion:             "us-east-1",
			expectedSetParameterCalls: 0,
			expectedSetRegionCalls:    0,
			expectError:               true,
		},
		{
			name:                      "UpdateFromDifferentAccount",
			eventFile:                 "../../fixtures/param-store-event-different-account.json",
			currentAccount:            "123456789012",
			currentRegion:             "us-east-2",
			expectedSetParameterCalls: 0,
			expectedSetRegionCalls:    0,
			expectError:               true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			event, err := util.UnmarshalFile[model.ParameterStoreChangeEvent](tt.eventFile)
			if err != nil {
				t.Fatalf("Failed to unmarshal event %s", tt.eventFile)
			}

			enabledRegions := getEnabledRegions(t)
			parameterOutput := getParameterValue(t)

			accountSvc := &mocks.AccountServiceMock{GetAccountResponse: tt.currentAccount, GetRegionResponse: tt.currentRegion, GetEnabledRegionsResponse: enabledRegions}
			ssmSvc := &mocks.SSMServiceMock{GetParameterOutput: parameterOutput}

			svc := &CurrentAccountSyncService{AccountService: accountSvc, SSMService: ssmSvc}

			err = svc.Sync(event)
			if (err != nil) != tt.expectError {
				t.Errorf("parseArn() error = %v, expectedErr %v", err, tt.expectError)
				return
			}

			assert.Equal(t, tt.expectedSetParameterCalls, len(ssmSvc.SetParameterCalls))
			assert.Equal(t, tt.expectedSetRegionCalls, len(ssmSvc.SetRegionCalls))
		})
	}
}
