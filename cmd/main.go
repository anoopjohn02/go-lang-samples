package main

import (
	"com/anoop/examples/internal/web"
	"log"
)

func main() {
	log.Printf("Starting Sample App...")
	webService := web.NewService()
	webService.Start()
	log.Printf("Application Started...")
}
