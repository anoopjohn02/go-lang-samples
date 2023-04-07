package service

import (
	"com/anoop/examples/internal/models"
)

type MeasurementService struct {
}

func NewMeasurementService() *MeasurementService {
	return &MeasurementService{}
}

func (s *MeasurementService) send(measurement models.Measurement) (*models.Measurement, error) {
	return &models.Measurement{}, nil
}

func (s *MeasurementService) get(id string) (*models.Measurement, error) {
	return &models.Measurement{}, nil
}

func (s *MeasurementService) getByDeviceId(deviceId string) (*[]models.Measurement, error) {
	return &[]models.Measurement{}, nil
}

func (s *MeasurementService) delete(id string) error {
	return nil
}
