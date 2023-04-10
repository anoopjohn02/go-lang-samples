package web

import (
	"com/anoop/examples/internal/service"

	"github.com/gin-gonic/gin"
)

type MeasurementController struct {
	service service.MeasurementService
}

func NewMeasurementController(service service.MeasurementService) *MeasurementController {
	return &MeasurementController{service: service}
}

func (s *MeasurementController) send(ctx *gin.Context) {

}

func (s *MeasurementController) get(ctx *gin.Context) {

}

func (s *MeasurementController) getByDeviceId(ctx *gin.Context) {

}

func (s *MeasurementController) delete(ctx *gin.Context) {

}
