package model

type Config struct {
	SSMBasePath          string `json:"ssm_base_path"`
	SSMSharedPath        string `json:"ssm_shared_config_path"`
	OrchestratorQueueURL string `json:"orchestrator_queue_url"`
}
