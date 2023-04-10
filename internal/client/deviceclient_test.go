package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

type dummyHttpClient struct {
	mock.Mock
}

func (d *dummyHttpClient) Do(req *http.Request) (string, error) {
	args := d.Called()
	return "", args.Error(0)
}

func TestGetDeviceProfile(t *testing.T) {
	/*profile := &models.DeviceProfile{UserName: "device"}
	response := &http.Response{Body: profile}
	client := new(dummyHttpClient)
	client.On("Do", mock.Anything).Return(response, nil)*/
}
