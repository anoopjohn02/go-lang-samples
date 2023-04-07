package repo

import "com/anoop/examples/internal/data/entity"

type MeasurementRepository interface {
	Save(entity entity.Measurements) (*entity.Measurements, error)
	Get(id string) (*entity.Measurements, error)
	ByDeviceId(id string) (*[]entity.Measurements, error)
	Delete(id string) error
}
