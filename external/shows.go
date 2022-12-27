package external

import "api_go/models"

// Genres is a map of genre ids to genre names
var Genres = map[int64]string{
	10759: "Action & Adventure",
	16:    "Animation",
	35:    "Comedy",
	80:    "Crime",
	99:    "Documentary",
	18:    "Drama",
	10751: "Family",
	10762: "Kids",
	9648:  "Mystery",
	10763: "News",
	10764: "Reality",
	10765: "Sci-Fi & Fantasy",
	10766: "Soap",
	10767: "Talk",
	10768: "War & Politics",
	37:    "Western",
}

func getGenre(ids []int64) models.DBGenre {
	// return first genre name or empty string
	for _, id := range ids {
		if name, ok := Genres[id]; ok {
			return models.DBGenre{
				Name:        name,
				Description: "",
			}
		}
	}
	return models.DBGenre{
		Name:        "",
		Description: "",
	}
}

func getNetwork(networks []Network) string {
	// return first network or empty string
	if len(networks) > 0 {
		return networks[0].Name
	}
	return ""
}

func getDirector(crew []Cast) models.Director {
	// return first director or empty director
	for _, person := range crew {
		if *person.Job == "Director" {
			return models.Director{
				Name: person.Name,
				Bio:  "",
			}
		}
	}
	return models.Director{
		Name: "",
		Bio:  "",
	}
}

func getActors(cast []Cast) []string {
	// return first 5 actors or empty actors
	var actors []string
	for i, person := range cast {
		if i > 2 {
			break
		}
		actors = append(actors, person.Name)
	}
	return actors
}
