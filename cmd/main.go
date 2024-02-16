package main

import (
	"log"
	"service-task-list/config"
	"service-task-list/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	log.Println("start running app")

	app.Run(cfg)
}
