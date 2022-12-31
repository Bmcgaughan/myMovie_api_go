package helpers

import (
	ext "api_go/external"
	"api_go/models"
	"context"
	"errors"
	"log"
	"sort"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	topMovies = 10
)

func CreateUser(client *mongo.Client, user models.User) (models.User, error) {
	collection := client.Database("myFlixDB").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result bson.M
	err := collection.FindOne(ctx, bson.M{"Username": user.UserName}).Decode(&result)
	if err == nil {
		return user, errors.New("user already exists")
	}

	if err == mongo.ErrNoDocuments {
		_, err = collection.InsertOne(ctx, user)
		if err != nil {
			log.Println(err)
		}
	}

	return user, nil
}

func GetUser(client *mongo.Client, username string) (models.User, error) {
	collection := client.Database("myFlixDB").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result bson.M
	err := collection.FindOne(ctx, bson.M{"Username": username}).Decode(&result)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	user := models.User{
		UserName: result["Username"].(string),
		Password: result["Password"].(string),
	}

	return user, nil
}

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

	go addToDB(client, shows)

	//sort by popularity
	sortShows(shows)

	return shows, nil
}

func GetTrendingTV(client *mongo.Client) (*[]models.Movie, error) {
	shows, err := ext.GetTrendingTMDB()
	if err != nil {
		return nil, err
	}

	go addToDB(client, shows)

	sortShows(shows)

	return shows, nil
}

func GetRecommendedTV(client *mongo.Client, id string) (*[]models.Movie, error) {
	// convert id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	// see if show has recommendations in db
	existShows, err := getRecommendedFromDB(client, idInt)
	if err != nil {
		log.Println("No recommendations in db")
	}

	if existShows != nil {
		return existShows, nil
	}

	shows, err := ext.GetRecommendedTMDB(id)
	if err != nil {
		return nil, err
	}

	sortShows(shows)

	go addToRecommended(client, idInt, shows)

	go addToDB(client, shows)

	return shows, nil
}

func GetMostRecommendedTV(client *mongo.Client, username string) (*[]models.Movie, error) {
	// get users FavoriteMovies from db
	userFavorites, err := getUserFavorites(client, username)
	if err != nil || len(*userFavorites) == 0 {
		log.Println(err)
		return nil, err
	}

	collection := client.Database("myFlixDB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"_id": bson.M{"$in": userFavorites}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var movies []models.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		log.Println(err)
		return nil, err
	}
	// tally up recommendations
	topMovies, err := tallyRecommended(client, &movies)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return topMovies, nil

}

func GetUserDetails(client *mongo.Client, username string) (models.User, error) {
	collection := client.Database("myFlixDB").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var user models.User
	err := collection.FindOne(ctx, bson.M{"Username": username}).Decode(&user)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func tallyRecommended(client *mongo.Client, movies *[]models.Movie) (*[]models.Movie, error) {
	recommendations := make(map[int]int)
	for _, movie := range *movies {
		for _, rec := range movie.Recommended {
			recommendations[rec]++
		}
	}

	// sort recommendation map by value
	type keyVal struct {
		Key   int
		Value int
	}

	var tallyIds []keyVal
	for k, v := range recommendations {
		tallyIds = append(tallyIds, keyVal{k, v})
	}

	sort.Slice(tallyIds, func(i, j int) bool {
		return tallyIds[i].Value > tallyIds[j].Value
	})

	var sortedReco []models.Movie
	for _, val := range tallyIds {
		movie, err := getMovieByID(client, val.Key)
		if err != nil {
			log.Println(err)
			continue
		}
		sortedReco = append(sortedReco, movie)
	}

	return &sortedReco, nil

}

func getMovieByID(client *mongo.Client, id int) (models.Movie, error) {
	collection := client.Database("myFlixDB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var movie models.Movie
	err := collection.FindOne(ctx, bson.M{"odbID": id}).Decode(&movie)
	if err != nil {
		log.Println(err)
		return movie, err
	}
	return movie, nil
}

func getUserFavorites(client *mongo.Client, username string) (*[]primitive.ObjectID, error) {
	collection := client.Database("myFlixDB").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var user models.User
	err := collection.FindOne(ctx, bson.M{"Username": username}).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.FavoriteMovies, nil
}

func SearchTV(client *mongo.Client, query string) (*[]models.Movie, error) {
	shows, err := ext.SearchTMDB(query)
	if err != nil {
		return nil, err
	}

	go addToDB(client, shows)

	return shows, nil
}

func getRecommendedFromDB(client *mongo.Client, id int) (*[]models.Movie, error) {
	collection := client.Database("myFlixDB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var movie models.Movie
	// return only Recommended field
	err := collection.FindOne(ctx, bson.M{"odbID": bson.M{"$eq": id}},
		options.FindOne().SetProjection(bson.M{"Recommended": 1, "_id": 0})).Decode(&movie)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// if movie has recommendations get them from db
	if len(movie.Recommended) > 0 {
		var shows []models.Movie
		cursor, err := collection.Find(ctx, bson.M{"odbID": bson.M{"$in": movie.Recommended}})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if err = cursor.All(ctx, &shows); err != nil {
			log.Println(err)
			return nil, err
		}
		sortShows(&shows)

		return &shows, nil

	}
	return nil, nil
}

func addToRecommended(client *mongo.Client, show int, shows *[]models.Movie) {
	var showIDsToAdd []int
	for _, show := range *shows {
		showIDsToAdd = append(showIDsToAdd, show.OdbID)
	}
	collection := client.Database("myFlixDB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var result bson.M
	updateErr := collection.FindOneAndUpdate(ctx, bson.M{"odbID": show}, bson.M{"$addToSet": bson.M{"Recommended": bson.M{"$each": showIDsToAdd}}}).Decode(&result)
	if updateErr != nil {
		log.Println(updateErr)
	}

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

func sortShows(shows *[]models.Movie) {
	sort.Slice(*shows, func(i, j int) bool {
		return (*shows)[i].Popularity > (*shows)[j].Popularity
	})
}
