package model

type Config struct {
	OrchestratorQueueURL string `json:"orchestrator_queue_url"`
	SSMBasePath          string `json:"ssm_base_path"`
	SSMSharedPath        string `json:"ssm_shared_config_path"`
}
