package web

import (
	"context"
	"log"
	"net/http"
	"os"

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

func NewService() *RestService {
	var port = os.Getenv("PORT")
	engine := gin.Default()
	setupRoutes(engine)

	return &RestService{
		engine: engine,
		server: &http.Server{Addr: ":" + port, Handler: engine},
		stop:   make(chan string),
	}
}

func setupRoutes(engine *gin.Engine) {
	engine.GET("/api/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
