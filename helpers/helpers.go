package helpers

import (
	ext "api_go/external"
	"api_go/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTV(client *mongo.Client) ([]models.Movie, error) {
	collection := client.Database("myFlixDB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var movies []models.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		log.Println(err)
		return nil, err
	}
	return movies, nil
}

func GetTVByTitle(client *mongo.Client, title string) (models.Movie, error) {
	collection := client.Database("myFlixDB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var movie models.Movie
	err := collection.FindOne(ctx, bson.M{"Title": title}).Decode(&movie)
	if err != nil {
		log.Println(err)
		return movie, err
	}
	return movie, nil
}

func AddFavorite(client *mongo.Client, username string, movieID string) (models.User, error) {
	collection := client.Database("myFlixDB").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var user models.User
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	// find one and update and return new record
	err := collection.FindOneAndUpdate(ctx, bson.M{"Username": username}, bson.M{"$addToSet": bson.M{"FavoriteMovies": movieID}}, options).Decode(&user)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func RemoveFavorite(client *mongo.Client, username string, movieID string) (models.User, error) {
	collection := client.Database("myFlixDB").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var user models.User
	// find one and update and return new record
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := collection.FindOneAndUpdate(ctx, bson.M{"Username": username}, bson.M{"$pull": bson.M{"FavoriteMovies": movieID}}, options).Decode(&user)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func GetPopularTV(client *mongo.Client) (*[]models.Movie, error) {
	shows, err := ext.GetPopularTMDB()
	if err != nil {
		return nil, err
	}

	addToDB(client, shows)

	return shows, nil

}

func addToDB(client *mongo.Client, shows *[]models.Movie) {
	// add to db if not already there
	collection := client.Database("myFlixDB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for _, show := range *shows {
		var result bson.M
		err := collection.FindOne(ctx, bson.M{"odbID": show.OdbID}).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				_, err = collection.InsertOne(ctx, show)
				if err != nil {
					log.Println(err)
				}
			}

		}
	}
}
