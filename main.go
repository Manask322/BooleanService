package main

import (
	"booleanservice/src/middleware"
	"booleanservice/src/router"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
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
	fmt.Println("DB started Sucssessfully")

	r := router.SetupRouter()
	r.Run()
	middleware.CloseDb()
}
