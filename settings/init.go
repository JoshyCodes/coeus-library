package settings

import (
	"log"
	"github.com/joho/godotenv"
)

func Init(status *bool) {
	var configFile string

	if *status {
		log.Println("Loading Production Settings...")
		configFile = "configs/env/prod/.env"
	} else {
		log.Println("Loading Development Settings...")
		configFile = "configs/env/dev/.env"
	}
	
	err := godotenv.Load(configFile)
	if err != nil {
	  log.Println("Error loading .env file")
	}
}