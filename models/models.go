package models

// make mongo User Model

// User is the model for a user
type User struct {
	UserName       string `json:"Username" bson:"Username"`
	Password       string `json:"Password" bson:"Password"`
	FavoriteMovies []int  `json:"FavoriteMovies" bson:"FavoriteMovies"`
}

// make mongo Movie Model
type Movie struct {
	ID          string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Genre       DBGenre  `json:"genre" bson:"genre"`
	Director    Director `json:"director" bson:"director"`
	Actors      []string `json:"actors" bson:"actors"`
	ImagePath   string   `json:"imagepath" bson:"imagepath"`
	Featured    bool     `json:"featured" bson:"featured"`
	Rating      float64  `json:"rating" bson:"rating"`
	Network     string   `json:"network" bson:"network"`
	Popularity  float64  `json:"popularity" bson:"popularity"`
	Trending    bool     `json:"trending" bson:"trending"`
	OdbID       int      `json:"odbid" bson:"odbid"`
	Recommended []int    `json:"recommended" bson:"recommended"`
}

type DBGenre struct {
	Name        string `json:"Name" bson:"Name"`
}

type Director struct {
	Name string `json:"Name" bson:"Name"`
}
