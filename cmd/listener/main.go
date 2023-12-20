package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/cloudposse/ssm-syncronizer/internal/model"
	"github.com/cloudposse/ssm-syncronizer/internal/service"
	awssvc "github.com/cloudposse/ssm-syncronizer/internal/service/aws"
	awsutil "github.com/cloudposse/ssm-syncronizer/internal/util/aws"
)

func handler(context context.Context, event events.SQSEvent) ([]byte, error) {

	// Instantiate concrete dependencies
	session := session.Must(session.NewSession())
	ec2Client := awssvc.NewEC2Client(session)
	stsClient := awssvc.NewSTSClient(session)
	accountService := awssvc.NewAccountService(*session.Config.Region, ec2Client, stsClient)
	ssmService := awssvc.NewSSMService()

	currentAccountSyncService := service.NewCurrentAccountSyncService(accountService, ssmService)

	for _, record := range event.Records {
		ssmEvent, err := awsutil.UnmarshalSQSEvent[model.ParameterStoreChangeEvent](record)
		if err != nil {
			return nil, err
		}
		currentAccountSyncService.Sync(ssmEvent)
	}
	return nil, nil
}

func main() {
	lambda.Start(handler)
}
