package service

import "github.com/gin-gonic/gin"

type AlertService struct {
}

func NewAlertService() *AlertService {
	return &AlertService{}
}

func (s *AlertService) send(ctx *gin.Context) {

}

func (s *AlertService) get(ctx *gin.Context) {

}

func (s *AlertService) getByDeviceId(ctx *gin.Context) {

}

func (s *AlertService) delete(ctx *gin.Context) {

}
