package service

import (
	"com/anoop/examples/internal/models"
)

type MeasurementService struct {
}

func NewMeasurementService() *MeasurementService {
	return &MeasurementService{}
}

func (s *MeasurementService) send(measurement models.Measurement) models.Measurement {

}

func (s *MeasurementService) get(id string) models.Measurement {

}

func (s *MeasurementService) getByDeviceId(deviceId string) models.Measurement {

}

func (s *MeasurementService) delete(id string) {

}
