package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	NotesCollection *mongo.Collection
)

func InitNotesCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("notesApp").Collection("notes")
}

func InitMongoDB(url string) (*mongo.Client, error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, nil
}

func InitCollections(client *mongo.Client) {
	NotesCollection = InitNotesCollection(client)
}
