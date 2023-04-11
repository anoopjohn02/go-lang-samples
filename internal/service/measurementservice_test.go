package service

import (
	"com/anoop/examples/internal/data/entity"
	"com/anoop/examples/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dummyMeasurementRepository struct {
	mock.Mock
}

func (d *dummyMeasurementRepository) Save(ent entity.Measurements) (*entity.Measurements, error) {
	args := d.Called(ent)
	return args.Get(0).(*entity.Measurements), args.Error(1)
}

func (d *dummyMeasurementRepository) Get(id string) (*entity.Measurements, error) {
	args := d.Called(id)
	return args.Get(0).(*entity.Measurements), args.Error(1)
}

func (d *dummyMeasurementRepository) ByDeviceId(id string) (*[]entity.Measurements, error) {
	args := d.Called(id)
	return args.Get(0).(*[]entity.Measurements), args.Error(1)
}

func (d *dummyMeasurementRepository) Delete(id string) error {
	args := d.Called(id)
	return args.Error(0)
}

func TestMeasurementSend(t *testing.T) {
	measurements := &entity.Measurements{
		ID:       primitive.NewObjectID(),
		DeviceId: "deviceId",
		Type:     "TEST",
	}
	dummyRepo := new(dummyMeasurementRepository)
	dummyRepo.On("Save", mock.Anything).Return(measurements, nil)
	dummyAlertService := dummyAlertService()

	measurementService := NewMeasurementService(dummyRepo, *dummyAlertService)
	t.Run("normal", func(t *testing.T) {
		measurement := &models.Measurement{DeviceId: "deviceId", Value: 500}
		result, err := measurementService.Send(*measurement)

		assert.Nil(t, err, "No error was expected but one was received")
		assert.NotNil(t, result, "Expected some data but nothing was received")
		assert.Equal(t, "deviceId", result.DeviceId, "Expected device id and recieved device id are different")
	})
	t.Run("trigger alert", func(t *testing.T) {
		measurement := &models.Measurement{DeviceId: "deviceId", Value: 1500}
		result, err := measurementService.Send(*measurement)

		assert.Nil(t, err, "No error was expected but one was received")
		assert.NotNil(t, result, "Expected some data but nothing was received")
		assert.Equal(t, "deviceId", result.DeviceId, "Expected device id and recieved device id are different")
	})
}

func TestMeasurementGet(t *testing.T) {
	measurements := &entity.Measurements{
		ID:       primitive.NewObjectID(),
		DeviceId: "deviceId",
		Type:     "TEST",
	}
	id := "object_id"
	dummyRepo := new(dummyMeasurementRepository)
	dummyRepo.On("Get", id).Return(measurements, nil)
	dummyAlertService := dummyAlertService()

	measurementService := NewMeasurementService(dummyRepo, *dummyAlertService)
	result, err := measurementService.Get(id)

	assert.Nil(t, err, "No error was expected but one was received")
	assert.NotNil(t, result, "Expected some data but nothing was received")
	assert.Equal(t, "deviceId", result.DeviceId, "Expected device id and recieved device id are different")
}

func TestMeasurementGetByDeviceId(t *testing.T) {
	measurements := entity.Measurements{
		ID:       primitive.NewObjectID(),
		DeviceId: "deviceId",
		Type:     "TEST",
	}
	id := "device_id"
	res := []entity.Measurements{}
	res = append(res, measurements)

	dummyRepo := new(dummyMeasurementRepository)
	dummyRepo.On("ByDeviceId", id).Return(&res, nil)
	dummyAlertService := dummyAlertService()

	measurementService := NewMeasurementService(dummyRepo, *dummyAlertService)
	result, err := measurementService.GetByDeviceId(id)

	assert.Nil(t, err, "No error was expected but one was received")
	assert.NotNil(t, result, "Expected some data but nothing was received")
	assert.Equal(t, 1, len(*result), "Expected one result but none or more results was received")
}

func dummyAlertService() *AlertService {
	alerts := &entity.Alerts{
		ID:       primitive.NewObjectID(),
		DeviceId: "deviceId",
	}
	dummyAlertRepo := new(dummyAlertRepository)
	dummyAlertRepo.On("Save", mock.Anything).Return(alerts, nil)
	dummyMqtt := new(dummyMessageService)
	dummyMqtt.On("Publish", mock.Anything).Return(nil, nil)
	alertSer := NewAlertService(dummyAlertRepo, dummyMqtt)
	return alertSer
}
