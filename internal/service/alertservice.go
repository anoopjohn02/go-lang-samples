package service

import (
	"com/anoop/examples/internal/data/entity"
	"com/anoop/examples/internal/data/repo"
	"com/anoop/examples/internal/models"
	"encoding/json"
	"fmt"
	"log"
)

type AlertService struct {
	repo           repo.AlertRepository
	messageService MessageService
}

func NewAlertService(repo repo.AlertRepository, messageService MessageService) *AlertService {
	return &AlertService{repo: repo, messageService: messageService}
}

func (s *AlertService) Send(alert models.Alert) (*models.Alert, error) {
	alerts := &entity.Alerts{
		DeviceId:    alert.DeviceId,
		Type:        alert.Type,
		Severity:    alert.Severity,
		Key:         alert.Key,
		Description: alert.Description,
		DateTime:    alert.DateTime,
	}

	result, error := s.repo.Save(*alerts)
	if error != nil {
		log.Printf("Unable to insert data into database: %v\n", error)
		return &models.Alert{}, error
	}
	alert.UId = result.ID.Hex()
	jsonAlert, err := json.Marshal(alert)
	if err != nil {
		fmt.Println(err)
	}
	message := models.IotoMessage{
		MessageType: "ALERT",
		DeviceId:    alert.DeviceId,
		Message:     string(jsonAlert),
	}
	s.messageService.Publish(message)

	return GetModel(result), nil
}

func (s *AlertService) Get(id string) (*models.Alert, error) {
	result, error := s.repo.Get(id)
	if error != nil {
		log.Printf("Unable to get data from database: %v\n", error)
		return &models.Alert{}, error
	}
	return GetModel(result), nil
}

func (s *AlertService) GetByDeviceId(deviceId string) (*[]models.Alert, error) {
	result, error := s.repo.ByDeviceId(deviceId)
	if error != nil {
		log.Printf("Unable to get data from database: %v\n", error)
		return &[]models.Alert{}, error
	}

	res := []models.Alert{}
	for _, entity := range *result {
		res = append(res, *GetModel(&entity))
	}
	return &res, nil
}

func (s *AlertService) Delete(id string) error {
	error := s.repo.Delete(id)
	if error != nil {
		log.Printf("Unable to delete data from database: %v\n", error)
		return error
	}
	return nil
}

func GetModel(entity *entity.Alerts) *models.Alert {
	return &models.Alert{
		Id:          entity.ID.Hex(),
		UId:         entity.ID.Hex(),
		DeviceId:    entity.DeviceId,
		Type:        entity.Type,
		Severity:    entity.Severity,
		Key:         entity.Key,
		Description: entity.Description,
		DateTime:    entity.DateTime,
	}
}
