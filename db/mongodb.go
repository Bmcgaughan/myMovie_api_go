package db

import (
	"api_go/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects to the database
func ConnectDB() *mongo.Client {
	// Set client options
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MainConfig.MongoDBURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}

// get collection names from database
func GetCollectionNames(client *mongo.Client) []string {
	collections, err := client.Database("myFlixDB").ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return collections
}
