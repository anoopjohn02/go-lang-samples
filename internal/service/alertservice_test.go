package service

import (
	"com/anoop/examples/internal/data/entity"

	"github.com/stretchr/testify/mock"
)

type dummyAlertRepository struct {
	mock.Mock
}

func (d *dummyAlertRepository) Save(entity entity.Alerts) (entity.Alerts, error) {
	args := d.Called()
	return &entity.Alerts{}, args.Error(0)
}
