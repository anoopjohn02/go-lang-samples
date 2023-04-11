package web

import (
	"context"
	"log"
	"net/http"

	"com/anoop/examples/internal/commons"
	"com/anoop/examples/internal/web/middlewares"

	"github.com/gin-gonic/gin"
)

type RestService struct {
	engine *gin.Engine
	server *http.Server
	stop   chan string
}

func (r *RestService) Start() {
	go func() {
		if err := r.server.ListenAndServe(); err != nil {
			log.Printf("Error with http server: %v", err)
		}
	}()

	<-r.stop
}

func (r *RestService) Stop() {
	var stopMsg = "stop"
	r.stop <- stopMsg
	err := r.server.Shutdown(context.Background())
	if err != nil {
		log.Printf("Error with http server: %v", err)
	}
}

func NewService(deviceContext *commons.DeviceContext) *RestService {
	engine := gin.Default()
	setupRoutes(engine, deviceContext)

	return &RestService{
		engine: engine,
		server: &http.Server{Addr: ":" + "8080", Handler: engine},
		stop:   make(chan string),
	}
}

func setupRoutes(engine *gin.Engine, deviceContext *commons.DeviceContext) {
	engine.GET("/api/ping", ping)
	v1 := engine.Group("/v1")
	{
		v1.Use(middlewares.JwtAuthMiddleware(deviceContext.TokenValidator))
		alerts := v1.Group("/alerts")
		{
			alertController := NewAlertController(deviceContext.AlertService)
			alerts.GET("/:id", alertController.Get)
			alerts.GET("", alertController.GetByDeviceId)
			alerts.POST("", alertController.Send)
			alerts.DELETE("/:id", alertController.Delete)
		}
		measurements := v1.Group("/measurements")
		{
			measurementController := NewMeasurementController(*deviceContext.MeasurementService)
			measurements.GET("/:id", measurementController.Get)
			measurements.GET("", measurementController.GetByDeviceId)
			measurements.POST("", measurementController.Send)
			measurements.DELETE("/:id", measurementController.Delete)
		}
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
