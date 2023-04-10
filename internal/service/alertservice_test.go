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

func (d *dummyAlertRepository) Save(entity entity.Alerts) (*entity.Alerts, error) {
	args := d.Called()
	return nil, args.Error(0)
}

func (d *dummyAlertRepository) Get(id string) (*entity.Alerts, error) {
	args := d.Called()
	return nil, args.Error(0)
}

func (d *dummyAlertRepository) ByDeviceId(id string) (*[]entity.Alerts, error) {
	args := d.Called()
	return nil, args.Error(0)
}

func (d *dummyAlertRepository) Delete(id string) error {
	args := d.Called()
	return args.Error(0)
}

type dummyIotoMqttConnection struct {
	mock.Mock
}

func (m *dummyIotoMqttConnection) Publish(message models.IotoMessage) {
	m.Called()
}

func (m *dummyIotoMqttConnection) Sub(topic string) {
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
	dummyMqtt := new(dummyIotoMqttConnection)
	dummyMqtt.On("Publish", mock.Anything).Return(nil, nil)

	alertService := NewAlertService(dummyRepo, dummyMqtt)
	alert := &models.Alert{DeviceId: "deviceId"}
	result, err := alertService.Send(*alert)

	assert.Nil(t, err, "No error was expected but one was received")
	assert.NotNil(t, result, "Expected some data but nothing was received")
	assert.Equal(t, "deviceId", "Expected device id and recieved device id are different")
}
