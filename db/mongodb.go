package db

import (
	"context"
	"log"
	"os"
	"time"

	"api_go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects to the database
func ConnectDB() *mongo.Client {
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

// function to get all Movies from Movies Collection
func GetAllMovies(client *mongo.Client) []models.Movie {
	collection := client.Database("myFlixDB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var movies []models.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		log.Fatal(err)
	}
	return movies
}
