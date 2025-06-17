package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse: ", log.LstdFlags|log.Lshortfile)

	app := server.NewServer(logger)

	logger.Println("Starting servera on :8080")
	err := app.Server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
