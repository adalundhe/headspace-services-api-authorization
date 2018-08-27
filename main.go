package main

import (
	"github.com/scorbettUM/headspace/client_application/api_services/authorization_service/tools"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tools.Init()
	StartServer()

}
