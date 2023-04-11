package web

import (
	"com/anoop/examples/internal/models"
	"com/anoop/examples/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlertController struct {
	ser service.AlertService
}

func NewAlertController(service *service.AlertService) *AlertController {
	return &AlertController{ser: *service}
}

func (req *AlertController) Send(ctx *gin.Context) {
	var dataRequest models.Alert
	if err := ctx.ShouldBindJSON(&dataRequest); err != nil {
		req.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := req.ser.Send(dataRequest); err != nil {
		req.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.String(http.StatusOK, "")
}

func (s *AlertController) Get(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	alert, err := s.ser.Get(id)
	if err != nil {
		s.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, alert)
}

func (s *AlertController) GetByDeviceId(ctx *gin.Context) {
	device := ctx.MustGet("User").(*models.DeviceProfile)
	alerts, err := s.ser.GetByDeviceId(device.UserName)
	if err != nil {
		s.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, alerts)
}

func (s *AlertController) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	err := s.ser.Delete(id)
	if err != nil {
		s.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, "")
}

func (req *AlertController) finishWithError(ctx *gin.Context, status int, err error) {
	var response = struct {
		Error string `json:"error"`
	}{Error: err.Error()}

	ctx.JSON(status, response)
}
