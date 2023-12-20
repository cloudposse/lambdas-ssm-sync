package model

import (
	"time"
)

type ParameterStoreChangeEvent EventBridgeEvent[*ParameterStoreChangeDetail]

func (a *ParameterStoreChangeEvent) SetDetail(detail ParameterStoreChangeDetail) {
	a.Detail = &detail
}

func (a *ParameterStoreChangeEvent) SetDetailType(detailType string) {
	a.DetailType = detailType
}

func (a *ParameterStoreChangeEvent) SetResources(resources []string) {
	a.Resources = resources
}

func (a *ParameterStoreChangeEvent) SetId(id string) {
	a.Id = id
}

func (a *ParameterStoreChangeEvent) SetSource(source string) {
	a.Source = source
}

func (a *ParameterStoreChangeEvent) SetTime(time time.Time) {
	a.Time = time
}

func (a *ParameterStoreChangeEvent) SetRegion(region string) {
	a.Region = region
}

func (a *ParameterStoreChangeEvent) SetVersion(version string) {
	a.Version = version
}

func (a *ParameterStoreChangeEvent) SetAccount(account string) {
	a.Account = account
}
