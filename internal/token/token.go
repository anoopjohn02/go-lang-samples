package token

import (
	"strings"

	"com/anoop/examples/internal/client"

	"github.com/gin-gonic/gin"
)

type TokenValidator struct {
	device *client.DeviceClient
}

func NewTokenValidator(deviceClient *client.DeviceClient) *TokenValidator {
	return &TokenValidator{device: deviceClient}
}

func (t *TokenValidator) TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := t.device.GetDeviceProfile(tokenString)
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
