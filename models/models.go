package models

// make mongo User Model

// User is the model for a user
type User struct {
	UserName       string   `json:"username" bson:"username"`
	Password       string   `json:"password" bson:"password"`
	FavoriteMovies []string `json:"favoriteMovies" bson:"favoriteMovies"`
}

// make mongo Movie Model
type Movie struct {
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Genre       Genre    `json:"genre" bson:"genre"`
	Director    Director `json:"director" bson:"director"`
	Actors      []string `json:"actors" bson:"actors"`
	ImagePath   string   `json:"imagePath" bson:"ImagePath"`
	Featured    bool     `json:"featured" bson:"featured"`
	Rating      float64  `json:"rating" bson:"rating"`
	Network     string   `json:"network" bson:"network"`
	Popularity  float64  `json:"popularity" bson:"popularity"`
	Trending    bool     `json:"trending" bson:"trending"`
	OdbID       int      `json:"odbID" bson:"odbID"`
	Recommended []int    `json:"recommended" bson:"recommended"`
}

type Genre struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type Director struct {
	Name string `json:"name" bson:"name"`
	Bio  string `json:"bio" bson:"bio"`
}
