package service

import (
	"com/anoop/examples/internal/data/entity"
	"com/anoop/examples/internal/data/repo"
	"com/anoop/examples/internal/models"
	"log"
	"time"
)

type MeasurementService struct {
	repo         repo.MeasurementRepository
	alertService AlertService
}

func NewMeasurementService(repo repo.MeasurementRepository, alertService AlertService) *MeasurementService {
	return &MeasurementService{repo: repo, alertService: alertService}
}

func (s *MeasurementService) Send(measurement models.Measurement) (*models.Measurement, error) {
	measurementEnt := &entity.Measurements{
		DeviceId: measurement.DeviceId,
		Type:     measurement.Type,
		Value:    measurement.Value,
		Unit:     measurement.Unit,
		Start:    measurement.Start,
		End:      measurement.End,
	}

	result, error := s.repo.Save(*measurementEnt)
	if error != nil {
		log.Printf("Unable to insert data into database: %v\n", error)
		return &models.Measurement{}, error
	}
	if measurement.Value > 1200 {
		excededAlert := models.Alert{
			DeviceId:    measurement.DeviceId,
			Type:        "ENERGY_CONSUMPTION_CROSSED",
			Severity:    "MAJOR",
			Key:         "app.notification.device.energy.usage.high",
			Description: "Energy Usage Crossed the limit",
			DateTime:    time.Now(),
		}
		s.alertService.Send(excededAlert)
	}
	return GetMeasurement(result), nil
}

func (s *MeasurementService) Get(id string) (*models.Measurement, error) {
	result, error := s.repo.Get(id)
	if error != nil {
		log.Printf("Unable to get data from database: %v\n", error)
		return &models.Measurement{}, error
	}
	return GetMeasurement(result), nil
}

func (s *MeasurementService) GetByDeviceId(deviceId string) (*[]models.Measurement, error) {
	result, error := s.repo.ByDeviceId(deviceId)
	if error != nil {
		log.Printf("Unable to get data from database: %v\n", error)
		return &[]models.Measurement{}, error
	}

	res := []models.Measurement{}
	for _, entity := range *result {
		res = append(res, *GetMeasurement(&entity))
	}
	return &res, nil
}

func (s *MeasurementService) Delete(id string) error {
	error := s.repo.Delete(id)
	if error != nil {
		log.Printf("Unable to delete data from database: %v\n", error)
		return error
	}
	return nil
}

func GetMeasurement(entity *entity.Measurements) *models.Measurement {
	return &models.Measurement{
		Id:       entity.ID.Hex(),
		DeviceId: entity.DeviceId,
		Type:     entity.Type,
		Value:    entity.Value,
		Unit:     entity.Unit,
		Start:    entity.Start,
		End:      entity.End,
	}
}
