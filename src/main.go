package main

import (
	"booleanservice/src/middleware"
	"booleanservice/src/router"
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := middleware.StartDb()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("DB started Sucssessfully")

	r := router.SetupRouter()
	r.Run()
	middleware.CloseDb()
}
