package main

import (
	"com/anoop/examples/internal/data/repo/mongorepo"
	"com/anoop/examples/internal/service"
	"com/anoop/examples/internal/web"
	"log"

	"com/anoop/examples/internal/commons"
)

func main() {
	log.Printf("Starting Sample App...")
	context := Context()
	webService := web.NewService(context)
	webService.Start()
	log.Printf("Application Started...")
}

func Context() *commons.DeviceContext {

	uri := "mongodb://root:root@localhost:27017/admin"
	mongorep := mongorepo.NewMongoRepo(uri)
	db := mongorep.Db()

	alertRepo := mongorepo.NewAlertRepository(db)
	alertService := service.NewAlertService(alertRepo)

	return &commons.DeviceContext{AlertService: alertService}
}
