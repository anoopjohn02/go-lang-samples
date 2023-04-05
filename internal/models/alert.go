package models

import "time"

type Alert struct {
	Id          string    `json:"id"`
	UId         string    `json:"uId"`
	DeviceId    string    `json:"deviceId"`
	Type        string    `json:"type"`
	Severity    string    `json:"severity"`
	Key         string    `json:"key"`
	Description string    `json:"description"`
	DateTime    time.Time `json:"dateTime"`
}
