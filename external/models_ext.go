package external

import (
	"api_go/models"
	"log"
)

type Shows struct {
	Page         int64     `json:"page"`
	Results      []Results `json:"results"`
	TotalPages   int64     `json:"total_pages"`
	TotalResults int64     `json:"total_results"`
}

type Results struct {
	Adult            bool             `json:"adult"`
	BackdropPath     string           `json:"backdrop_path"`
	ID               int64            `json:"id"`
	Name             string           `json:"name"`
	OriginalLanguage OriginalLanguage `json:"original_language"`
	OriginalName     string           `json:"original_name"`
	Overview         string           `json:"overview"`
	PosterPath       string           `json:"poster_path"`
	MediaType        MediaType        `json:"media_type"`
	GenreIDS         []int64          `json:"genre_ids"`
	Popularity       float64          `json:"popularity"`
	FirstAirDate     string           `json:"first_air_date"`
	VoteAverage      float64          `json:"vote_average"`
	VoteCount        int64            `json:"vote_count"`
	OriginCountry    []string         `json:"origin_country"`
}

type MediaType string

const (
	Tv MediaType = "tv"
)

type OriginalLanguage string

const (
	En OriginalLanguage = "en"
	Ja OriginalLanguage = "ja"
	Ko OriginalLanguage = "ko"
)

// Generated by https://quicktype.io

type Details struct {
	Genres   []Genre   `json:"genres"`
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Networks []Network `json:"networks"`
	Credits  Credits   `json:"credits"`
}

type Credits struct {
	Cast []Cast `json:"cast"`
	Crew []Cast `json:"crew"`
}

type Cast struct {
	Adult              bool    `json:"adult"`
	Gender             int64   `json:"gender"`
	ID                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	Character          *string `json:"character,omitempty"`
	CreditID           string  `json:"credit_id"`
	Order              *int64  `json:"order,omitempty"`
	Department         *string `json:"department,omitempty"`
	Job                *string `json:"job,omitempty"`
}

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Network struct {
	ID            int64         `json:"id"`
	Name          string        `json:"name"`
	LogoPath      string        `json:"logo_path"`
	OriginCountry OriginCountry `json:"origin_country"`
}

type OriginCountry string

const (
	Jp OriginCountry = "JP"
	US OriginCountry = "US"
)

func (s *Shows) ConvertToMovie() *[]models.Movie {
	baseURL := "http://image.tmdb.org/t/p/original"
	var movies []models.Movie

	for _, show := range s.Results {
		details, err := getDetails(show.ID)
		if err != nil {
			log.Println("Error getting details for show: ", show.Name)
		}

		parsed := models.Movie{
			Title:       show.Name,
			Description: show.Overview,
			OdbID:       int(show.ID),
			ImagePath:   baseURL + show.PosterPath,
			Popularity:  show.Popularity,
			Rating:      show.VoteAverage,
			Genre:       getGenre(show.GenreIDS),
			Network:     getNetwork(details.Networks),
			Director:    getDirector(details.Credits.Crew),
			Actors:      getActors(details.Credits.Cast),
		}
		if parsed.Description == "" {
			parsed.Description = "No description available"
		}
		movies = append(movies, parsed)
	}
	return &movies
}
