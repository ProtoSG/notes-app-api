package main

import (
	"log"
	"notes-app-api/cmd/config"
	"notes-app-api/cmd/db"
	"notes-app-api/cmd/utils"
)

func main() {
	cfg := config.Load()

	client, err := db.InitMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}

	db.InitCollections(client)

	server := utils.NewServer()

	println("Server running on port 8080")
	log.Fatal(server.ListenAndServe())
}
