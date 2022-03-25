package main

import (
	"fuckoff-server/model"
	"fuckoff-server/route"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	db, _ := model.DBConnection()
	route.SetupRoutes(db)
}
