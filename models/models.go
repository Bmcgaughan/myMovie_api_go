package models

// make mongo User Model

// User is the model for a user
type User struct {
	UserName       string  `json:"Username" bson:"Username"`
	Password       string  `json:"Password" bson:"Password"`
	FavoriteMovies []int32 `json:"FavoriteMovies" bson:"FavoriteMovies"`
}

// make mongo Movie Model
type Movie struct {
	ID          string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Genre       DBGenre  `json:"genre" bson:"genre"`
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

type DBGenre struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type Director struct {
	Name string `json:"name" bson:"name"`
	Bio  string `json:"bio" bson:"bio"`
}
