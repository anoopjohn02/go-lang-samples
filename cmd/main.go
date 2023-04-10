package main

import (
	"com/anoop/examples/internal/client"
	"com/anoop/examples/internal/data/repo/mongorepo"
	"com/anoop/examples/internal/service"
	"com/anoop/examples/internal/service/mqtt"
	"com/anoop/examples/internal/token"
	"com/anoop/examples/internal/web"
	"fmt"
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
	accountUrl := "https://demo-accounts.ioto.cloud/v1"
	mqttUrl := "tcp://demo-mqcon.ioto.cloud:8880"
	mqttUserName := "dev-device"
	mqttPassword := "test123"
	mqttTopic := fmt.Sprintf("app/message/%s", mqttUserName)

	mongorep := mongorepo.NewMongoRepo(uri)
	db := mongorep.Db()
	mqttMessageService := mqtt.NewMqttMessageService(mqttUrl, mqttUserName, mqttPassword, mqttTopic)

	alertRepo := mongorepo.NewAlertRepository(db)
	alertService := service.NewAlertService(alertRepo, mqttMessageService)
	deviceClient := client.NewDeviceClient(accountUrl)
	tokenValidator := token.NewTokenValidator(deviceClient)

	return &commons.DeviceContext{AlertService: alertService,
		TokenValidator: tokenValidator, MessageService: mqttMessageService}
}
