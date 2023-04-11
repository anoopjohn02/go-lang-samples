package web

import (
	"com/anoop/examples/internal/models"
	"com/anoop/examples/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MeasurementController struct {
	service service.MeasurementService
}

func NewMeasurementController(service service.MeasurementService) *MeasurementController {
	return &MeasurementController{service: service}
}

func (c *MeasurementController) Send(ctx *gin.Context) {
	var dataRequest models.Measurement
	if err := ctx.ShouldBindJSON(&dataRequest); err != nil {
		c.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := c.service.Send(dataRequest); err != nil {
		c.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.String(http.StatusOK, "")
}

func (c *MeasurementController) Get(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	measurement, err := c.service.Get(id)
	if err != nil {
		c.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, measurement)
}

func (c *MeasurementController) GetByDeviceId(ctx *gin.Context) {
	device := ctx.MustGet("User").(*models.DeviceProfile)
	measurements, err := c.service.GetByDeviceId(device.UserName)
	if err != nil {
		c.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, measurements)
}

func (c *MeasurementController) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	err := c.service.Delete(id)
	if err != nil {
		c.finishWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, "")
}

func (req *MeasurementController) finishWithError(ctx *gin.Context, status int, err error) {
	var response = struct {
		Error string `json:"error"`
	}{Error: err.Error()}

	ctx.JSON(status, response)
}
