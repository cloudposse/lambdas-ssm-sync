package service

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/cloudposse/ssm-syncronizer/internal/model"
	"github.com/cloudposse/ssm-syncronizer/internal/util"
	"github.com/cloudposse/ssm-syncronizer/mocks"
	"github.com/stretchr/testify/assert"
)

func getParametersValue(t *testing.T) []*ssm.Parameter {
	file := "../../fixtures/get-parameters-by-path-output.json"
	param, err := util.UnmarshalFile[[]*ssm.Parameter](file)
	if err != nil {
		t.Fatalf("Failed to unmarshal %s", file)
	}
	return param
}

func TestOrchestratorService_Sync(t *testing.T) {
	tests := []struct {
		name           string
		eventFile      string
		currentAccount string
		currentRegion  string
		expectError    bool
	}{
		{
			name:           "Foo",
			eventFile:      "../../fixtures/param-store-event-same-account-and-region.json",
			currentAccount: "111111222222",
			currentRegion:  "us-east-2",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			event, err := util.UnmarshalFile[model.ParameterStoreChangeEvent](tt.eventFile)
			if err != nil {
				t.Fatalf("Failed to unmarshal event %s", tt.eventFile)
			}

			getParametersValue := getParametersValue(t)

			accountSvc := &mocks.AccountServiceMock{GetAccountResponse: tt.currentAccount, GetRegionResponse: tt.currentRegion}
			ssmSvc := &mocks.SSMServiceMock{GetParametersByPathPagesOutput: getParametersValue}
			sqsSvc := &mocks.SQSServiceMock{}

			svc := &OrchestratorService{AccountService: accountSvc, ConfigPath: "/terraform/config", SSMService: ssmSvc, SQSService: sqsSvc}

			err = svc.Sync(event)
			if (err != nil) != tt.expectError {
				t.Errorf("parseArn() error = %v, expectedErr %v", err, tt.expectError)
				return
			}

			assert.Equal(t, 2, sqsSvc.SendMessageCalls)
			assert.Equal(t, 2, sqsSvc.SetQueueUrlCalls)
		})
	}
}
