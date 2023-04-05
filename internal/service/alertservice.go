package service

import (
	"com/anoop/examples/internal/models"
)

type AlertService struct {
}

func NewAlertService() *AlertService {
	return &AlertService{}
}

func (s *AlertService) send(alert models.Alert) models.Alert {

}

func (s *AlertService) get(id string) models.Alert {

}

func (s *AlertService) getByDeviceId(deviceId string) models.Alert {

}

func (s *AlertService) delete(id string) {

}
