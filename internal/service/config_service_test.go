package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigService(t *testing.T) {
	t.Setenv("ORCHESTRATOR_QUEUE_URL", "https://sqs.us-east-1.amazonaws.com/123456789012/orchestrator")

	svc, err := GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, "/terraform", svc.SSMBasePath)
	assert.Equal(t, "/terraform/shared", svc.SSMSharedPath)
	assert.Equal(t, "https://sqs.us-east-1.amazonaws.com/123456789012/orchestrator", svc.OrchestratorQueueURL)
}

func TestConfigService_OverrideBasePath(t *testing.T) {
	t.Setenv("ORCHESTRATOR_QUEUE_URL", "https://sqs.us-east-1.amazonaws.com/123456789012/orchestrator")
	t.Setenv("SSM_BASE_PATH", "/terraform/override")

	svc, err := GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, "/terraform/override", svc.SSMBasePath)
	assert.Equal(t, "/terraform/override/shared", svc.SSMSharedPath)
	assert.Equal(t, "https://sqs.us-east-1.amazonaws.com/123456789012/orchestrator", svc.OrchestratorQueueURL)
}

func TestConfigService_OverrideSharedPath(t *testing.T) {
	t.Setenv("ORCHESTRATOR_QUEUE_URL", "https://sqs.us-east-1.amazonaws.com/123456789012/orchestrator")
	t.Setenv("SSM_SHARED_PATH", "public")

	svc, err := GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, "/terraform", svc.SSMBasePath)
	assert.Equal(t, "/terraform/public", svc.SSMSharedPath)
}
