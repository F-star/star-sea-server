package main

import (
	"log"
	"os"
	"star-sea-server/api/view"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error readin .env: ", err)
		os.Exit(1)
	}

	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	view.StartServer()
}
