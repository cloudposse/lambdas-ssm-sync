package model

import "time"

type EventBridgeEvent[T any] struct {
	Detail     T         `json:"detail"`
	DetailType string    `json:"detail-type"`
	Resources  []string  `json:"resources"`
	Id         string    `json:"id"`
	Source     string    `json:"source"`
	Time       time.Time `json:"time"`
	Region     string    `json:"region"`
	Version    string    `json:"version"`
	Account    string    `json:"account"`
}
