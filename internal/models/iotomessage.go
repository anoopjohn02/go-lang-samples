package models

type IotoMessage struct {
	MessageType string `json:"messageType"`
	DeviceId    string `json:"deviceId"`
	Message     string `json:"message"`
}
