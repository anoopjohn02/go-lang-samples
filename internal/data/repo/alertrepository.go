package repo

import "com/anoop/examples/internal/data/entity"

type AlertRepository interface {
	Save(entity entity.Alerts) (*entity.Alerts, error)
	Get(id string) (*entity.Alerts, error)
	ByDeviceId(id string) (*[]entity.Alerts, error)
	Delete(id string) error
}
