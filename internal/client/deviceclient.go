package client

import (
	"bytes"
	"com/anoop/examples/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type DeviceClient struct {
	cache       *cache.Cache
	accountHost string
	httpClient  *http.Client
}

func NewDeviceClient(host string) *DeviceClient {
	cache := cache.New(60*time.Minute, 60*time.Minute)
	client := http.Client{}
	return &DeviceClient{cache: cache, accountHost: host, httpClient: &client}
}

func (dc *DeviceClient) GetDeviceProfile(token string) (*models.DeviceProfile, error) {
	/*furnaceId, found := dc.cache.Get(strconv.Itoa(token))
	if found {
		return furnaceId.(string), nil
	}*/

	url := fmt.Sprintf("%v/user/info", dc.accountHost)
	log.Printf("Accessing URL : %v\n", url)

	bearer := "Bearer " + token
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Add("Authorization", bearer)
	response, err := dc.httpClient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return &models.DeviceProfile{}, err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &models.DeviceProfile{}, err
	}

	var profile models.DeviceProfile
	if err := json.Unmarshal(responseData, &profile); err != nil {
		return &models.DeviceProfile{}, err
	}

	//dc.cache.Set(strconv.Itoa(deviceId), furnace.Value, cache.DefaultExpiration)
	return &profile, nil
}
