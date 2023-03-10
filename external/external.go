package external

import (
	"api_go/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	popular  = "https://api.themoviedb.org/3/tv/popular?api_key="
	trending = "https://api.themoviedb.org/3/trending/tv/week?api_key="
	baseURL  = "https://api.themoviedb.org/3/tv/"

	language  = "&language=en-US"
	page      = "&page="
	pageCount = 3
	credits   = "&append_to_response=credits"
)

// GetPopularTV returns a list of popular tv shows
func GetPopularTMDB() (*[]models.Movie, error) {
	var showResults = []Results{}

	for i := 1; i <= pageCount; i++ {
		url := popular + os.Getenv("API_KEY") + language + page + fmt.Sprint(i)
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
		for _, show := range shows.Result {
			if strings.ToLower(string(show.OriginalLanguage)) != "en" {
				continue
			}
			showResults = append(showResults, show)
		}
	}
	processShows := Shows{Result: showResults}
	returnShows := processShows.ConvertToMovie()

	return returnShows, nil
}

// GetTrendingTV returns a list of trending tv shows
func GetTrendingTMDB() (*[]models.Movie, error) {
	var trendResults = []Results{}

	for i := 1; i <= pageCount; i++ {
		url := trending + os.Getenv("API_KEY") + language + page + fmt.Sprint(i)
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
		trendResults = append(trendResults, shows.Result...)
	}

	processShows := Shows{Result: trendResults}
	returnShows := processShows.ConvertToMovie()

	return returnShows, nil
}

// GetRecommendedTV returns a list of recommended tv shows
func GetRecommendedTMDB(id string) (*[]models.Movie, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/tv/%s/recommendations?api_key=%s%s&page=1%s", id, os.Getenv("API_KEY"), language, credits)

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

// SearchTMDB returns a list of tv shows based on search query
func SearchTMDB(query string) (*[]models.Movie, error) {
	uriEncodeQuery := url.QueryEscape(query)
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/tv?api_key=%s%s&page=1&query=%s&include_adult=false", os.Getenv("API_KEY"), language, uriEncodeQuery)

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
	url := fmt.Sprintf("https://api.themoviedb.org/3/tv/%d?api_key=%s%s%s", id, os.Getenv("API_KEY"), language, credits)

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
