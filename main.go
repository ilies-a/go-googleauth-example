package main

import (
	"log"
	"os"

	"github.com/ilies-a/go-googleauth-example/app"
	"github.com/ilies-a/go-googleauth-example/app/models"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "prod" {
		loadDotEnv()
	}
	models.InitFakeDb()
	app.StartGinServer()
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
