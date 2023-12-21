package service

import (
	"fmt"
	"os"

	"github.com/cloudposse/ssm-syncronizer/internal/model"
)

func GetConfig() (*model.Config, error) {
	basePath := os.Getenv("SSM_BASE_PATH")
	if basePath == "" {
		basePath = "/terraform"
	}

	sharedSubPath := os.Getenv("SSM_SHARED_PATH")
	if sharedSubPath == "" {
		sharedSubPath = "shared"
	}

	sharedPath := fmt.Sprintf("%s/%s", basePath, sharedSubPath)

	queueUrl := os.Getenv("ORCHESTRATOR_QUEUE_URL")
	if queueUrl == "" {
		return nil, fmt.Errorf("ORCHESTRATOR_QUEUE_URL is not set")
	}

	return &model.Config{
		SSMBasePath:          basePath,
		SSMSharedPath:        sharedPath,
		OrchestratorQueueURL: queueUrl,
	}, nil
}
