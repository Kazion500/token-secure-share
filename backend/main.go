package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kazion500/secure-share/app"
	"github.com/kazion500/secure-share/database"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := LoadEnv(); err != nil {
		log.Fatal("Error loading .env file")
	}

	err := database.Connect()

	if err != nil {
		panic(err)
	}

	app := app.CreateApp()

	app.Listen("0.0.0.0:5000")
}
