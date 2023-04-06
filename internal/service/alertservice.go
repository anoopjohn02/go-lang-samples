package service

import (
	"com/anoop/examples/internal/data/repo"
	"com/anoop/examples/internal/data/repo/mongorepo"
	"com/anoop/examples/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type AlertService struct {
	repo repo.AlertRepository
}

func NewAlertService(mongo *mongo.Database) *AlertService {
	repo := mongorepo.NewAlertRepository(mongo)
	return &AlertService{repo: repo}
}

func (s *AlertService) send(alert models.Alert) models.Alert {

}

func (s *AlertService) get(id string) models.Alert {

}

func (s *AlertService) getByDeviceId(deviceId string) models.Alert {

}

func (s *AlertService) delete(id string) {

}
