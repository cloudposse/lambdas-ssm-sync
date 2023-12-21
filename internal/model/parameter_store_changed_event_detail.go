package model

type ParameterStoreChangeDetail struct {
	Operation          string `json:"operation"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	Description        string `json:"description"`
	Value              string `json:"value"`
	IsFromOrchestrator bool   `json:"isFromOrchestrator"`
}

func (e *ParameterStoreChangeDetail) SetOperation(operation string) {
	e.Operation = operation
}

func (e *ParameterStoreChangeDetail) SetName(name string) {
	e.Name = name
}

func (e *ParameterStoreChangeDetail) SetType(parameterType string) {
	e.Type = parameterType
}

func (e *ParameterStoreChangeDetail) SetDescription(description string) {
	e.Description = description
}

func (e *ParameterStoreChangeDetail) SetValue(value string) {
	e.Value = value
}

func (e *ParameterStoreChangeDetail) SetIsFromOrchestrator(value bool) {
	e.IsFromOrchestrator = value
}
