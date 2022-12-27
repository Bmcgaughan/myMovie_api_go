package external

import (
	"api_go/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	popular = "https://api.themoviedb.org/3/tv/popular?api_key="

	language = "&language=en-US"
	page     = "&page="
	credits  = "&append_to_response=credits"
)

// GetPopularTV returns a list of popular tv shows
func GetPopularTMDB() (*[]models.Movie, error) {
	url := popular + os.Getenv("API_KEY") + language

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var shows Shows
	err = json.NewDecoder(response.Body).Decode(&shows)
	if err != nil {
		return nil, err
	}

	returnShows := shows.ConvertToMovie()

	return returnShows, nil
}

func getDetails(id int64) (*Details, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/tv/%d?api_key=%s&language=en-US&%s", id, os.Getenv("API_KEY"), credits)

	response, err := http.Get(url)
	if err != nil {
		return &Details{}, err
	}
	defer response.Body.Close()

	var details Details
	err = json.NewDecoder(response.Body).Decode(&details)
	if err != nil {
		return &Details{}, err
	}

	return &details, nil
}
