package models

type IotoMessage struct {
	MessageType string `json:"messageType"`
	DeviceId    string `json:"deviceId"`
	Message     []byte `json:"message"`
}
