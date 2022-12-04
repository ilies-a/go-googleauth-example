package main

import (
	"log"
	"os"
	"runtime"

	"github.com/ilies-a/go-googleauth-example/app"
	"github.com/ilies-a/go-googleauth-example/app/models"
	"github.com/joho/godotenv"
)

func main() {
	ConfigRuntime()
	if os.Getenv("APP_ENV") != "prod" {
		loadDotEnv()
	}
	models.InitFakeDb()
	app.StartGinServer()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.Printf("Running with %d CPUs\n", nuCPU)
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
