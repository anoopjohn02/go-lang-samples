package service

import "github.com/gin-gonic/gin"

type MeasurementService struct {
}

func NewMeasurementService() *MeasurementService {
	return &MeasurementService{}
}

func (s *MeasurementService) send(ctx *gin.Context) {

}

func (s *MeasurementService) get(ctx *gin.Context) {

}

func (s *MeasurementService) getByDeviceId(ctx *gin.Context) {

}

func (s *MeasurementService) delete(ctx *gin.Context) {

}
