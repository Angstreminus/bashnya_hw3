package main

import (
	"log"

	"uniq/internal/app"
)

func main() {
	app := app.New()
	if err := app.Run(); err != nil {
		log.Fatalf("Error to run app: %v", err.Error())
	}
}
