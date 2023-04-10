package commons

import (
	"com/anoop/examples/internal/service"
	"com/anoop/examples/internal/token"
)

type DeviceContext struct {
	AlertService       *service.AlertService
	TokenValidator     *token.TokenValidator
	MeasurementService *service.MeasurementService
	MessageService     service.MessageService
}
