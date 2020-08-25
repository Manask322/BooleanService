package main

import (
	"booleanservice/src/middleware"
	"booleanservice/src/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = middleware.StartDb()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	r := router.SetupRouter()
	r.Run()
	middleware.CloseDb()
}
