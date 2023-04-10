package service

import (
	"com/anoop/examples/internal/data/entity"
	"com/anoop/examples/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dummyAlertRepository struct {
	mock.Mock
}

func (d *dummyAlertRepository) Save(ent entity.Alerts) (*entity.Alerts, error) {
	args := d.Called(ent)
	return args.Get(0).(*entity.Alerts), args.Error(1)
}

func (d *dummyAlertRepository) Get(id string) (*entity.Alerts, error) {
	args := d.Called(id)
	return args.Get(0).(*entity.Alerts), args.Error(1)
}

func (d *dummyAlertRepository) ByDeviceId(id string) (*[]entity.Alerts, error) {
	args := d.Called(id)
	return args.Get(0).(*[]entity.Alerts), args.Error(1)
}

func (d *dummyAlertRepository) Delete(id string) error {
	args := d.Called(id)
	return args.Error(0)
}

type dummyMessageService struct {
	mock.Mock
}

func (m *dummyMessageService) Publish(message models.IotoMessage) {
	m.Called()
}

func (m *dummyMessageService) Sub(topic string) {
	m.Called()
}

func TestSend(t *testing.T) {
	alerts := &entity.Alerts{
		ID:          primitive.NewObjectID(),
		DeviceId:    "deviceId",
		Type:        "TEST",
		Severity:    "MAJOR",
		Key:         "key",
		Description: "description",
		DateTime:    time.Now(),
	}
	dummyRepo := new(dummyAlertRepository)
	dummyRepo.On("Save", mock.Anything).Return(alerts, nil)
	dummyMqtt := new(dummyMessageService)
	dummyMqtt.On("Publish", mock.Anything).Return(nil, nil)

	alertService := NewAlertService(dummyRepo, dummyMqtt)
	alert := &models.Alert{DeviceId: "deviceId"}
	result, err := alertService.Send(*alert)

	assert.Nil(t, err, "No error was expected but one was received")
	assert.NotNil(t, result, "Expected some data but nothing was received")
	assert.Equal(t, "deviceId", result.DeviceId, "Expected device id and recieved device id are different")
}

func TestGet(t *testing.T) {
	alerts := &entity.Alerts{
		ID:          primitive.NewObjectID(),
		DeviceId:    "deviceId",
		Type:        "TEST",
		Severity:    "MAJOR",
		Key:         "key",
		Description: "description",
		DateTime:    time.Now(),
	}
	id := "object_id"
	dummyRepo := new(dummyAlertRepository)
	dummyRepo.On("Get", id).Return(alerts, nil)
	dummyMqtt := new(dummyMessageService)

	alertService := NewAlertService(dummyRepo, dummyMqtt)
	result, err := alertService.Get(id)

	assert.Nil(t, err, "No error was expected but one was received")
	assert.NotNil(t, result, "Expected some data but nothing was received")
	assert.Equal(t, "deviceId", result.DeviceId, "Expected device id and recieved device id are different")
}

func TestGetByDeviceId(t *testing.T) {
	ent := entity.Alerts{
		ID:          primitive.NewObjectID(),
		DeviceId:    "deviceId",
		Type:        "TEST",
		Severity:    "MAJOR",
		Key:         "key",
		Description: "description",
		DateTime:    time.Now(),
	}
	id := "device_id"
	res := []entity.Alerts{}
	res = append(res, ent)

	dummyRepo := new(dummyAlertRepository)
	dummyRepo.On("ByDeviceId", id).Return(&res, nil)
	dummyMqtt := new(dummyMessageService)

	alertService := NewAlertService(dummyRepo, dummyMqtt)
	result, err := alertService.GetByDeviceId(id)

	assert.Nil(t, err, "No error was expected but one was received")
	assert.NotNil(t, result, "Expected some data but nothing was received")
	assert.Equal(t, 1, len(*result), "Expected one result but none or more results was received")
}
