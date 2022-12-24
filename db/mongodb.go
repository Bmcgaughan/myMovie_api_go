package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectDB connects to the database
func ConnectDB() {
	// Set client options
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	log.Println("Connected to MongoDB!")
}

// get collection names from database
func GetCollectionNames(client *mongo.Client) []string {
	collections, err := client.Database("myFlixDB").ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return collections
}
