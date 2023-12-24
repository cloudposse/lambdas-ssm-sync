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

	configSubPath := os.Getenv("SSM_CONFIG_PATH")
	if configSubPath == "" {
		configSubPath = "config"
	}
	configPath := fmt.Sprintf("%s/%s", basePath, configSubPath)

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
		SSMConfigPath:        configPath,
		SSMSharedPath:        sharedPath,
		OrchestratorQueueURL: queueUrl,
	}, nil
}
