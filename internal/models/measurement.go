package models

import "time"

type Measurement struct {
	Id       string    `json:"id"`
	DeviceId string    `json:"deviceId"`
	Type     string    `json:"type"`
	Value    int       `json:"value"`
	Unit     string    `json:"unit"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}
