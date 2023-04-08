package client

import (
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

type DeviceClient struct {
	cache       *cache.Cache
	accountHost string
}

func NewDeviceClient(host string) *DeviceClient {
	cache := cache.New(60*time.Minute, 60*time.Minute)
	return &DeviceClient{cache, host}
}

func (dc *DeviceClient) GetDeviceProfile(token string) (string, error) {
	furnaceId, found := dc.cache.Get(strconv.Itoa(deviceId))
	if found {
		return furnaceId.(string), nil
	}

	url := fmt.Sprintf("%v/devices/%v/attributes", dc.host, deviceId)

	log.Printf("Accessing URL : %v\n", url)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var furnaces []Furnace
	if err := json.Unmarshal(responseData, &furnaces); err != nil {
		return "", err
	}
	for _, furnace := range furnaces {
		dc.cache.Set(strconv.Itoa(deviceId), furnace.Value, cache.DefaultExpiration)
		return furnace.Value, nil
	}
	return "", &furnaceNotFoundError{deviceId}
}
